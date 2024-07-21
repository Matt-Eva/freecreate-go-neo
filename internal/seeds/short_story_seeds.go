package seeds

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/models"

	"github.com/go-faker/faker/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedShortStories(ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client) err.Error {

	return err.Error{}
}

func makeShortStory(){
	w := models.PostedWriting{
		Title: faker.Sentence(),
	}
	p := models.PostedShortStory{
		w,
	}
	fmt.Println(p.Title)
}
