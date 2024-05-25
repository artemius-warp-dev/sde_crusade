//go:generate gojson -input=../entities/json_schemas/event.json -o=event/event.go -name=Event

package value_objects

//go:generate gofmt -w event/event.go
