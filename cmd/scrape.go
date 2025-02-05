package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
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
		save,_:=cmd.Flags().GetString("save")
		if url==""{
			log.Fatal("--url flag missing")
		}
		if format==""{
			log.Fatal("--format flag missing - specify json format")
		}
		JsonFormat:=readJSONFile(format)
		data:=scraper.ScrapeURL(url,JsonFormat)
		if save!=""{
			saveCsv(data,save)
		}else{
			fmt.Println(data)
		}
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
func saveCsv(JSON_str string,filepath string){
	JSON_str="{\n"+JSON_str+"\n}"
		// Regex to match the unquoted "response"
	re := regexp.MustCompile(`(?m)(\bresponse)(\s*:)`)
	// Add double quotes around "response"
	JSON_str = re.ReplaceAllString(JSON_str, `"response"$2`)
	print(JSON_str)
var response map[string][]map[string]interface{}
err:=json.Unmarshal([]byte(JSON_str),&response)
if err!=nil{
	fmt.Println("Error in converting JSON to data in csv code",err)
}
data :=response["response"]
csvFile,err:=os.Create(filepath)
if err!=nil{
	fmt.Println("Error in creating csv file")
}
defer csvFile.Close()
writer:=csv.NewWriter(csvFile)
defer writer.Flush()
if len(data)>0{
	header:=[]string{}
	for key:=range data[0]{
		header=append(header,key)
	}
	writer.Write(header)
	for _,row:=range data{
		record:=[]string{}
		for _,key:=range header{
			record=append(record,fmt.Sprintf("%v",row[key]))
		}
		writer.Write(record)
	}
	fmt.Println("Written to csv")
}
}