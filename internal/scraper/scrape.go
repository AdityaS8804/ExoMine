package scraper

import (
	"context"
	"fmt"
	"time"

	"github.com/AdityaS8804/ExoMine.git/internal/processor"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)
type Product struct{
	Url, Image,Name,Price string
}
func ScrapeURL(url string,JsonFormat string)string{
	fmt.Println("Getting Response...")
	ctx,cancel:=chromedp.NewContext(
		context.Background(),
	)
	defer cancel()
	var html,data string
	err:=chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(2000*time.Millisecond),
		chromedp.ActionFunc(func(ctx context.Context)error{
			rootNode, err := dom.GetDocument().Do(ctx)
              if err != nil {
                 return err
              }
              html, err = dom.GetOuterHTML().WithNodeID(rootNode.NodeID).Do(ctx)
              
			data=processor.LLMFetch([]byte(html),JsonFormat)
			return err
		}),
	)
	if err!=nil{
		fmt.Println("Error in loading chromedp")
	}
	return data

	
}
/*
func ScrapeURL(url string,JsonFormat string){
	fmt.Println("Getting Response...")
	c:=colly.NewCollector()
	c.OnResponse(func(r *colly.Response){
		processor.LLMFetch(r.Body,JsonFormat)
	})
	c.Visit(url)
}
	*/
