package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:"exomine",
	Short:"ExoMine is a smart data extraction and processing tool",
	Long:`ExoMine can scrape websites, process content using LLMs, 
and provide JSON outputs optionally dumped to databases.`,
}

func Execute() error{
return rootCmd.Execute()
}
func init(){
scrapeCmd.Flags().StringP("url","u","","URL of website to be scraped")
scrapeCmd.Flags().StringP("format","f","","JSON format expected by the user")
scrapeCmd.Flags().StringP("save","s","","Save to a file")
rootCmd.AddCommand(scrapeCmd)
}