package scraper

import (
	"fmt"

	"github.com/gocolly/colly"

	"github.com/AdityaS8804/ExoMine.git/internal/processor"
)
type Product struct{
	Url, Image,Name,Price string
}

func ScrapeURL(url string,JsonFormat string){
	fmt.Println("Getting Response...")
	c:=colly.NewCollector()
	c.OnResponse(func(r *colly.Response){
		processor.LLMFetch(r.Body,JsonFormat)
	})
	c.Visit(url)
}
