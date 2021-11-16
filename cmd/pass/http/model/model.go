package model

//go:generate easyjson

//easyjson:json
type (
	GeneratePasswordRequest struct {
		Length           int  `json:"length"`
		SymbolsUppercase bool `json:"uppercase"`
		Numbers          bool `json:"numbers"`
		SpecialSymbols   bool `json:"special_symbols"`
	}

	GeneratePasswordResponse struct {
		Password string `json:"password"`
	}
)
