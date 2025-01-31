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
	err:=utils.LoadEnvVariables()
	if err!=nil{
		fmt.Println("Error loading environment variables")
	}
	apiURL:="https://api.mistral.ai/v1/agents/completions"
	cleanedBody:=JsonFormat+removeExtraNewlines(string(removeStyle(body)))
	
	maxLength :=  131072
if len(cleanedBody) > maxLength {
    cleanedBody = cleanedBody[:maxLength]
}
print(cleanedBody)
	payload:=RequestPayload{

		Messages: []Message{
			{
				Role:    "user",
				Content: cleanedBody,
			},
		},
		AgentID: utils.GetAPIKey(),
	}
	payloadBytes,err:=json.Marshal(payload)
	if err!=nil{
		fmt.Println("Error in payload loading to json")
	}
	req,err:=http.NewRequest("POST",apiURL,bytes.NewBuffer(payloadBytes))
	if err!=nil{
		fmt.Println("Error with making http request object")
	}
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("Authorization", "Bearer ADRiKpjC1ss6ud6RmuQ1rbAdzU3pA0gR")

	client:=&http.Client{}
	res,err:=client.Do(req)
	if err!=nil{
		fmt.Println("Error getting response")
	}
	defer res.Body.Close()
	switch res.StatusCode {
	case http.StatusOK: // 200
		fmt.Println("Success: Status 200 OK")
		resBody, err:=io.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("Error in reading response body")
	}
	var responsePayload ResponsePayload
	err = json.Unmarshal(resBody,&responsePayload)
	if err!=nil{
		fmt.Println("Error in unmarshalling to json")
	}
	fmt.Println("Model Response:", responsePayload)
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
func removeStyle(htmlcontent []byte)[]byte{
doc,err:=html.Parse(bytes.NewReader(htmlcontent))
if err!=nil{
	fmt.Println("Error in parsing html")
}
var b bytes.Buffer
removeNodes(doc,"style")
removeNodes(doc,"header")
removeNodes(doc,"footer")
removeNodes(doc,"head")
removeNodes(doc,"script")
removeNodes(doc,"noscript")
removeNodes(doc,"svg")
removeNodes(doc,"img")
traverse(doc,"style")
traverse(doc,"class")
traverse(doc,"id")

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
func traverse(n *html.Node,tag string) {
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
