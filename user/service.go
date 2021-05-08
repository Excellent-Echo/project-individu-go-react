package user

type Service interface {
	GetAllUser() ([]UserFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllUser() ([]UserFormat, error) {
	// bisnis logic
	users, err := s.repository.FindAll()

	var formatUsers []UserFormat

	for _, user := range users {
		formatUser := FormatUser(user)
		formatUsers = append(formatUsers, formatUser)
	}

	if err != nil {
		return formatUsers, err
	}

	return formatUsers, nil
}
