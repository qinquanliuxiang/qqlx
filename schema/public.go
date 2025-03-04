package schema

type IDRequest struct {
	ID int `json:"id" form:"id" binding:"required,gte=1"`
}

type ListRequest struct {
	Page     int `form:"page" binding:"required,gt=0" json:"page"`
	PageSize int `form:"pageSize" binding:"required,gt=0" json:"pageSize"`
}
