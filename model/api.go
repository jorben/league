package model

type Api struct {
	ID      uint   `gorm:"primarykey" json:"ID"`
	Name    string `gorm:"size:255;not null;default:''" json:"name"`                              // 接口名称
	Path    string `gorm:"size:255;uniqueIndex:idx_path_method;not null;default:''" json:"path"`  // 接口地址
	Method  string `gorm:"size:64;uniqueIndex:idx_path_method;not null;default:''" json:"method"` // 请求方法
	Comment string `gorm:"size:255;not null;default:''" json:"comment"`                           // 备注
}

type ApiList struct {
	Count int64  //总记录数
	List  []*Api // Api信息
}
