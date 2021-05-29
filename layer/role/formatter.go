package role

import (
	"project-individu-go-react/entity"
	"time"
)

type RoleFormat struct {
	ID          int    `json:"id"`
	NamaRole    string `json:"nama_role"`
	Description string `json:"description"`
}

func FormatRole(role entity.Role) RoleFormat {
	var formatRole = RoleFormat{
		ID:          role.ID,
		NamaRole:    role.NamaRole,
		Description: role.Description,
	}
	return formatRole
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatDeleteRole(message string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    message,
		TimeDelete: time.Now(),
	}
	return deleteFormat
}
