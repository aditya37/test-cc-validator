package transport

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aditya37/test-cc-validator/model"
	"github.com/aditya37/test-cc-validator/repository"
	"github.com/aditya37/test-cc-validator/service"
	"github.com/aditya37/test-cc-validator/utils"
)

type Transport struct {
	service service.Iservice
}

func NewTransport(svc service.Iservice) *Transport {
	return &Transport{
		service: svc,
	}
}

func (t *Transport) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("oke"))
}

func (t *Transport) GetCardNumber(w http.ResponseWriter, r *http.Request) {
	// set header to json
	w.Header().Set("Content-Type", "application/json")

	// encode payload
	var payload model.RequestValidateCard
	json.NewDecoder(r.Body).Decode(&payload)

	// service or usecase
	resp, err := t.service.GetCardNumber(r.Context(), payload.CardNumber)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) || errors.Is(err, service.ErrUnknownCard) {
			utils.MappError(w, err, http.StatusFound)
			return
		}
		utils.MappError(w, err, http.StatusBadRequest)
		return
	}

	// convert to json
	buf, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

func (t *Transport) InsertCard(w http.ResponseWriter, r *http.Request) {
	// set header to json
	w.Header().Set("Content-Type", "application/json")

	// encode payload
	var payload model.RequestAddCard
	json.NewDecoder(r.Body).Decode(&payload)
	resp, err := t.service.InsertCardNumber(r.Context(), payload)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicate) {
			utils.MappError(w, err, http.StatusUnprocessableEntity)
			return
		}
		utils.MappError(w, err, http.StatusBadRequest)
		return
	}

	// convert to json
	buf, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusCreated)
	w.Write(buf)
}
