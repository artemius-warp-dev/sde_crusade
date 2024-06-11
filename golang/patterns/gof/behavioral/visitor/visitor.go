package main

import "fmt"

// Visitor interface declares visiting methods for each concrete element.
type Visitor interface {
	VisitSimpleDateTemperatureService(*SimpleDateTemperatureService)
	VisitCustomDateTemperatureService(*CustomDateTemperatureService)
}

// ReportVisitor is a concrete visitor that generates reports.
type ReportVisitor struct{}

func (v *ReportVisitor) VisitSimpleDateTemperatureService(s *SimpleDateTemperatureService) {
	report := fmt.Sprintf("Simple Date: %s, Temperature: %.1f", s.date, s.temperature)
	fmt.Println(report)
}

func (v *ReportVisitor) VisitCustomDateTemperatureService(c *CustomDateTemperatureService) {
	report := fmt.Sprintf("Custom Temperature Report: %s - %.1f°C", c.date, c.temperature)
	fmt.Println(report)
}

// TemperatureConverterVisitor is another concrete visitor that converts temperature to Fahrenheit.
type TemperatureConverterVisitor struct{}

func (v *TemperatureConverterVisitor) VisitSimpleDateTemperatureService(s *SimpleDateTemperatureService) {
	fahrenheit := (s.temperature * 9 / 5) + 32
	fmt.Printf("Simple Date: %s, Temperature: %.1f°F\n", s.date, fahrenheit)
}

func (v *TemperatureConverterVisitor) VisitCustomDateTemperatureService(c *CustomDateTemperatureService) {
	fahrenheit := (c.temperature * 9 / 5) + 32
	fmt.Printf("Custom Temperature Report: %s - %.1f°F\n", c.date, fahrenheit)
}

// DetailVisitor is a concrete visitor that provides detailed reports.
type DetailVisitor struct{}

func (v *DetailVisitor) VisitSimpleDateTemperatureService(s *SimpleDateTemperatureService) {
	details := fmt.Sprintf("Detailed Report: On %s, the temperature recorded was %.1f°C. It was a pleasant day.", s.date, s.temperature)
	fmt.Println(details)
}

func (v *DetailVisitor) VisitCustomDateTemperatureService(c *CustomDateTemperatureService) {
	details := fmt.Sprintf("Detailed Temperature Report: On %s, the temperature was %.1f°C, recorded with precision. It was a notable day.", c.date, c.temperature)
	fmt.Println(details)
}
