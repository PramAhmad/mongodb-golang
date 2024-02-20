package repository

import (
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://pramudita:tasikmalaya123..@sisfor23.u6a0v29.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	var err error
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
}

func GetPost(c *gin.Context) {
	// binding page from params
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid page"})
		return
	}
	postCollection := client.Database("myappdb").Collection("posts")
	cursor, err := postCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	// paging per page100 data
	var posts []bson.M
	for cursor.Next(context.Background()) {
		var post bson.M
		if err = cursor.Decode(&post); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		posts = append(posts, post)
	}
	// get  data
	totalData, err := postCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// get total
	totalPage := totalData / 100
	if totalData%100 != 0 {
		totalPage++
	}
	// get  per page
	var dataPerPage []bson.M
	for i := 0; i < 100; i++ {
		if i+(page-1)*100 < len(posts) {
			dataPerPage = append(dataPerPage, posts[i+(page-1)*100])
		}
	}
	// response
	c.JSON(200, gin.H{
		"totalData":   totalData,
		"totalPage":   totalPage,
		"dataPerPage": dataPerPage,
	})

}
