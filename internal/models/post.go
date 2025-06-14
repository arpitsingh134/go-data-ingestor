package models

import "time"

type Post struct {
    UserID     int       `json:"userId" bson:"userId"`
    ID         int       `json:"id" bson:"id"`
    Title      string    `json:"title" bson:"title"`
    Body       string    `json:"body" bson:"body"`
    IngestedAt time.Time `bson:"ingested_at"`
    Source     string    `bson:"source"`
}
