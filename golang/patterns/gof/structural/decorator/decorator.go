package main

import "fmt"

type BaseDecorator struct {
	WrappedService DateTemperatureService
}

func (bd *BaseDecorator) GetDate() string {

	return bd.WrappedService.GetDate()

}

func (bd *BaseDecorator) GetTemperature() float64 {
	return bd.WrappedService.GetTemperature()
}

type LoggingDecorator struct {
	BaseDecorator
}

func NewLoggingDecorator(service DateTemperatureService) *LoggingDecorator {

	return &LoggingDecorator{
		BaseDecorator: BaseDecorator{WrappedService: service},
	}

}

func (ld *LoggingDecorator) GetDate() string {
	fmt.Println("LoggingDecorator: Calling GetDate")
	result := ld.WrappedService.GetDate()
	fmt.Println("LoggingDecorator: GetDate returned", result)
	return result
}

func (ld *LoggingDecorator) GetTemperature() float64 {
	fmt.Println("LoggingDecorator: Calling GetTemperature")
	result := ld.WrappedService.GetTemperature()
	fmt.Println("LoggingDecorator: GetTemperature returned", result)
	return result
}

type CachingDecorator struct {
	BaseDecorator
	cachedDate            string
	cachedTemperature     float64
	dateCacheValid        bool
	temperatureCacheValid bool
}

func NewCachingDecorator(service DateTemperatureService) *CachingDecorator {

	return &CachingDecorator{

		BaseDecorator: BaseDecorator{WrappedService: service},
	}

}

func (cd *CachingDecorator) GetDate() string {
	if !cd.dateCacheValid {
		cd.cachedDate = cd.WrappedService.GetDate()
		cd.dateCacheValid = true
	}

	return cd.cachedDate
}

func (cd *CachingDecorator) GetTemperature() float64 {

	if !cd.temperatureCacheValid {
		cd.cachedTemperature = cd.WrappedService.GetTemperature()
		cd.temperatureCacheValid = true
	}

	return cd.cachedTemperature

}
