package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

const baseURL = "http://localhost:8080"

type CarRecord struct {
	ID       string `json:"id"`
	Make     string `json:"make"`
	Model    string `json:"model"`
	Category string `json:"category"`
	Package  string `json:"package"`
	Color    string `json:"color"`
	Year     int    `json:"year"`
	Mileage  int    `json:"mileage"`
	Price    int    `json:"price"`
}

func TestCarsAPIIntegration(t *testing.T) {
	// 1. POST a new car
	newCar := CarRecord{
		ID:       "test-car-1",
		Make:     "TestMake",
		Model:    "TestModel",
		Category: "Test",
		Package:  "TestPackage",
		Color:    "Blue",
		Year:     2022,
		Mileage:  0,
		Price:    25000,
	}

	data, err := json.Marshal(newCar)
	if err != nil {
		t.Fatalf("Failed to marshal car data: %v", err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/car", baseURL), "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("Failed to POST new car: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		t.Fatalf("Expected status code %d, got %d", http.StatusAccepted, resp.StatusCode)
	}

	// 2. GET the list of cars and check our car
	resp, err = http.Get(fmt.Sprintf("%s/cars", baseURL))
	if err != nil {
		t.Fatalf("Failed to GET cars: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var cars []CarRecord
	err = json.Unmarshal(body, &cars)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(cars) == 0 {
		t.Fatal("No cars found")
	}

	found := false
	for _, car := range cars {
		if car.ID == newCar.ID {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("Newly added car with ID %s not found", newCar.ID)
	}

	// 3. GET specific car by ID
	resp, err = http.Get(fmt.Sprintf("%s/car?id=%s", baseURL, newCar.ID))
	if err != nil {
		t.Fatalf("Failed to GET car by ID: %v", err)
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	var car CarRecord
	err = json.Unmarshal(body, &car)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if car.ID != newCar.ID {
		t.Fatalf("Expected car ID %s, got %s", newCar.ID, car.ID)
	}

	// 4. PUT to update car and validate
	newCar.Color = "Red"
	data, err = json.Marshal(newCar)
	if err != nil {
		t.Fatalf("Failed to marshal car data for update: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/car", baseURL), bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("Failed to create PUT request: %v", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to PUT update for car: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		t.Fatalf("Expected status code %d for update, got %d", http.StatusAccepted, resp.StatusCode)
	}
}

func TestMain(m *testing.M) {
	if os.Getenv("INTEGRATION") == "" {
		// Skip Integration Tests
		os.Exit(0)
	}
	// Run Integration Tests
	os.Exit(m.Run())
}
