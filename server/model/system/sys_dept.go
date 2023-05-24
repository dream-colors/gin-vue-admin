package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysDepartment struct {
	global.GVA_MODEL
	ID       uint   `gorm:"primary_key"`
	Name     string `json:"name" gorm:"not null;comment:部门名称"`             // 部门名称
	ParentID uint   `json:"parentID" gorm:"column:parent_id;comment:父级id"` // 父部门ID
	Sort     uint   `json:"sort" gorm:"comment:排序"`                        // 部门排序
	Head     string `json:"head" gorm:"comment:负责人"`                       // 负责人
	Email    string `json:"email" gorm:"comment:邮箱"`                       // 邮箱
	Phone    string `json:"phone" gorm:"comment:联系电话"`                     // 联系电话
	Desc     string `json:"desc" gorm:"comment:描述"`                        // 部门描述
}

func (SysDepartment) TableName() string {
	return "sys_dept"
}
