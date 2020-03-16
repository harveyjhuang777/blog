package storage

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/jwjhuang/blog/service/app/logger"
)

var projectID = "gorich"

// Use a service account
func NewFirestore() *firestore.Client {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		logger.Log().Fatal(err)
	}
	return client
}
