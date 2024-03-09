package model

type TransportType int

const (
	Walking TransportType = iota
	Bicycle
	Car
)

type TourDuration struct {
	Duration      int
	TransportType TransportType
}
