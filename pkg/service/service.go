package service

import (
	"main.go/pkg/repository"
)

type Service struct {
	Translator
}

type Translator interface {
	Translate(string) ([]string, error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Translator: NewTranslateService(repos.Translator),
	}
}