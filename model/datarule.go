package model

import "time"

// DataRule 数据权限表规则
type DataRule struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Group     string `gorm:"index:idx_group;size:255;not null" json:"group"`                      // 所属用户组
	Table     string `gorm:"index:idx_table;size:255;not null" json:"table"`                      // 所属数据表
	RuleName  string `gorm:"size:255;not null" json:"rule_name"`                                  // 规则名称
	Relation  uint8  `gorm:"default:0;not null;comment:0-AND,1-OR" json:"relation"`               // 表内逻辑关系
	Status    uint8  `gorm:"index:idx_status;default:0;not null;comment:0-启用,1-禁用" json:"status"` // 状态
	CreatedAt time.Time
	UpdatedAt time.Time
}

// FieldRule 数据权限字段规则
type FieldRule struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	DataRuleId uint   `gorm:"index:idx_rule_id;not null" json:"data_rule_id"`        // 所属规则ID
	Field      string `gorm:"size:255;not null" json:"field"`                        // 字段名
	Method     string `gorm:"size:64;not null" json:"method"`                        // 比较方式
	Value      string `gorm:"size:10240;not null" json:"value"`                      // 比较值
	Relation   uint8  `gorm:"default:0;not null;comment:0-AND,1-OR" json:"relation"` // 字段间逻辑关系
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
