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

func (ts *TranslateService) Translate(word string) ([]string, error) {

	allWords, err := ts.repo.GetWords()
	if err != nil {
		return nil, err
	}
	translate := FindWord(word, allWords)

	return translate, nil
}

func HtmTranslate(word string) ([]string, error) {
	return make([]string, 0), nil
}

func PdfTranslate(word string) ([]string, error) {
	return make([]string, 0), nil
}

func SplitTextToArray(text string) []string {
	bodyText := strings.Split(text, "body")[1]
	splitText := strings.Split(bodyText, "<p>")

	slice := []string{"start"}
	index := 0

	for i := 0; i < len(splitText); i++ {
		if strings.Contains(splitText[i], "style=\"font-weight:bold;\"") {
			index++
			slice = append(slice, splitText[i])
			continue
		}
		slice[index] += splitText[i]
	}
	return slice
}

func GetSplitTranslate(str string) []string {
	res := strings.Split(str, "</p>")

	return res
}

func FindWord(word string, str string) []string {
	slice := SplitTextToArray(str)

	res := make([]string, 0)
	for i := 0; i < len(slice); i++ {
		split := GetSplitTranslate(slice[i])
		if strings.Contains(split[0], word) {
			res = append(res, strings.ReplaceAll(slice[i], "</p>", ""))
		}
	}

	return res
}
