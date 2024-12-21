package main

import "fmt"

type Set struct {
	hashMap map[interface{}]bool
}

// NewSet will initialize and return a new object of Set.
func NewSet() *Set {
	s := new(Set)
	s.hashMap = make(map[interface{}]bool)
	return s
}

// Add will add the value in the Set.
func (s *Set) Add(value interface{}) {
	s.hashMap[value] = true
}

// Delete will delete the value from the set.
func (s *Set) Delete(value interface{}) {
	delete(s.hashMap, value)
}

// Exists will check if the value exists in the set or not.
func (s *Set) Exists(value interface{}) bool {
	_, ok := s.hashMap[value]
	return ok
}

func substring(runes []rune, i, k int) string {
	fmt.Println(i, k, len(runes))
	str := ""
	if k < len(runes) {
		str = string(runes[i:k])
	} else {
		str = string(runes[i:len(runes)])
	}
	return str
}

func findRepeatedSequences(dna string, k int) Set {
	runes := []rune(dna)
	output := *NewSet()
	temp := *NewSet()
	for i, _ := range runes {
		substr := substring(runes, i, k+i)
		if temp.Exists(substr) {
			output.Add(substr)
		} else {
			temp.Add(substr)
		}
	}
	return output
}
