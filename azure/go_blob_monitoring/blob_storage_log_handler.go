package goblobmonitoring

import (
	"encoding/json"
	"os"
)
type BlobStorageLogHandler struct {
	Context map[string]interface{}
}

func NewBlobStorageLogHandler(context map[string]interface{}) BlobStorageLogHandler {
	return BlobStorageLogHandler{Context: context}
}

func (h BlobStorageLogHandler) handleLogs(logs []string) []interface{} {
	// handle logs logic here
	return []interface{}{}
}

var DDLogSplittingConfig = map[string]interface{}{
    // "azure.datafactory": {
    //     "paths": []interface{}{[]interface{}{"properties", "Output", "value"}},
    //     "keep_original_log": true,
    //     "preserve_fields": true,
    // },
}

func getLogSplittingConfig() map[string]interface{} {
    config, ok := os.LookupEnv("DD_LOG_SPLITTING_CONFIG")
    if !ok {
        return DDLogSplittingConfig
    }

    var configMap map[string]interface{}
    err := json.Unmarshal([]byte(config), &configMap)
    if err != nil {
        return DDLogSplittingConfig
    }

    return configMap
}

