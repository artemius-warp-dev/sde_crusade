package main

import "fmt"

func main() {

	collection := NewTemperatureRecordCollection()
	collection.Add(&TemperatureRecord{Date: "2024-06-09", Temperature: 25.5})
	collection.Add(&TemperatureRecord{Date: "2024-06-10", Temperature: 26.0})
	collection.Add(&TemperatureRecord{Date: "2024-06-11", Temperature: 24.8})

	iterator := collection.CreateIterator()

	for iterator.HasNext() {
		record := iterator.Next()
		fmt.Printf("Date: %s, Temperature: %.1f\n", record.Date, record.Temperature)
	}

}
