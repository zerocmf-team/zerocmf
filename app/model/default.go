package model

type Paginate struct {
	Data     interface{} `json:"data"`
	Current  string      `json:"current"`
	PageSize string      `json:"page_size"`
	Total    int64         `json:"total"`
}

type ReturnData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
