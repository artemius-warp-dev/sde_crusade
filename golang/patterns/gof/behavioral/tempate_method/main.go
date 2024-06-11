package main

func main() {

	// Create a simple date temperature service
	simpleService := &SimpleDateTemperatureService{
		date:        "2024-06-09",
		temperature: 25.5,
	}

	// Create a custom date temperature service

	customService := &CustomDateTemperatureService{
		date:        "2024-06-09",
		temperature: 25.5,
	}

	b := BaseDateTemperatureService{service: customService}

	b.Report()

	b = BaseDateTemperatureService{service: simpleService}

	b.Report()

}
