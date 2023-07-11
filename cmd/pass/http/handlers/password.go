package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"passback/cmd/pass/generator"
	"passback/cmd/pass/http/model"
)

type PasswordHandler struct {
	generator *generator.Generator
}

func NewPasswordHandler(generator *generator.Generator) *PasswordHandler {
	return &PasswordHandler{
		generator: generator,
	}
}

func (h *PasswordHandler) Route() string {
	return "/pass"
}

func (h *PasswordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
	if r.Method == http.MethodOptions {
		return
	}

	request := model.GeneratePasswordRequest{}
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithError(err).Error("error on body read")

		return
	}

	err = request.UnmarshalJSON(raw)
	if err != nil {
		logrus.WithError(err).Error("error on request unmarshal")

		return
	}

	response := model.GeneratePasswordResponse{}

	task := generator.NewTask(generator.Config{
		Length:           request.Length,
		SymbolsUppercase: request.SymbolsUppercase,
		Numbers:          request.Numbers,
		SpecialSymbols:   request.SpecialSymbols,
	})

	h.generator.Generate(&task)
	response.Password = task.Done()

	resp, err := response.MarshalJSON()
	if err != nil {
		logrus.WithError(err).Error("error on marshal json")

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		logrus.WithError(err).Error("error on write to output")

		return
	}

	return
}

func (h *PasswordHandler) Methods() []string {
	return []string{http.MethodPost, http.MethodOptions}
}
