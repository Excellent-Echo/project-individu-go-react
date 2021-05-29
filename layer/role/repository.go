package role

import (
	"gorm.io/gorm"
	"project-individu-go-react/entity"
)

type Repository interface {
	FindAllRoles() ([]entity.Role, error)
	CreateRole(role entity.Role) (entity.Role, error)
	FindRoleByID(ID string) (entity.Role, error)
	FindandDeleteRoleByID(ID string) (string, error)
	FindAndUpdateRoleByID(ID string, roleUpdate map[string]interface{}) (entity.Role, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllRoles() ([]entity.Role, error) {
	var roles []entity.Role
	if err := r.db.Find(&roles).Error; err != nil {
		return roles, err
	}
	return roles, nil
}

func (r *repository) CreateRole(role entity.Role) (entity.Role, error) {
	if err := r.db.Create(&role).Error; err != nil {
		return role, nil
	}
	return role, nil
}

func (r *repository) FindRoleByID(ID string) (entity.Role, error) {
	var roleFind entity.Role
	if err := r.db.Where("id = ?", ID).Find(&roleFind).Error; err != nil {
		return roleFind, err
	}
	return roleFind, nil
}

func (r *repository) FindandDeleteRoleByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.Role{}).Error; err != nil {
		return "error", err
	}
	return "success", nil
}

func (r *repository) FindAndUpdateRoleByID(ID string, roleUpdate map[string]interface{}) (entity.Role, error) {
	var roleWillUpdate entity.Role

	if err := r.db.Model(&roleWillUpdate).Where("id = ?", ID).Updates(roleUpdate).Error; err != nil {
		return roleWillUpdate, err
	}

	if err := r.db.Where("id = ?", ID).Find(&roleWillUpdate).Error; err != nil {
		return roleWillUpdate, err
	}

	return roleWillUpdate, nil
}
