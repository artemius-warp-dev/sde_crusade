package main

type BridgeDateTemperatureService struct {
	implementor DateTemperatureImplementor
}

func NewBridgeDateTemperatureService(implementor DateTemperatureImplementor) *BridgeDateTemperatureService {

	return &BridgeDateTemperatureService{implementor: implementor}

}

func (b *BridgeDateTemperatureService) GetDate() string {

	return b.implementor.FetchDate()

}

func (b *BridgeDateTemperatureService) GetTemperature() float64 {
	return b.implementor.FetchTemperature()
}
