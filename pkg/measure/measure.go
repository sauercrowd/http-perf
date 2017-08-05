package measure

type Measurement interface {
	GetChart() ([]byte, error)
	GetAVG() float64
}
