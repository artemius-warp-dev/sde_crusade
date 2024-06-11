package main

type TemperatureRecord struct {
	Date        string
	Temperature float64
}

type Iterator interface {
	HasNext() bool
	Next() *TemperatureRecord
}

type TemperatureRecordCollection struct {
	records []*TemperatureRecord
}

func NewTemperatureRecordCollection() *TemperatureRecordCollection {
	return &TemperatureRecordCollection{
		records: []*TemperatureRecord{},
	}
}

func (c *TemperatureRecordCollection) Add(record *TemperatureRecord) {
	c.records = append(c.records, record)
}

func (c *TemperatureRecordCollection) CreateIterator() Iterator {
	return &TemperatureRecordIterator{
		collection: c,
		index:      0,
	}
}

type TemperatureRecordIterator struct {
	collection *TemperatureRecordCollection
	index      int
}

func (i *TemperatureRecordIterator) HasNext() bool {
	return i.index < len(i.collection.records)
}

func (i *TemperatureRecordIterator) Next() *TemperatureRecord {
	if i.HasNext() {
		record := i.collection.records[i.index]
		i.index++
		return record
	}
	return nil

}
