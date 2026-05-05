package dto

import (
	"encoding/json"
	"time"
)

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func NewErrorDTO(err error) *ErrorDTO {
	return &ErrorDTO{
		Message: err.Error(),
		Time: time.Now(),
	}
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ");
	if err != nil {
		panic(err);
	}

	return string(b);
}