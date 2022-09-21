package repositories

type _UserInterface interface {
	Login(string, string) // (Username, Password)
}

type _UserRepository struct{}

func (repo _UserRepository) Login(username, password string) {

}

func RegisterUserRepository() _UserInterface {
	var repository _UserInterface = _UserRepository{}
	return repository
}
