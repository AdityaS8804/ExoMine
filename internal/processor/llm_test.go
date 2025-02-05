package processor

import (
	"strings"
	"testing"
)

func TestGetJSON(t *testing.T) {
	// Test input containing JSON between backticks
	content := "Here is some JSON:\n```json\n{\"key\": \"value\"}\n```\nEnd of content."
	expected := "{\"key\": \"value\"}"
	result:=strings.TrimSpace(getJSON(content))
	if result!=expected{
		t.Errorf("Expected : %s\tResult : %s",expected,result)
	}
	//Edge Case
	emptyContent:="This is an empty json"
	result=getJSON(emptyContent)
	if result!=""{
		t.Errorf("Expected an empty string. Recieved : %s",result)
	}
}