package sizedto

type SizeRequest struct {
	Size string `db:"column:size" json:"size" form:"size"`
}
type SizeColorRequest struct {
	ID string `db:"column:id" json:"id" form:"id"`
}
