package project_usecase

type AddProjectInput struct {
	Name        string
	Description string
	Logo        []byte
	Background  []byte
	ServiceFee  uint
	Languages   []string
}

func NewAddProjectInput(
	name, description string,
	logo, background []byte,
	serviceFee uint,
	languages []string) *AddProjectInput {
	return &AddProjectInput{
		Name:        name,
		Description: description,
		Logo:        logo,
		Background:  background,
		ServiceFee:  serviceFee,
		Languages:   languages,
	}
}
