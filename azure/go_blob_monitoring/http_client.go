package goblobmonitoring

type HTTPClient struct {
	Context map[string]interface{}
}

func (c HTTPClient) sendAll(logs []interface{}) []bool {
	// HTTP request logic here
	return []bool{}
}