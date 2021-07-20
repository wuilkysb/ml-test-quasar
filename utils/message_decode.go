package utils

import (
	"strings"

	"github.com/thoas/go-funk"
)

func printUniqueValue(arr []string) map[string]int {
	dict:= make(map[string]int)
	for _ , num :=  range arr {
		dict[num] = dict[num]+1
	}
	return dict
}

func MinOf(vars map[int]int) int {
	min := vars[0]
	var index int
	for idx, i := range vars {
		if min > i {
			min = i
			index = idx
		}
	}

	return index
}

func GetMostCompleteMessage(message1, message2, message3 []string) []string {
	count1 := printUniqueValue(message1)
	count2 := printUniqueValue(message2)
	count3 := printUniqueValue(message3)

	counts := map[int]int{
		0: count1[""],
		1: count2[""],
		3: count3[""],
	}
	messages := map[int][]string{
		0: message1,
		1: message2,
		3: message3,
	}

	return messages[MinOf(counts)]
}

func GetCompleteMessage(message1, message2, message3 []string) string {
	messages := map[string][]string{
		strings.Join(message1, ","): message1,
		strings.Join(message2, ","): message2,
		strings.Join(message3, ","): message3,
	}


	mostComplete := GetMostCompleteMessage(message1, message2, message3)
	mostCompleteIdx := strings.Join(mostComplete, ",")
	completeLen := len(mostComplete)
	delete(messages, mostCompleteIdx)

	for _, message := range messages {
		factorIdx := completeLen - len(message)
		_, messageDiff := funk.Difference(mostComplete, message)
		if len(messageDiff.([]string)) > 0 {
			for _, v := range messageDiff.([]string) {
				messageIdx := funk.IndexOf(message, v)
				mostComplete[factorIdx+messageIdx] = v
			}
		}
	}
	return strings.Trim(strings.Join(mostComplete, " "), " ")
}
