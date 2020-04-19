
# Setting up the service

# Prerequesites
--Postgres running on port 5432

Create the Database 
```
./setup_environment.sh
```

# Initialize models and load test data into database
```
go build ./cmd/init && ./init
```

# Run the service to loop through ALL feeds and scrape them and store the results
```
go build ./cmd/run && ./run
```



----
# If you ever want to reset your database, run this shortcut to drop all tables
```
  go build ./cmd/dropmodels && ./dropmodels
```