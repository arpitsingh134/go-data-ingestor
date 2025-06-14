package main

import (
    "fmt"
    "log"

    "go-data-ingestor/config"
    "go-data-ingestor/internal/fetcher"
    "go-data-ingestor/internal/repository"
    "go-data-ingestor/internal/transformer"
)

func main() {
    cfg := config.Load()

    posts, err := fetcher.FetchPosts(cfg.APIUrl)
    if err != nil {
        log.Fatal("Error fetching posts: ", err)
    }

    transformed := transformer.Transform(posts)

    repo, err := repository.NewMongoRepo(cfg.MongoURI, cfg.Database, cfg.Collection)
    if err != nil {
        log.Fatal("Error connecting to DB: ", err)
    }

    err = repo.Save(transformed)
    if err != nil {
        log.Fatal("Error saving data: ", err)
    }

    fmt.Println("Data ingestion completed.")
}
