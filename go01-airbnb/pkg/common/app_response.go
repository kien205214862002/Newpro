package common

type appResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
}

func NewSuccessResponse(data, paging interface{}) *appResponse {
	return &appResponse{data, paging}
}

func Response(data interface{}) *appResponse {
	return NewSuccessResponse(data, nil)
}

func ResponseWithPaging(data, paging interface{}) *appResponse {
	return NewSuccessResponse(data, paging)
}
