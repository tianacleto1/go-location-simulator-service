package route

type Route struct {
	ID        string
	ClientId  string
	Positions []Position
}

type Position struct {
	Lat  float64
	Long float64
}
