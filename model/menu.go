package model

const (
	MenuTypeIndex = "index" // 前台用户菜单
	MenuTypeAdmin = "admin" // 管理后台菜单
)

type Menu struct {
	ID       uint    `gorm:"primarykey" json:"ID"`
	Key      string  `gorm:"uniqueIndex:idx_key_type;size:255;not null" json:"key"`
	Type     string  `gorm:"uniqueIndex:idx_key_type;size:64;not null" json:"type"`
	Icon     string  `gorm:"size:255;not null;default:''" json:"icon"`
	Parent   string  `gorm:"size:255;not null;default:''" json:"parent"`
	Label    string  `gorm:"size:64;not null" json:"label"`
	Order    int     `gorm:"not null;default:0" json:"order"`
	Children []*Menu `gorm:"-" json:"children,omitempty"`
}
