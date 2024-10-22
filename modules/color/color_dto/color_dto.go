package colordto

type ColorRequest struct {
	Name     string `db:"column:name" json:"name" form:"name"`
	Code     string `db:"column:code" json:"code" form:"code"`
	CodeName string `db:"column:code_name" json:"code_name" form:"code_name"`
}
type DeleteColorRequest struct {
	ID string `db:"column:id" json:"id" form:"id"`
}
