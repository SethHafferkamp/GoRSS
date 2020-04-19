package feedparsing

import (
	"time"

	"github.com/SethHafferkamp/GoRSS/models"
)

func RegisterFeed(feedUrl string) {
	db := models.GetDB()
	defer db.Close()
	db.Create(&models.FeedUrl{FeedUrl: feedUrl})
}

func ParseDBFeedAndCreateFeedItems(feedID int) {
	db := models.GetDB()
	var feedUrl models.FeedUrl
	db.First(&feedUrl, feedID)
	pasedFeed := parseFeed(feedUrl.FeedUrl)
	createItemsFromParsedFeed(feedID, pasedFeed, db)
}

type ParseJob struct {
	feedUrl string
	feedId  int
}

// ParseAllFeeds : Uses a Buffered Channel to create a worker pool of a set size of goroutines and iterate over all FeedUrls
func ParseAllFeeds(numWorkers int) {

	db := models.GetDB()
	defer db.Close()

	var numJobs int
	db.Model(&models.FeedUrl{}).Count(&numJobs)

	jobs := make(chan ParseJob, numJobs)
	results := make(chan int, numJobs)

	var feedUrls []models.FeedUrl

	// Iterating over a slice is just to avoid a for loop. I like this but I'm not sure it's a Go best practice.
	a := make([]int, numWorkers)
	for w := range a {
		go feedWorker(w, jobs, results)
	}

	db.Find(&feedUrls)
	for _, feedUrl := range feedUrls {
		feedUrl.LastKickoffDate = &[]time.Time{time.Now()}[0]
		db.Save(&feedUrl)
		newJob := ParseJob{feedUrl: feedUrl.FeedUrl, feedId: int(feedUrl.ID)}
		jobs <- newJob
	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
