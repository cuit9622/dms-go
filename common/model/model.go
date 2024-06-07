package model

type PageRequest struct {
	Page     int32 `form:"page"`
	PageSize int32 `form:"pageSize"`
}

type PageResult struct {
	Total   int64 `json:"total"`
	Records any   `json:"records"`
}
