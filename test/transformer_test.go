package test

import (
	"testing"

	"go-data-ingestor/internal/models"
	"go-data-ingestor/internal/transformer"
)

func TestTransform(t *testing.T) {
	posts := []models.Post{
		{UserID: 1, ID: 1, Title: "Test", Body: "Test Body"},
	}
	result := transformer.Transform(posts)

	if result[0].Source != "placeholder_api" {
		t.Errorf("Expected source to be placeholder_api, got %s", result[0].Source)
	}

	if result[0].IngestedAt.IsZero() {
		t.Error("Expected IngestedAt to be set")
	}
}
