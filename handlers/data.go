package handlers

type Response struct {
	Success  bool     `json:"success"`
	Data     any      `json:"data"`
	Metadata Metadata `json:"metadata"`
}

type Data struct {
	YourData any `json:"yourData"`
}

type Metadata struct {
	Timestamp int `json:"timestamp"`
}
