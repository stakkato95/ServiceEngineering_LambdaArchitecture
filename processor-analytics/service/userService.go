package service

// "github.com/stakkato95/lambda-architecture/ingress/domain"
// "github.com/stakkato95/lambda-architecture/ingress/errs"

type UserService interface {
	GetUserCount() int
}

type simpleUserService struct {
	// repo domain.UserRepository
}

func (s *simpleUserService) GetUserCount() int {
	return 100500 //s.repo.InjestUser(user)
}

func NewSimpleUserService( /*repo domain.UserRepository*/ ) UserService {
	return &simpleUserService{ /*repo*/ }
}
