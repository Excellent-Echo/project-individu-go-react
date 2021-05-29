package role

import (
	"errors"
	"fmt"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
)

type Service interface {
	GetAllRole() ([]RoleFormat, error)
	SaveNewRole(role entity.RoleInput) (RoleFormat, error)
	GetRoleByID(roleID string) (RoleFormat, error)
	GetandDeleteRoleByID(roleID string) (interface{}, error)
	GetandUpdateRoleByID(roleID string, dataRoleInput entity.RoleInputUpdate) (RoleFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllRole() ([]RoleFormat, error) {
	roles, err := s.repository.FindAllRoles()
	var formatRoles []RoleFormat

	for _, role := range roles {
		formatrole := FormatRole(role)
		formatRoles = append(formatRoles, formatrole)
	}

	if err != nil {
		return formatRoles, err
	}
	return formatRoles, nil
}

func (s *service) SaveNewRole(role entity.RoleInput) (RoleFormat, error) {

	var newRole = entity.Role{
		NamaRole:    role.NamaRole,
		Description: role.Description,
	}

	createRole, err := s.repository.CreateRole(newRole)
	formatRole := FormatRole(createRole)

	if err != nil {
		return formatRole, err
	}

	return formatRole, nil
}

func (s *service) GetRoleByID(roleID string) (RoleFormat, error) {
	if err := helper.ValidateIDNumber(roleID); err != nil {
		return RoleFormat{}, err
	}

	role, err := s.repository.FindRoleByID(roleID)

	if err != nil {
		return RoleFormat{}, err
	}

	if role.ID == 0 {
		newError := fmt.Sprintf("role id %s not found", roleID)
		return RoleFormat{}, errors.New(newError)
	}

	formatRole := FormatRole(role)

	return formatRole, nil
}

func (s *service) GetandDeleteRoleByID(roleID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(roleID); err != nil {
		return RoleFormat{}, err
	}

	role, err := s.repository.FindRoleByID(roleID)

	if err != nil {
		return nil, err
	}

	if role.ID == 0 {
		newError := fmt.Sprintf("role id %s not found", roleID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.FindandDeleteRoleByID(roleID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	message := fmt.Sprintf("success delete role by ID : %s", roleID)

	formatDelete := FormatDeleteRole(message)

	return formatDelete, nil
}

func (s *service) GetandUpdateRoleByID(roleID string, dataRoleInput entity.RoleInputUpdate) (RoleFormat, error) {

	var dataRoleUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(roleID); err != nil {
		return RoleFormat{}, err
	}

	role, err := s.repository.FindRoleByID(roleID)

	if err != nil {
		return RoleFormat{}, err
	}

	if role.ID == 0 {
		newError := fmt.Sprintf("role id not found : %s", roleID)
		return RoleFormat{}, errors.New(newError)
	}

	if dataRoleInput.NamaRole != "" || len(dataRoleInput.NamaRole) != 0 {
		dataRoleUpdate["nama_role"] = dataRoleInput.NamaRole
	}
	if dataRoleInput.Description != "" || len(dataRoleInput.Description) != 0 {
		dataRoleUpdate["description"] = dataRoleInput.Description
	}

	fmt.Println(dataRoleUpdate)

	roleUpdated, err := s.repository.FindAndUpdateRoleByID(roleID, dataRoleUpdate)

	if err != nil {
		return RoleFormat{}, err
	}

	formatRole := FormatRole(roleUpdated)

	return formatRole, nil
}
