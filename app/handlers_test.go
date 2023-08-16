package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/YoungOak/GoAPI/internal/car"
	"github.com/YoungOak/GoAPI/internal/data"
)

var testRecord = car.Record{
	ID:       "123",
	Make:     "Toyota",
	Model:    "Camry",
	Category: "Sedan",
	Package:  "Standard",
	Color:    "Blue",
	Year:     time.Now().Year(),
	Mileage:  1000,
	Price:    10000,
}

func TestPOSTCars(t *testing.T) {
	CarManager = data.NewManager()
	body, _ := json.Marshal(testRecord)

	req, err := http.NewRequest(http.MethodPost, "/car", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(POSTCar)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusAccepted {
		t.Errorf("Expected response code %v, got: %v", http.StatusAccepted, rr.Code)
	}

	expectedResponse := fmt.Sprintf("added car '%s' to database", testRecord.ID)
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response: %v, got: %v", expectedResponse, rr.Body.String())
	}
}

func TestGETCar(t *testing.T) {
	CarManager = data.NewManager()
	_ = CarManager.Add(testRecord)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/car?id=%s", testRecord.ID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GETCar)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected response code %v, got %v", http.StatusOK, rr.Code)
	}

	var gotRecord car.Record
	err = json.Unmarshal(rr.Body.Bytes(), &gotRecord)
	if err != nil {
		t.Fatalf("Failed unmarshalling response: %v", err)
	}

	if !reflect.DeepEqual(gotRecord, testRecord) {
		t.Fatalf("Unexpected record obtained, wanted: %v, got: %v", testRecord, gotRecord)
	}
}

func TestGETCars(t *testing.T) {
	CarManager = data.NewManager()
	_ = CarManager.Add(testRecord)
	req, err := http.NewRequest(http.MethodGet, "/cars", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GETCars)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected response code %v, got %v", http.StatusOK, rr.Code)
	}

	var gotRecords []car.Record
	err = json.Unmarshal(rr.Body.Bytes(), &gotRecords)
	if err != nil {
		t.Fatalf("Failed unmarshalling response: %v", err)
	}

	wantRecords := []car.Record{testRecord}
	if !reflect.DeepEqual(gotRecords, wantRecords) {
		t.Fatalf("Unexpected record obtained, wanted: %v, got: %v", wantRecords, gotRecords)
	}
}

func TestPUTCar(t *testing.T) {
	CarManager = data.NewManager()
	_ = CarManager.Add(testRecord)

	updatedRecord := testRecord
	updatedRecord.Make = "Mitsubishi"

	body, _ := json.Marshal(updatedRecord)

	req, err := http.NewRequest(http.MethodPut, "/cars", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PUTCar)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusAccepted {
		t.Errorf("Expected response code %v, got %v", http.StatusAccepted, rr.Code)
	}
}
