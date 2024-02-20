package repository

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func GetPost(c *gin.Context) {
	postCollection := client.Database("sisfor23").Collection("posts")
	cursor, err := postCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var posts []bson.M

	if err := cursor.All(context.Background(), &posts); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if len(posts) > 100 {
		posts = posts[:100]
	}

	c.JSON(200, posts)

}
