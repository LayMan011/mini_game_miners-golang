package errors

import (
	"encoding/json"
	"errors"
	"myproj/data/dto"
	"net/http"
)

/// for equipment
var ErrEquipmentNotFound = errors.New("equipment not found");
var ErrEquipmentAlreadyPurchased = errors.New("equipment already buy");
var ErrEquipmentNotEnoughMoney = errors.New("Not enough money for equipment");

/// for miner
var ErrMinerNotFound = errors.New("miner not found");
var ErrClassMinerNotFound = errors.New("miner class not found");
var ErrMinerNotEnoughMoney = errors.New("Not enough money for miner");

/// for http
var ErrMinerAlreadyExists = errors.New("miner already buy");

func HttpErrorBadRequest(w http.ResponseWriter, err error) {
	errDTO := dto.NewErrorDTO(err);

	http.Error(w, errDTO.ToString(), http.StatusBadRequest);
}

func HttpErrorConflict(w http.ResponseWriter, err error) {
	errDTO := dto.NewErrorDTO(err);

	if errors.Is(err, ErrMinerAlreadyExists) {
		http.Error(w, errDTO.ToString(), http.StatusConflict);
	} else {
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError);
	}
}

func HttpErrorMinerNotFound(w http.ResponseWriter, err error) {
	errDTO := dto.NewErrorDTO(err);

	if errors.Is(err, ErrMinerNotFound) {
		http.Error(w, errDTO.ToString(), http.StatusNotFound);
	} else {
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError);
	}
}

func HttpErrorEquipmentBuy(w http.ResponseWriter, err error) {
	errDTO := dto.NewErrorDTO(err);

	if errors.Is(err, ErrEquipmentNotFound) {
		http.Error(w, errDTO.ToString(), http.StatusNotFound);
	} else if errors.Is(err, ErrEquipmentAlreadyPurchased) {
		http.Error(w, errDTO.ToString(), http.StatusBadRequest);
	} else {
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError);
	}
}

func JsonMarhalInd(book any) []byte {
	b, err := json.MarshalIndent(book, "", "    ");
	if err != nil {
		panic(err);
	}

	return b;
}