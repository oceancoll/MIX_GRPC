package common


type JSONStruct struct {
	Code int `json:"code"`
	Data map[string]interface{} `json:"data"`
	Error string `json:"error"`

}

type ListJSONStruct struct {
	Code int `json:"code"`
	Data []map[string]interface{} `json:"data"`
	Error string `json:"error"`

}

