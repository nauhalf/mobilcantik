package response

type ResponseError struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	ErrorCode interface{} `json:"errors"`
}

type ResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Metadata struct {
	Page        int `json:"page"`
	TotalPage   int `json:"total_pages"`
	ItemPerPage int `json:"item_per_page"`
	TotalItems  int `json:"total_items"`
}
