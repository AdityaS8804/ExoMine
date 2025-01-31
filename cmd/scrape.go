package cmd

import (
	"github.com/spf13/cobra"

	"github.com/AdityaS8804/ExoMine.git/internal/scraper"
)


var scrapeCmd=&cobra.Command{
	Use:"scrape [url] [json]",
	Short:"To scrape details from the given link",
	Args:cobra.ExactArgs(2),
	Run:func(cmd *cobra.Command, args []string){
		url:=args[0]
		JsonFormat:=args[1]
		scraper.ScrapeURL(url,JsonFormat)
	},
}

