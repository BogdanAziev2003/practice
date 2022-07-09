package service

import (
	"strings"

	"main.go/pkg/repository"
)

type TranslateService struct {
	repo repository.Translator
}

func NewTranslateService(repo repository.Translator) *TranslateService {
	return &TranslateService{
		repo: repo,
	}
}

func (ts *TranslateService) TranslateWord(word string) ([]string, error) {

	allWords, err := ts.repo.GetWords()
	if err != nil {
		return nil, err
	}
	translate := FindWord(word, allWords)

	return translate, nil
}

func (ts *TranslateService) GetAllWords() ([]string, error) {
	allWords, err := ts.repo.GetWords()
	if err != nil {
		return nil, err
	}
	return SplitTextToArray(allWords), nil
}

func SplitTextToArray(text string) []string {
	bodyText := strings.Split(text, "body")[1]
	splitText := strings.Split(bodyText, "<p>")

	slice := []string{"start"}
	index := 0

	for i := 0; i < len(splitText); i++ {
		if strings.Contains(splitText[i], "style=\"font-weight:bold;\"") {
			index++
			slice = append(slice, "<p>"+splitText[i])
			continue
		}
		slice[index] += "<p>" + splitText[i]
	}
	return slice
}

func GetSplitTranslate(str string) []string {

	return strings.Split(str, "</p>")
}

func FindWord(word string, str string) []string {
	slice := SplitTextToArray(str)
	res := make([]string, 0)
	for i := 0; i < len(slice); i++ {
		split := strings.Split(slice[i], "</p>") //GetSplitTranslate(slice[i])
		if strings.Contains(split[0], word) {
			res = append(res, slice[i]) //strings.ReplaceAll(slice[i], `'`, ""))
		}
	}

	return res
}
