package main

type CompositeDateTemperatureService struct {
	services []DateTemperatureService
}

func NewCompositeDateTemperatureService() *CompositeDateTemperatureService {

	return &CompositeDateTemperatureService{}

}

func (c *CompositeDateTemperatureService) Add(service DateTemperatureService) {
	c.services = append(c.services, service)
}

// GetDate returns the date from the first service in the composite.
func (c *CompositeDateTemperatureService) GetDate() string {
	if len(c.services) == 0 {
		return ""
	}
	return c.services[0].GetDate()
}

// GetTemperature returns the average temperature of all services in the composite.
func (c *CompositeDateTemperatureService) GetTemperature() float64 {
	if len(c.services) == 0 {
		return 0
	}

	totalTemperature := 0.0
	for _, service := range c.services {
		totalTemperature += service.GetTemperature()
	}

	return totalTemperature / float64(len(c.services))
}
