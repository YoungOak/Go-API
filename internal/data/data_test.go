package data

import (
	"reflect"
	"slices"
	"testing"
	"time"

	"github.com/YoungOak/GoAPI/internal/car"
)

func TestManager_Add(t *testing.T) {
	testManager := NewManager()

	validRecord := car.Record{
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

	invalidRecord := validRecord
	invalidRecord.ID = ""

	tests := []struct {
		name    string
		record  car.Record
		wantErr error
	}{
		// Valid
		{
			name:    "Add valid record",
			record:  validRecord,
			wantErr: nil,
		},
		// Errors
		{
			name:    "Duplicate record ID",
			record:  validRecord,
			wantErr: ErrorAlreadyExists{"123"},
		},
		{
			name:    "Invalid record",
			record:  invalidRecord,
			wantErr: car.ErrorFieldMissing{Field: "ID"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testManager.Add(tt.record)
			if err != nil {
				if tt.wantErr == nil {
					t.Fatalf("unexpected error, wanted success, got: %v", err)
				}
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("unexpected error, wanted: %v, got: %v", tt.wantErr, err)
				}
			} else if tt.wantErr != nil {
				t.Fatalf("unexpected success, expected error: %s", tt.wantErr.Error())
			}
			if m, _ := testManager.(*manager); !reflect.DeepEqual(m.records[validRecord.ID], validRecord) {
				t.Fatalf("unexpected records in manager, expected to find one record: %v, but got: %v", validRecord, m.records)
			}
		})
	}
}

func TestManager_Get(t *testing.T) {
	testManager := NewManager()

	record := car.Record{
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

	_ = testManager.Add(record)

	tests := []struct {
		name       string
		id         string
		wantRecord car.Record
		wantErr    error
	}{
		// Valid
		{
			name:       "Get valid record",
			id:         "123",
			wantRecord: record,
			wantErr:    nil,
		},
		// Error
		{
			name:    "Record not found",
			id:      "456",
			wantErr: ErrorRecordNotFound{"456"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecord, gotErr := testManager.Get(tt.id)
			if gotErr != nil {
				if tt.wantErr == nil {
					t.Fatalf("unexpected error, wanted success, got: %v", gotErr)
				}
				if gotErr.Error() != tt.wantErr.Error() {
					t.Fatalf("unexpected error, wanted: %v, got: %v", tt.wantErr, gotErr)
				}
			} else if tt.wantErr != nil {
				t.Fatalf("unexpected success, expected error: %s", tt.wantErr.Error())
			}
			if gotErr == nil && !reflect.DeepEqual(gotRecord, record) {
				t.Fatalf("unexpected record obtained, expected: %v, got: %v", record, gotRecord)
			}
		})
	}
}

func TestManager_List(t *testing.T) {
	testManager := NewManager()

	// Setup some records
	record1 := car.Record{
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

	record2 := car.Record{
		ID:       "124",
		Make:     "Honda",
		Model:    "Civic",
		Category: "Coupe",
		Package:  "Premium",
		Color:    "Red",
		Year:     time.Now().Year(),
		Mileage:  500,
		Price:    12000,
	}

	_ = testManager.Add(record1)
	_ = testManager.Add(record2)

	tests := []struct {
		name     string
		wantList []car.Record
	}{
		{
			name:     "List records",
			wantList: []car.Record{record1, record2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList := testManager.List()
			for _, record := range tt.wantList {
				if !slices.Contains[[]car.Record, car.Record](gotList, record) {
					t.Fatalf("unexpected list, wanted: %v, got: %v", tt.wantList, gotList)
				}
			}
		})
	}
}

func TestManager_Update(t *testing.T) {
	testManager := NewManager()

	record := car.Record{
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

	_ = testManager.Add(record)

	updatedRecord := record
	updatedRecord.Make = "Honda"

	nonexistentRecord := record
	nonexistentRecord.ID = "456"

	tests := []struct {
		name    string
		record  car.Record
		wantErr error
	}{
		{
			name:    "Update valid record",
			record:  updatedRecord,
			wantErr: nil,
		},
		{
			name:    "Record not found",
			record:  nonexistentRecord,
			wantErr: ErrorRecordNotFound{"456"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testManager.Update(tt.record)
			if err != nil {
				if tt.wantErr == nil {
					t.Fatalf("unexpected error, wanted success, got: %v", err)
				}
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("unexpected error, wanted: %v, got: %v", tt.wantErr, err)
				}
			} else if tt.wantErr != nil {
				t.Fatalf("unexpected success, expected error: %s", tt.wantErr.Error())
			}
			if m, _ := testManager.(*manager); err == nil && reflect.DeepEqual(m.records[record.ID], record) {
				t.Fatalf("unexpected record, expected: %v, but got: %v", record, m.records[record.ID])
			}
		})
	}
}
