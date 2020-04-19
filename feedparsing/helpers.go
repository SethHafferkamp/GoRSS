package feedparsing

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/SethHafferkamp/GoRSS/models"

	"github.com/jinzhu/gorm"
	"github.com/mmcdole/gofeed"
)

func parseFeed(feedUrl string) *gofeed.Feed {
	resp, err := http.Get("https://www.mongodb.com/blog/rss")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fp := gofeed.NewParser()
	feed, _ := fp.ParseString(string(body))
	return feed
}

func createItemsFromParsedFeed(feedID int, feed *gofeed.Feed, db *gorm.DB) {
	for _, item := range feed.Items {
		var feedItem models.FeedItem
		db.FirstOrCreate(&feedItem, &models.FeedItem{
			FeedUrlID:       feedID,
			Title:           item.Title,
			Description:     item.Description,
			Content:         item.Content,
			Link:            item.Link,
			Updated:         item.Updated,
			UpdatedParsed:   item.UpdatedParsed,
			Published:       item.Published,
			PublishedParsed: item.PublishedParsed,
			GUID:            item.GUID,
		})
	}
}

func parseFeedAndSavePull(feedUrl string, feedID int, db *gorm.DB) {
	parsedFeed := parseFeed(feedUrl)
	db.Create(&models.FeedPull{FeedUrlID: feedID})
	createItemsFromParsedFeed(feedID, parsedFeed, db)
	var feed models.FeedUrl
	db.First(&feed, feedID)
	db.Model(&feed).Update("LastCompletedDate", &[]time.Time{time.Now()}[0])
}

func feedWorker(id int, jobs <-chan ParseJob, results chan<- int) {

	for j := range jobs {
		wdb := models.GetDB()
		fmt.Println("worker", id, "started  job", j)
		parseFeedAndSavePull(j.feedUrl, j.feedId, wdb)
		fmt.Println("worker", id, "finished job", j)
		results <- id
		wdb.Close()
	}
}
