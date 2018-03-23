package main

import (
	"log"
	"os"

	lcs "github.com/emmansun/go-lcs"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	var algorithm, modelStr, sampleStr string

	app := cli.NewApp()
	app.Name = "LCS"
	app.Description = "Longest common subsequence problem algorithm"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "algorithm,a",
			Usage:       "used algorithm (fast or normal)",
			Value:       "fast",
			Destination: &algorithm,
		},
		cli.StringFlag{
			Name:        "model,m",
			Usage:       "model characters",
			Destination: &modelStr,
		},
		cli.StringFlag{
			Name:        "sample,s",
			Usage:       "sample characters",
			Destination: &sampleStr,
		},
	}
	app.Action = func(c *cli.Context) error {
		if modelStr == "" || sampleStr == "" {
			log.Fatalf("modle and sample can't be empty!")
		}
		var lcsImpl lcs.LCSInterface
		var pairs []*lcs.LcsPair
		if algorithm == "fast" {
			lcsImpl = lcs.NewFastLCSString(modelStr, sampleStr)
			pairs = lcsImpl.FindAllLcsPairs()
		} else {
			lcsImpl = lcs.NewNormalLCSString(modelStr, sampleStr)
			pairs = lcsImpl.FindAllLcsPairs()
		}
		log.Printf("[FAST] Max lcs len=%d, Candidate size=%d, Candicates=%v\n\n", lcsImpl.MaxLcsLen(), len(pairs), pairs)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
