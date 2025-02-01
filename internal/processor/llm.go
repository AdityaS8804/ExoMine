package processor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/AdityaS8804/ExoMine.git/utils"
	"golang.org/x/net/html"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestPayload struct{
	Messages []Message `json:"messages"`
	AgentID string `json:"agent_id"`
}


type ResponsePayload struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func LLMFetch(body []byte,JsonFormat string){
	req:=postLLM(body,JsonFormat) //return a request object
	client:=&http.Client{}
	res,err:=client.Do(req)//sends the request and gets response
	if err!=nil{
		fmt.Println("Error getting response")
	}
	defer res.Body.Close()
	switch res.StatusCode {
	case http.StatusOK: // 200
		fmt.Println("Success: Status 200 OK")
		processResponse(res)//processes the response
	case http.StatusUnprocessableEntity: // 422
		fmt.Println("Error: Status 422 Unprocessable Entity")
	default:
		fmt.Printf("Unexpected status code: %d\n", res.StatusCode)
		failedBody,err:=io.ReadAll(res.Body)
		if err!=nil{
			fmt.Println("Error in default case")
		}
		fmt.Println(string(failedBody))
	}


	
}
func processResponse(res *http.Response){
		resBody, err:=io.ReadAll(res.Body)
		
	if err!=nil{
		fmt.Println("Error in reading response body")
	}
	var responsePayload ResponsePayload
	err = json.Unmarshal(resBody,&responsePayload)
	if err!=nil{
		fmt.Println("Error in unmarshalling to json")
	}
	if len(responsePayload.Choices)>0{
		fmt.Println(getJSON(responsePayload.Choices[0].Message.Content))
	}

}
func postLLM(body []byte,JsonFormat string) *http.Request{
	err:=utils.LoadEnvVariables()
	if err!=nil{
		fmt.Println("Error loading environment variables")
	}
	apiURL:="https://api.mistral.ai/v1/agents/completions"
	cleanedBody:=JsonFormat+removeExtraNewlines(string(cleanBody(body)))
	
	maxLength :=  131072
if len(cleanedBody) > maxLength {
    cleanedBody = cleanedBody[:maxLength]
}
//print(cleanedBody)
	payload:=RequestPayload{

		Messages: []Message{
			{
				Role:    "user",
				Content: cleanedBody,
			},
		},
		AgentID: "ag:1b69b193:20250131:webscrapper:105fca8c",
	}
	payloadBytes,err:=json.Marshal(payload)
	if err!=nil{
		fmt.Println("Error in payload loading to json")
	}
	req,err:=http.NewRequest("POST",apiURL,bytes.NewBuffer(payloadBytes))
	if err!=nil{
		fmt.Println("Error with making http request object")
	}
	//Set Header - Authorization
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("Authorization", "Bearer "+ utils.GetAPIKey())
return req
}
func cleanBody(htmlcontent []byte)[]byte{
doc,err:=html.Parse(bytes.NewReader(htmlcontent))
if err!=nil{
	fmt.Println("Error in parsing html")
}
var b bytes.Buffer
nodesRemoved:=[]string{"style","header","footer","head","script","noscript","svg","img"}
tagsRemoved:=[]string{"style","class","id"}
for _,node:=range nodesRemoved{
	removeNodes(doc,node)
}
for _,tag:=range tagsRemoved{
	traverse(doc,tag)
}
if err:=html.Render(&b,doc);err!=nil{
fmt.Println("Error in removing styles")
}
return b.Bytes()
}


func removeNodes(n *html.Node,tag string){
	
	var prevSibling *html.Node
	for c:=n.FirstChild;c!=nil;{
		next:=c.NextSibling
		if c.Type==html.ElementNode && c.Data==tag{
			if prevSibling!=nil{
				prevSibling.NextSibling=c.NextSibling
			}else{
				n.FirstChild=c.NextSibling
			}
		}else{
			removeNodes(c,tag)
			prevSibling=c
		}
		c=next
	}
}
func traverse(n *html.Node, tag string) {
    if n == nil {
        return
    }

    // Process current node if it's an element
    if n.Type == html.ElementNode {
        // Remove the tag attribute if it exists
        for i := 0; i < len(n.Attr); i++ {
            if n.Attr[i].Key == tag {
                // Remove the tag attribute by replacing it with the last attribute
                // and truncating the slice
                n.Attr[i] = n.Attr[len(n.Attr)-1]
                n.Attr = n.Attr[:len(n.Attr)-1]
                i-- // Adjust index since we modified the slice
            }
        }
    }

    // Recursively process all child nodes
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        traverse(c, tag)
    }
}


func removeExtraNewlines(content string) string {
    // Replace multiple newlines with a single newline
    re := regexp.MustCompile(`\n\s*\n`)
    content = re.ReplaceAllString(content, "\n")
    // Remove leading and trailing whitespace
    content = strings.TrimSpace(content)
    
    // Remove newlines and spaces between HTML tags
    re = regexp.MustCompile(`>\s+<`)
    content = re.ReplaceAllString(content, "><")
    
    return content
}
func getJSON(content string) string{
pattern := `(?s)` + "```json\n(.*?)```"
re := regexp.MustCompile(pattern)

matches := re.FindStringSubmatch(content)
return matches[1]
}