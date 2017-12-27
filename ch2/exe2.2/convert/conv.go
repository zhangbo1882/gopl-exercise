package convert

//temperature
type Kelvin float64
type Celsius float64
type Fahrenhert float64

//length
type Meters float64
type Feet float64

//weight
type Pounds float64
type Kilograms float64

const (
	AbsoluteZeroC float64 = -273.15
	MetersPerFeet float64 = 0.3048
	KiloPerPounds float64 = 0.45359237
)

func CToF(c Celsius) Fahrenhert {
	return Fahrenhert(c*9/5 + 32)
}

func FToC(f Fahrenhert) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k + Kelvin(AbsoluteZeroC))
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c - Celsius(AbsoluteZeroC))
}

func FeetToMeters(feet Feet) Meters {
	return Meters(feet * Feet(MetersPerFeet))
}

func MetersToFeet(meters Meters) Feet {
	return Feet(meters / Meters(MetersPerFeet))
}

func PoundsToKilo(pounds Pounds) Kilograms {
	return Kilograms(pounds * Pounds(KiloPerPounds))
}

func KiloToPounds(kilo Kilograms) Pounds {
	return Pounds(kilo / Kilograms(KiloPerPounds))
}
