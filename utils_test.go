package main

import "testing"

func TestConfigUnitToAPIUnit(t *testing.T) {
	tests := []struct {
		ConfigUnit string
		APIUnit    string
	}{
		{TempCelsius, "metric"},
		{TempFahrenheit, "imperial"},
		{TempKelvin, ""},
	}

	for _, test := range tests {
		t.Run(test.ConfigUnit, func(t *testing.T) {
			got := configUnitToAPIUnit(test.ConfigUnit)

			if got != test.APIUnit {
				t.Errorf("got %v, want %v", got, test.APIUnit)
			}
		})
	}
}
