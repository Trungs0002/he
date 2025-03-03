package branch

type CreateInput struct {
	Name  string
	Alias string
	Code  string
}

type UpdateInput struct {
	ID    string
	Name  string
	Alias string
	Code  string
}