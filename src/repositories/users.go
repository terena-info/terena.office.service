package repositories

type UserInterface interface {
	Login(string, string) // (Username, Password)
	GenerateAccessToken() string
}

type _UserRepository struct{}

func (repo _UserRepository) Login(username, password string) {

}

func (repo _UserRepository) GenerateAccessToken() string {
	return ""
}

func RegisterUserRepository() UserInterface {
	var repository UserInterface = _UserRepository{}
	return repository
}
