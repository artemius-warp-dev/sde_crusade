package main

// DateTemperatureService is the element interface that accepts a visitor.
type DateTemperatureService interface {
	Accept(visitor Visitor)
}

// SimpleDateTemperatureService is a concrete element.
type SimpleDateTemperatureService struct {
	date        string
	temperature float64
}

func (s *SimpleDateTemperatureService) Accept(visitor Visitor) {
	visitor.VisitSimpleDateTemperatureService(s)
}

// CustomDateTemperatureService is another concrete element.
type CustomDateTemperatureService struct {
	date        string
	temperature float64
}

func (c *CustomDateTemperatureService) Accept(visitor Visitor) {
	visitor.VisitCustomDateTemperatureService(c)
}
