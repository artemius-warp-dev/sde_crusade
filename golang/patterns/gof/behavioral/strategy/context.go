package main

type Context struct {
	service  DateTemperatureService
	strategy FormatStrategy
}

func NewContext(service DateTemperatureService, strategy FormatStrategy) *Context {

	return &Context{
		service:  service,
		strategy: strategy,
	}

}

func (c *Context) SetStrategy(strategy FormatStrategy) {
	c.strategy = strategy
}

func (c *Context) GetFormattedData() string {

	date := c.service.GetDate()
	temperature := c.service.GetTemperature()
	return c.strategy.Format(date, temperature)

}
