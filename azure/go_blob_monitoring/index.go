package goblobmonitoring

import (
	"encoding/json"
	"os"
	"log"
	"strings"
)

func run(context map[string]interface{}, blobContent interface{}) {
	if DD_API_KEY == "" || DD_API_KEY == "<DATADOG_API_KEY>" {
		log.Println("You must configure your API key before starting this function (see ## Parameters section)")
		return
	}

	var logs []string
	switch v := blobContent.(type) {
	case string:
		logs = strings.Split(strings.TrimSpace(v), "\n")
	case []byte:
		logs = strings.Split(strings.TrimSpace(string(v)), "\n")
	default:
		logsJSON, _ := json.Marshal(blobContent)
		logs = strings.Split(strings.TrimSpace(string(logsJSON)), "\n")
	}

	for i, log := range logs {
		logs[i] = strings.ReplaceAll(log, "'", "\"")
	}

	handler := NewBlobStorageLogHandler(context)
	parsedLogs := handler.handleLogs(logs)

	results := HTTPClient{context}.sendAll(parsedLogs)

	for _, v := range results {
		if v != true {
			log.Println("Some messages were unable to be sent. See other logs for details.")
			break
		}
	}

	//context["done"]()
}

func testLocalFile(){
	data, err := os.ReadFile("/Users/nina.rei/Downloads/azure_log.json")
	if err != nil {
		log.Fatal(err)
	}

	context := map[string]interface{}{
		"log": log.New(os.Stderr, "", 0),
		"executionContext": map[string]string{
			"functionName": "test",
		},
		"done": func() {},
	}

	run(context, data)
}

func main(context map[string]interface{}, blobContent interface{}) {
	go run(context, blobContent)
}
