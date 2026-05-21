package helper

import (
	"context"
	"os"
	"time"

	"backend/config"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetDB returns the database name from env, fallback ke "kampus"
func GetDB() string {
	db := os.Getenv("MONGODB_NAME")
	if db == "" {
		return "kampus"
	}
	return db
}

// GetCollection returns a MongoDB collection reference dari database utama.
// Semua modul wajib pakai fungsi ini — jangan akses config.MongoClient langsung.
func GetCollection(collectionName string) *mongo.Collection {
	return config.MongoClient.Database(GetDB()).Collection(collectionName)
}

// GetContext returns a context with timeout for MongoDB operations
func GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
