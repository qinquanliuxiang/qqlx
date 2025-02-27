package schema

type IDRequest struct {
	ID uint `json:"id" form:"id" binding:"required,gte=1"`
}

type ListRequest struct {
	Page     int `form:"page" binding:"required,gt=0"`
	PageSize int `form:"pageSize" binding:"required,gt=0"`
}
