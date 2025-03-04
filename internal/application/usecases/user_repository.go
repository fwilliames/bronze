package usecases

type UserRepository interface {
	SaveUser(name string, age int) error
}
