package repository

type Translator interface {
	GetWords() (string, error)
}

type Repository struct {
	Translator
}

func NewRepository(file string) *Repository {
	return &Repository{
		Translator: NewFileTranslate(file),
	}
}
