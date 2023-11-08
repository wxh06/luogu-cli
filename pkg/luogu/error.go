package luogu

type ErrorResponse struct {
	Status       int    `json:"status"`
	Data         string `json:"data"`
	ErrorMessage string `json:"errorMessage"`
	Trace        string `json:"trace"`
	CustomData   []any  `json:"customData"`
}
