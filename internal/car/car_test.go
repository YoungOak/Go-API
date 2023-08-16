package car

import (
	"testing"
	"time"
)

func TestCar_Validate(t *testing.T) {
	currentYear := time.Now().Year()
	validRecord := Record{
		ID:       "123",
		Make:     "Toyota",
		Model:    "Camry",
		Category: "Sedan",
		Package:  "Standard",
		Color:    "Blue",
		Year:     currentYear,
		Mileage:  1000,
		Price:    10000,
	}

	missingIDRecord := validRecord
	missingIDRecord.ID = ""

	missingMakeRecord := validRecord
	missingMakeRecord.Make = ""

	missingModelRecord := validRecord
	missingModelRecord.Model = ""

	missingCategoryRecord := validRecord
	missingCategoryRecord.Category = ""

	missingPackageRecord := validRecord
	missingPackageRecord.Package = ""

	missingColorRecord := validRecord
	missingColorRecord.Color = ""

	yearLessThan1900Record := validRecord
	yearLessThan1900Record.Year = 1899

	yearGreaterThanCurrentRecord := validRecord
	yearGreaterThanCurrentRecord.Year = currentYear + 1

	negativeMileageRecord := validRecord
	negativeMileageRecord.Mileage = -1

	zeroPriceRecord := validRecord
	zeroPriceRecord.Price = 0

	negativePriceRecord := validRecord
	negativePriceRecord.Price = -1

	tests := []struct {
		name    string
		record  Record
		wantErr error
	}{
		// Valid
		{
			name:    "valid record",
			record:  validRecord,
			wantErr: nil,
		},
		// Missing field
		{
			name:    "missing ID",
			record:  missingIDRecord,
			wantErr: ErrorFieldMissing{"ID"},
		},
		{
			name:    "missing Make",
			record:  missingMakeRecord,
			wantErr: ErrorFieldMissing{"Make"},
		},
		{
			name:    "missing Model",
			record:  missingModelRecord,
			wantErr: ErrorFieldMissing{"Model"},
		},
		{
			name:    "missing Category",
			record:  missingCategoryRecord,
			wantErr: ErrorFieldMissing{"Category"},
		},
		{
			name:    "missing Package",
			record:  missingPackageRecord,
			wantErr: ErrorFieldMissing{"Package"},
		},
		{
			name:    "missing Color",
			record:  missingColorRecord,
			wantErr: ErrorFieldMissing{"Color"},
		},
		// Invalid field scenarios
		{
			name:    "year less than 1900",
			record:  yearLessThan1900Record,
			wantErr: ErrorFieldInvalid{"Year", 1899},
		},
		{
			name:    "year greater than current year",
			record:  yearGreaterThanCurrentRecord,
			wantErr: ErrorFieldInvalid{"Year", currentYear + 1},
		},
		{
			name:    "negative Mileage",
			record:  negativeMileageRecord,
			wantErr: ErrorFieldInvalid{"Mileage", -1},
		},
		{
			name:    "zero Price",
			record:  zeroPriceRecord,
			wantErr: ErrorFieldInvalid{"Price", 0},
		},
		{
			name:    "negative Price",
			record:  negativePriceRecord,
			wantErr: ErrorFieldInvalid{"Price", -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.record.Validate()
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
		})
	}
}
