package main

import (
	"log"
	"os"

	"github.com/gocolly/colly"
)

type PageContent struct {
	URL   string `json:"url"`
	Text  string `json:"text"`
	Error string `json:"error,omitempty"`
}

func main() {
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

	for _, url := range urls {
		err := scrapeAndSave(url)
		if err != nil {
			log.Printf("Error scraping URL %s: %v", url, err)
		}
	}
}

// create function to scrape text and save to putput
func scrapeAndSave(url string) error {
	//create new collector
	c := colly.NewCollector()

	//create .jl output file
	outputFileName := "output.jl"
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
