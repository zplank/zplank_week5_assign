package main

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/gocolly/colly"
)

func TestScraping(t *testing.T) {
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	//test file iterates through each url and create test .jl file
	for _, url := range urls {
		t.Run(url, func(t *testing.T) {
			// Run scraping logic
			var outputFileName = "test_output.jl"
			err := scrapeAndSave(url, outputFileName)
			if err != nil {
				t.Errorf("Error scraping URL %s: %v", url, err)
			}

			validateOutputFile(t, outputFileName)
		})
	}
}

func scrapeAndSave(url, outputFileName string) error {
	//create new collector
	c := colly.NewCollector()

	//save to .jl test output file
	file, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	c.OnHTML("p", func(e *colly.HTMLElement) {
		file.WriteString(e.Text + "\n")
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err = c.Visit(url)
	if err != nil {
		return err
	}

	c.Wait()

	return nil
}

// create function to validate output
func validateOutputFile(t *testing.T, outputFileName string) {
	file, err := os.Open(outputFileName)
	if err != nil {
		t.Errorf("Error opening output file %s: %v", outputFileName, err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	for {
		var pageContent PageContent
		if err := decoder.Decode(&pageContent); err != nil {
			break // End of file
		}

		if pageContent.URL == "" {
			t.Errorf("Invalid output format: missing URL")
		}
	}
}
