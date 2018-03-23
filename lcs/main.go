package main

import (
	"bytes"
	"fmt"
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
		log.Printf("[FAST] Model=%v, Sample=%v, Max lcs len=%d, Candidate size=%d\n", modelStr, sampleStr, lcsImpl.MaxLcsLen(), len(pairs))
		var modelCandidate, sampleCandidate bytes.Buffer
		for i, pair := range pairs {
			log.Printf("The %d candidate:\n", i+1)
			modelCandidate.WriteString("model  index: ")
			sampleCandidate.WriteString("sample index: ")
			for _, mi := range pair.ModelIndexes {
				modelCandidate.WriteString(fmt.Sprintf("%c:%d ", modelStr[mi], mi))
			}
			for _, si := range pair.SampleIndexes {
				sampleCandidate.WriteString(fmt.Sprintf("%c:%d ", sampleStr[si], si))
			}
			log.Println(modelCandidate.String())
			log.Println(sampleCandidate.String())
			log.Println()
			modelCandidate.Reset()
			sampleCandidate.Reset()
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
