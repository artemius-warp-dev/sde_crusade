package main

// DateTemperatureService interface represents a service that provides date and temperature information.
type DateTemperatureService interface {
	GetDate() string
	GetTemperature() float64
	Reset()
}

// DateTemperatureServiceContext is the context that holds the current state.
type DateTemperatureServiceContext struct {
	currentState State
}

// NewDateTemperatureServiceContext creates a new DateTemperatureServiceContext with the initial state.
func NewDateTemperatureServiceContext(initialState State) *DateTemperatureServiceContext {
	return &DateTemperatureServiceContext{currentState: initialState}
}

// SetState sets the current state of the service.
func (c *DateTemperatureServiceContext) SetState(state State) {
	c.currentState = state
}

// GetDate delegates the call to the current state's GetDate method.
func (c *DateTemperatureServiceContext) GetDate() string {
	return c.currentState.GetDate(c)
}

// GetTemperature delegates the call to the current state's GetTemperature method.
func (c *DateTemperatureServiceContext) GetTemperature() float64 {
	return c.currentState.GetTemperature(c)
}

// Reset delegates the call to the current state's Reset method.
func (c *DateTemperatureServiceContext) Reset() {
	c.currentState.Reset(c)
}
