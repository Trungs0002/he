package branch

type CreateOptions struct {
	Name  string
	Alias string
	Code  string
}

type UpdateOptions struct {
	ID    string
	Name  string
	Alias string
	Code  string
}
