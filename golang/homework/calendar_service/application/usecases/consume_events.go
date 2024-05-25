package usecases

func ConsumeEvents() (<-chan []byte, error) {

	return EventQueue.Consume(EventQueue.GetQueue())

}
