package main

type OldDateTemperatureService struct{}

func (ots OldDateTemperatureService) GetDate() string {

	return "2024-06-09"

}

func (ots OldDateTemperatureService) GetTemperature() float64 {
	return 25.5
}
