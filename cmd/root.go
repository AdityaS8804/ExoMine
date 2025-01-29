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
rootCmd.AddCommand(scrapeCmd)
}