package main

import (
	"context"
	"fmt"
	"os"

	cli "github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	if err := NewApp().Run(os.Args); err != nil {
		panic(err)
	}
}

func NewApp() *cli.App {
	return &cli.App{
		Name: "mongo-docker-network",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "db-host",
				EnvVars: []string{"DB_HOST"},
			},
			&cli.StringFlag{
				Name:    "db-port",
				EnvVars: []string{"DB_PORT"},
			},
			&cli.StringFlag{
				Name:    "db-name",
				EnvVars: []string{"DB_NAME"},
			},
			&cli.StringFlag{
				Name:    "db-username",
				EnvVars: []string{"DB_USERNAME"},
			},
			&cli.StringFlag{
				Name:    "db-password",
				EnvVars: []string{"DB_PASSWORD"},
			},
		},
		Action: func(c *cli.Context) error {
			uri := fmt.Sprintf(
				"%s://%s:%s@%s:%s",
				"mongodb",
				c.String("db-username"),
				c.String("db-password"),
				c.String("db-host"),
				c.String("db-port"),
			)

			fmt.Println("trying to connect to:", uri)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			opts := options.Client().
				ApplyURI(uri)

			client, err := mongo.Connect(ctx, opts)
			if err != nil {
				return err
			}

			defer func() {
				if err = client.Disconnect(ctx); err != nil {
					panic(err)
				}
			}()

			// db := client.Database(c.String("db-name"))

			ctx, cancel = context.WithCancel(context.Background())
			defer cancel()

			if err := client.Ping(ctx, readpref.Primary()); err != nil {
				return err
			}

			fmt.Println("it didn't error")

			return nil
		},
	}
}
