package usecase

type usersUseCase interface {
	Get() string
	Post() string
	Put() string
	Delete() string
}

type usersUseCase struct {
	// Add your dependencies here
}

func NewusersUseCase() usersUseCase {
	return &usersUseCase{}
}

func (u *usersUseCase) Get() string {
	// Implement your logic here
	return "Get users"
}

func (u *usersUseCase) Post() string {
	// Implement your logic here
	return "Post users"
}

func (u *usersUseCase) Put() string {
	// Implement your logic here
	return "Put users"
}

func (u *usersUseCase) Delete() string {
	// Implement your logic here
	return "Delete users"
}
