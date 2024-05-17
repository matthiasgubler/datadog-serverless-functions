package goblobmonitoring

import (
	"os"
	"regexp"
)

const VERSION = "1.0.2"

const (
    STRING = "string" // example: 'some message'
    STRING_ARRAY = "string-array" // example: ['one message', 'two message', ...]
    JSON_OBJECT = "json-object" // example: {"key": "value"}
    JSON_ARRAY = "json-array" // example: [{"key": "value"}, {"key": "value"}, ...] or [{"records": [{}, {}, ...]}, {"records": [{}, {}, ...]}, ...]
    BUFFER_ARRAY = "buffer-array" // example: [<Buffer obj>, <Buffer obj>]
    JSON_STRING = "json-string" // example: '{"key": "value"}'
    JSON_STRING_ARRAY = "json-string-array" // example: ['{"records": [{}, {}]}'] or ['{"key": "value"}']
    INVALID = "invalid"
    JSON_TYPE = "json"
    STRING_TYPE = "string"
)

// To scrub PII from your logs, uncomment the applicable configs below. If you'd like to scrub more than just
// emails and IP addresses, add your own config to this map in the format
// "REDACT_IP": {
//     Pattern:     regexp.MustCompile(`[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`),
//     Replacement: "xxx.xxx.xxx.xxx",
// },
// "REDACT_EMAIL": {
//     Pattern:     regexp.MustCompile(`[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+`),
//     Replacement: "xxxxx@xxxxx.com",
// }
type scrubberConfig struct{
    Pattern     *regexp.Regexp
    Replacement string
}

type SCRUBBER_RULE_CONFIGS map[string]scrubberConfig

// To split array-type fields in your logs into individual logs, you can add sections to the map below. An example of
// a potential use case with azure.datafactory is there to show the format:
// {
//   source_type:
//    paths: [list of [list of fields in the log payload to iterate through to find the one to split]],
//    keep_original_log: bool, if you'd like to preserve the original log in addition to the split ones or not,
//    preserve_fields: bool, whether or not to keep the original log fields in the new split logs
// }
// You can also set the DD_LOG_SPLITTING_CONFIG env var with a JSON string in this format.
// "azure.datafactory": {
//     Paths:           [][]string{{"properties", "Output", "value"}},
//     KeepOriginalLog: true,
//     PreserveFields:  true,
// }
type logSplitConfig struct{
    Paths           [][]string
    KeepOriginalLog bool
    PreserveFields  bool
}

type DD_LOG_SPLITTING_CONFIG map[string]scrubberConfig


var (
    DD_API_KEY = getEnvOrDefault("DD_API_KEY", "xx")
    DD_SITE = getEnvOrDefault("DD_SITE", "datad0g.com")
    DD_HTTP_URL = getEnvOrDefault("DD_URL", "http-intake.logs." + DD_SITE)
    DD_HTTP_PORT = getEnvOrDefault("DD_PORT", "443")
    DD_TAGS = getEnvOrDefault("DD_TAGS", "") // Replace '' by your comma-separated list of tags
    DD_SERVICE = getEnvOrDefault("DD_SERVICE", "azure")
    DD_SOURCE = getEnvOrDefault("DD_SOURCE", "azure")
    DD_SOURCE_CATEGORY = getEnvOrDefault("DD_SOURCE_CATEGORY", "azure")
)

func getEnvOrDefault(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}
