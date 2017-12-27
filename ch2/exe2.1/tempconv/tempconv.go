package tempconv

type Kelvin float64
type Celsius float64

const (
	AbsoluteZeroC float64 = -273.15
)

func KToC(k Kelvin) Celsius {
	return Celsius(k + Kelvin(AbsoluteZeroC))
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c - Celsius(AbsoluteZeroC))
}
