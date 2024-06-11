package main

// DateTemperatureService interface represents a service that provides date and temperature information.
type DateTemperatureService interface {
	GetDate() string
	GetTemperature() float64
	Reset()
}

// SimpleDateTemperatureService is a basic implementation of DateTemperatureService.
type SimpleDateTemperatureService struct {
	date        string
	temperature float64
	observers   []Observer
}

// GetDate returns the current date.
func (s *SimpleDateTemperatureService) GetDate() string {
	return s.date
}

// GetTemperature returns the current temperature.
func (s *SimpleDateTemperatureService) GetTemperature() float64 {
	return s.temperature
}

// SetDate sets the current date and notifies observers.
func (s *SimpleDateTemperatureService) SetDate(date string) {
	s.date = date
	s.notifyObservers()
}

// SetTemperature sets the current temperature and notifies observers.
func (s *SimpleDateTemperatureService) SetTemperature(temperature float64) {
	s.temperature = temperature
	s.notifyObservers()
}

// Reset resets the service.
func (s *SimpleDateTemperatureService) Reset() {
	s.date = "2024-06-09"
	s.temperature = 25.5
	s.notifyObservers()
}

func (s *SimpleDateTemperatureService) AddObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *SimpleDateTemperatureService) RemoveObserver(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *SimpleDateTemperatureService) notifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s)
	}
}
