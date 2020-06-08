package response

type ResponseError struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	ErrorCode interface{} `json:"errorcode"`
}

type ResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
