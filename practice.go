package main

import (
	"fmt"
	"log"

	"encoding/csv"
	"os"

	"github.com/gocolly/colly"
)
type Product struct{
	Url, Image,Name,Price string
}

func practice(){
	fmt.Println("App Started...")
	var products []Product
	c:=colly.NewCollector()
	c.OnHTML(".product",func(r *colly.HTMLElement){
		product:=Product{
			Url:r.ChildAttr("a","href"),
			Image:r.ChildAttr("img","src"),
			Name:r.ChildText(".product-name"),
			Price:r.ChildText(".price"),
		}
		products=append(products,product)
	})
	c.OnScraped(func(r *colly.Response){
		file,err:=os.Create("WebScraped.csv")
		if err!=nil{
			log.Fatalln("Failed to create output csv")
		}
		defer file.Close()
		writer:=csv.NewWriter(file)
		writer.Write([]string{
			"Url",
			"Image",
			"Name",
			"Price",
		})
		for _,product := range products{
			writer.Write([]string{
				product.Url,
				product.Image,
				product.Name,
				product.Price,
			})
		}
		defer writer.Flush()
	})
	c.Visit("https://www.scrapingcourse.com/ecommerce")


}