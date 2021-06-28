package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	e := echo.New()

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	host := os.Getenv("MONGO_HOST")

	databases := connectToMongo(host)
	var listDatabase string
	for _, database := range databases {
		listDatabase += database + " "
	}

	html := `
	<h1>Hello World</h1>
	<h1> Nama saya "` + name + `"</h1>
	<h1> Umur saya "` + age + `"</h1>
	<h1> List Database Mongo "` + listDatabase + `"</h1>
	`

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, html)
	})

	if err := e.Start(":9000"); err != nil {
		panic(err)
	}
}

func connectToMongo(host string) []string {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + host + ":27017/?connect=direct"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	return databases
}
