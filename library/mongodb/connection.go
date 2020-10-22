package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type Config struct {
	DatabaseName string
	URI          string
}

type Connection struct {
	client   *mongo.Client
	database *mongo.Database
}

func Connect(uri string) (*Connection, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	cs, err := connstring.Parse(uri)
	if err != nil {
		panic(err)
	}
	databaseName := cs.Database

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return &Connection{
		client:   client,
		database: client.Database(databaseName),
	}, nil
}

func (c *Connection) Model(name string) *Model {
	return &Model{
		collection: c.database.Collection(name),
	}
}

func (c *Connection) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := c.client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	return nil
}
