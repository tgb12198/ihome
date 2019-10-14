package commons

type Result struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
	Data interface{} `json:"data"`
}
