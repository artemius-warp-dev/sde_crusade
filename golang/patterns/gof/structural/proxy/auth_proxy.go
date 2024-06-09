package main

import "fmt"

// AuthProxy represents a proxy for the DateTemperatureService with authentication.
type AuthProxy struct {
	RealService DateTemperatureService
	Username    string
	Password    string
}

// authenticate checks the provided credentials.
func (ap *AuthProxy) authenticate() bool {
	// Implement authentication logic here. For simplicity, we'll just check if the credentials match some hardcoded values.
	return ap.Username == "admin" && ap.Password == "password"
}

// DateSensor checks authentication and proxies the call to the real service's DateSensor method.
func (ap *AuthProxy) DateSensor() string {
	if !ap.authenticate() {
		fmt.Println("AuthProxy: Authentication failed for DateSensor")
		return "Unauthorized"
	}
	fmt.Println("AuthProxy: Authentication successful for DateSensor")
	return ap.RealService.DateSensor()
}

// TemperatureSensor checks authentication and proxies the call to the real service's TemperatureSensor method.
func (ap *AuthProxy) TemperatureSensor() float64 {
	if !ap.authenticate() {
		fmt.Println("AuthProxy: Authentication failed for TemperatureSensor")
		return 0.0
	}
	fmt.Println("AuthProxy: Authentication successful for TemperatureSensor")
	return ap.RealService.TemperatureSensor()
}
