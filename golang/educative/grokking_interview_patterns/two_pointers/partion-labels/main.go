package main

import (
	"fmt"
	"strings"
)

func partitionLabels(s string) []int {
	lastOccurence := make([]int, 26)
	for i := 0; i < len(s); i++ {
		lastOccurence[s[i]-'a'] = i
	}

	partitionEnd := 0
	partitionStart := 0
	partitionSizes := []int{}

	for i := 0; i < len(s); i++ {
		if lastOccurence[s[i]-'a'] > partitionEnd {
			partitionEnd = lastOccurence[s[i]-'a']
		}

		if i == partitionEnd {
			partitionSizes = append(partitionSizes, i-partitionStart+1)
			partitionStart = i + 1
		}
	}

	return partitionSizes

}

// Driver code
func main() {
	stringsList := []string{
		"ababcbacadefegdehijhklij",
		"eccbbbbdec",
		"caedbdedda",
		"abcdef",
		"bcbcdd",
	}

	for i, s := range stringsList {
		fmt.Printf("%d.\ts: %s\n", i+1, s)
		result := partitionLabels(s)
		fmt.Printf("\n\tPartition sizes: %v\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
