package cmd

import (
	"github.com/spf13/cobra"

	"github.com/AdityaS8804/ExoMine.git/internal/scraper"
)


var scrapeCmd=&cobra.Command{
	Use:"scrape [url]",
	Short:"To scrape details from the given link",
	Args:cobra.ExactArgs(1),
	Run:func(cmd *cobra.Command, args []string){
		url:=args[0]
		scraper.ScrapeURL(url)
	},
}

