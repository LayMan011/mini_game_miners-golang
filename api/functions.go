package api

import (
	"fmt"
	"myproj/errors"
	"net/http"
)

func WriteData(my_data any , w http.ResponseWriter) {
	b := errors.JsonMarhalInd(my_data)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err.Error())
		return
	}
}