package temperature

type Temperature float64

func CelsiusToFahrenheit(t Temperature) Temperature {
	return Temperature((t * 9 / 5) + 32)
}

func FahrenheitToCelsius(t Temperature) Temperature {
	return Temperature((t - 32) * 5 / 9)
}
