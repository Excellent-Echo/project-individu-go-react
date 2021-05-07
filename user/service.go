package user

type Service interface {
	GetAllUsers() ([]UserFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllUsers() ([]UserFormat, error) {
	Users, err := s.repository.GetAll()

	var usersFormat []UserFormat

	for _, user := range Users {
		var userFormat = FormattingUser(user)
		usersFormat = append(usersFormat, userFormat)
	}

	if err != nil {
		return usersFormat, err
	}

	return usersFormat, nil
}
