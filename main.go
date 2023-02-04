package main

import (
	"log"
	"math/rand"
	"os"

	"github.com/circa10a/postcard-mailer/internal/config"
	"github.com/circa10a/postcard-mailer/internal/images"
	"github.com/circa10a/postcard-mailer/internal/order"
	"github.com/circa10a/postcard-mailer/internal/pdf"
)

func main() {
	logger := log.Default()

	config := &config.Config{}
	err := config.Read("config.yaml")
	if err != nil {
		logger.Fatal(err)
	}

	err = config.Validate()
	if err != nil {
		log.Fatal(err)
	}

	imageFiles, err := images.Find(config.Config.ImageDirectory)
	if err != nil {
		logger.Fatal(err)
	}

	for _, image := range imageFiles {
		err := pdf.ConvertImage(image, config.Config.OutputDirectory)
		if err != nil {
			log.Fatal(err)
		}
	}

	logger.Printf("found %d images\n", len(imageFiles))

	pdfFiles, err := pdf.Find(config.Config.OutputDirectory)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("generated %d pdfs\n", len(pdfFiles))

	if config.Config.DryRun {
		logger.Printf("dryrun set. exiting...")
		os.Exit(0)
	}

	orderCost := 0
	for orderCount := 0; orderCount < config.Config.Count; orderCount++ {
		if orderCount != 0 {
			logger.Printf("orders created: %d\n", orderCount)
		}

		randomPDFFile := getRandomElementFromSlice(pdfFiles)
		postcardOrder, err := order.Create(randomPDFFile, *config)
		if err != nil {
			logger.Fatal(err)
		}

		orderCost += postcardOrder.Data.Total
		logger.Printf("current accumulated cost: $%d\n", orderCost/100)
	}

	logger.Printf("total cost of order: $%d\n", orderCost/100)
}

// getRandomElementFromSlice returns a random string from a slice of strings
func getRandomElementFromSlice(filenames []string) string {
	randomIndex := rand.Intn(len(filenames))
	return filenames[randomIndex]
}
