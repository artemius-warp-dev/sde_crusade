package main

func main() {

	// Create concrete elements
	simpleService := &SimpleDateTemperatureService{
		date:        "2024-06-09",
		temperature: 25.5,
	}

	customService := &CustomDateTemperatureService{
		date:        "2024-06-09",
		temperature: 25.5,
	}

	// Create visitors
	reportVisitor := &ReportVisitor{}
	temperatureConverterVisitor := &TemperatureConverterVisitor{}
	detailVisitor := &DetailVisitor{}

	// Use the visitors
	simpleService.Accept(reportVisitor)
	customService.Accept(reportVisitor)

	simpleService.Accept(temperatureConverterVisitor)
	customService.Accept(temperatureConverterVisitor)

	simpleService.Accept(detailVisitor)
	customService.Accept(detailVisitor)

}
