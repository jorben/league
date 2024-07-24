package model

// Increase the column size to 512.
type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:512;uniqueIndex:unique_index"`
	V0    string `gorm:"size:512;uniqueIndex:unique_index"`
	V1    string `gorm:"size:512;uniqueIndex:unique_index"`
	V2    string `gorm:"size:512;uniqueIndex:unique_index"`
	V3    string `gorm:"size:512;uniqueIndex:unique_index"`
	V4    string `gorm:"size:512;uniqueIndex:unique_index"`
	V5    string `gorm:"size:512;uniqueIndex:unique_index"`
}

// CasbinReq 原库使用sub,obj,act比较抽象，具象一下
type CasbinReq struct {
	UserId string // the user that wants to access a resource.
	Path   string // the resource that is going to be accessed.
	Method string // the operation that the user performs on the resource.
}

type Policy struct {
	ID       uint   `json:"ID"`
	Subject  string `json:"subject"`   // 用户或用户组
	Path     string `json:"path"`      // 路径
	PathName string `json:"path_name"` // 路径名称
	Method   string `json:"method"`    // 请求方法
	Result   string `json:"result"`    // 判定结论 allow or deny
}
