package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/YoungOak/GoAPI/internal/car"
	"github.com/YoungOak/GoAPI/internal/data"
)

func carsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GETCars(w, r)
	default:
		methodNotAllowedError(w, r)
	}
}

func carHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		POSTCar(w, r)
	case http.MethodGet:
		GETCar(w, r)
	case http.MethodPut:
		PUTCar(w, r)
	default:
		methodNotAllowedError(w, r)
	}
}

func POSTCar(w http.ResponseWriter, r *http.Request) {
	var record car.Record

	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		slog.WarnContext(r.Context(), "error decoding body: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error decoding body: %s", err.Error())))
		return
	}

	err = CarManager.Add(record)
	if err != nil {
		_, invalid := err.(car.ErrorFieldInvalid)
		_, missing := err.(car.ErrorFieldMissing)
		_, alreadyExists := err.(data.ErrorAlreadyExists)
		if invalid || missing {
			slog.WarnContext(r.Context(), err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else if alreadyExists {
			slog.WarnContext(r.Context(), err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else {
			slog.ErrorContext(r.Context(), fmt.Sprintf("error getting car: %s", err.Error()))
			internalServerError(w, r)
		}
		return
	}

	slog.Info(fmt.Sprintf("added new car with id: '%s'", record.ID))
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(fmt.Sprintf("added car '%s' to database", record.ID)))
}

func GETCars(w http.ResponseWriter, r *http.Request) {
	records := CarManager.List()

	slog.Info(fmt.Sprintf("listing all %v cars", len(records)))
	jsonRecords, err := json.Marshal(records)
	if err != nil {
		slog.ErrorContext(r.Context(), "error marshalling records: %w", err)
		internalServerError(w, r)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-type", "application/json")
	w.Write(jsonRecords)
}

func GETCar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	record, err := CarManager.Get(id)
	if err != nil {
		_, invalid := err.(car.ErrorFieldInvalid)
		_, missing := err.(car.ErrorFieldMissing)
		_, notFound := err.(data.ErrorRecordNotFound)
		if invalid || missing {
			slog.WarnContext(r.Context(), err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else if notFound {
			slog.WarnContext(r.Context(), err.Error())
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		} else {
			slog.ErrorContext(r.Context(), fmt.Sprintf("error getting car: %s", err.Error()))
			internalServerError(w, r)
		}
		return
	}

	slog.Info(fmt.Sprintf("found car with id: '%s'", id))
	jsonRecord, err := json.Marshal(record)
	if err != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error marshalling record: %s", err.Error()))
		internalServerError(w, r)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-type", "application/json")
	w.Write(jsonRecord)
}

func PUTCar(w http.ResponseWriter, r *http.Request) {
	var record car.Record

	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		slog.WarnContext(r.Context(), fmt.Sprintf("error decoding body: %s", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error decoding body: %s", err.Error())))
		return
	}

	err = CarManager.Update(record)
	if err != nil {
		_, invalid := err.(car.ErrorFieldInvalid)
		_, missing := err.(car.ErrorFieldMissing)
		_, notFound := err.(data.ErrorRecordNotFound)
		if invalid || missing {
			slog.WarnContext(r.Context(), err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else if notFound {
			slog.WarnContext(r.Context(), err.Error())
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		} else {
			slog.ErrorContext(r.Context(), fmt.Sprintf("error getting car: %s", err.Error()))
			internalServerError(w, r)
		}
		return
	}

	slog.Info(fmt.Sprintf("updated car with id: '%s'", record.ID))
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(fmt.Sprintf("updated car '%s'", record.ID)))
}

func methodNotAllowedError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(fmt.Sprintf("Method %s not allowed", r.Method)))
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("unexpected internal error, please retry later"))
}
