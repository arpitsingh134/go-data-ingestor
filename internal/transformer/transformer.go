package transformer

import (
    "time"

    "go-data-ingestor/internal/models"
)

func Transform(posts []models.Post) []models.Post {
    now := time.Now().UTC()
    source := "placeholder_api"
    for i := range posts {
        posts[i].IngestedAt = now
        posts[i].Source = source
    }
    return posts
}
