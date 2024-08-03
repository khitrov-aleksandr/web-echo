package handlers

type Response struct {
	Success  bool     `json:"success"`
	Data     Data     `json:"data"`
	Metadata Metadata `json:"metadata"`
}

type Data struct {
	Request  RequestData `json:"request"`
	Response any         `json:"response"`
}

type RequestData struct {
	Method      string            `json:"method"`
	Path        string            `json:"path"`
	Headers     map[string]string `json:"headers"`
	Cookies     map[string]string `json:"cookies"`
	Payload     any               `json:"payload"`
	QueryString string            `json:"queryString"`
	QueryParams map[string]string `json:"queryParams"`
	Host        string            `json:"host"`
	Ip          string            `json:"ip"`
}

type Metadata struct {
	Timestamp int `json:"timestamp"`
}
