package temperature

import (
	"testing"
)

const checkmark = "\u2713"
const ballotX = "\u2717"

var temperatures =[]struct {
	celsius float64
	fahrenheit float64	
}{ {29.0, 84.2}, {15.0, 49.0}, {25.0, 77.0} }

func TestCelsiusToFahrenheit(t *testing.T) {
	ctemp := 29.0
	ftemp := 84.2

	if (CelsiusToFahrenheit(ctemp) != ftemp) {
		t.Fatal("\nError on converting from Celsius to Fahrenheit", ballotX)
	} else {
		t.Log("\nConverting from Celsius to Fahrenheit: correct", checkmark)
	}
}

func TestMCelsiusToFahrenheit(t *testing.T) {
	t.Parallel()
	for _, temp := range temperatures {
		result := CelsiusToFahrenheit(temp.celsius)
		expected := temp.fahrenheit
		if (result != expected) {
			t.Errorf("\nError on converting from Celsius to Fahrenheit %s\nExpected %.2f, result %.2f", ballotX, expected, result)
		} else {
			t.Log("\nConverting from Celsius to Fahrenheit: correct", checkmark)
		}
	}
}