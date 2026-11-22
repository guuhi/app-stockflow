package usecase

type ExampleRepository interface {
	// Define methods for the repository here
}

type ExampleUsecase struct {
	repo ExampleRepository
}

func NewExampleUsecase(repo ExampleRepository) *ExampleUsecase {
	return &ExampleUsecase{repo: repo}
}
