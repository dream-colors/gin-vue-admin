package system

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type DepartmentService struct {
}

func (departmentService *DepartmentService) AddDepartment(department system.SysDepartment) error {
	if !errors.Is(global.GVA_DB.Where("name = ?", department.Name).First(&system.SysDepartment{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("部门名称已存在，请修改名称")
	}
	return global.GVA_DB.Create(department).Error
}

func (departmentService *DepartmentService) DeleteDepartment(departmentId int) error {
	var department system.SysDepartment
	return global.GVA_DB.Where("id = ?", departmentId).Delete(&department).Error
}

func (departmentService *DepartmentService) UpdateDepartment(d system.SysDepartment) error {
	var department system.SysDepartment
	if err := global.GVA_DB.Where("id = ?", d.ID).First(&department).Error; err != nil {
		return err
	}

	department.Name = d.Name
	department.ParentID = d.ParentID
	department.Sort = d.Sort
	department.Head = d.Head
	department.Email = d.Email
	department.Phone = d.Phone
	department.Desc = d.Desc

	return global.GVA_DB.Save(&department).Error
}
