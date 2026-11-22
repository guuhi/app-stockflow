package handler

type ExampleUsecase interface {
	// Define methods for the usecase here
}

type ExampleHandler struct {
	uc ExampleUsecase
}

func NewExampleHandler(uc ExampleUsecase) *ExampleHandler {
	return &ExampleHandler{uc: uc}
}