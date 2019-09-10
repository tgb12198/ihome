package commons

type Result struct {
	ErrNo int `json:"errno"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
}
