package structs

type BoardRequest struct {
	UserId  string   `json:"userid"`
	BoardId string   `json:"boardid"`
	ToDo    []string `json:"todo"`
	Doing   []string `json:"doing"`
	Done    []string `json:"done"`
}

type BoardResponseGet struct {
	Boards []map[string]interface{} `json:"boards"`
}

type BoardResponseSave struct {
	StatusCode int    `json:"statusCode"`
	BoardId    string `json:"boardid"`
	Msg        string `json:"msg"`
}
