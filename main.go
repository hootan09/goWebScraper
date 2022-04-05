package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type Item struct {
	Name     string
	Image    string
	Price    string
	Url      string
	Location string
}

func main() {
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)
	items := make([]Item, 0)

	// Callbacks

	c.OnHTML(".browse-post-list", func(e *colly.HTMLElement) {
		e.ForEach("div.post-card-item a", func(i int, h *colly.HTMLElement) {
			item := Item{}
			item.Name = h.ChildText(".kt-post-card__title")
			item.Image = h.ChildAttr("img", "src")
			item.Price = h.ChildText(".kt-post-card__description")
			item.Url = "https://divar.ir/v/" + h.Attr("href")
			item.Location = h.ChildText(".kt-post-card__bottom")
			items = append(items, item)
		})

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		js, err := json.MarshalIndent(items, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Writing data to file")
		if err := os.WriteFile("data.json", js, 0664); err == nil {
			fmt.Println("Data written to file successfully")
		}

	})

	c.Visit("https://divar.ir/s/tehran")
}
