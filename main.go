package main

import (
	"fmt"
	"os"

	"github.com/andreybleme/cloud-scan/aws"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "s3",
				Aliases: []string{"s3"},
				Usage:   "Scan your AWS S3 buckets",
				Action: func(c *cli.Context) error {
					region := c.String("region")
					ignoredBucket := c.String("ignore")
					aws.CheckBuckets(region, ignoredBucket)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "region",
						Value:   "us-east-1",
						Aliases: []string{"l"},
						Usage:   "availability region",
					},
					&cli.StringFlag{
						Name:    "ignore",
						Value:   "my-bucket",
						Aliases: []string{"i"},
						Usage:   "bucket name to be ignored",
					},
				},
			},
		},
	}
	app.Name = "cloud-scan"
	app.HelpName = app.Name
	app.Version = "v0.1.0"
	app.Usage = "A CLI tool to scan your cloud resources and show security issues"

	err := app.Run(os.Args)
	if err != nil {
		fmt.Print(err)
	}
}
