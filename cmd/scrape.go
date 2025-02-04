package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/AdityaS8804/ExoMine.git/internal/scraper"
)


var scrapeCmd=&cobra.Command{
	Use:"scrape",
	Short:"To scrape details from the given link",
	
	Run:func(cmd *cobra.Command, args []string){
		url,_:= cmd.Flags().GetString("url")
		format,_:=cmd.Flags().GetString("format")
		if url==""{
			log.Fatal("--url flag missing")
		}
		if format==""{
			log.Fatal("--format flag missing - specify json format")
		}
		JsonFormat:=readJSONFile(format)
		scraper.ScrapeURL(url,JsonFormat)
	},
}

func readJSONFile(filepath string)string{
	file,err:=os.ReadFile(filepath)
	if err!=nil{
		log.Fatal("Error in reading JSON file")
	}
	return splitLines(string(file))

}
func splitLines(data string) string {
	lines := []string{}
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}
	return strings.Join(lines, "")
}