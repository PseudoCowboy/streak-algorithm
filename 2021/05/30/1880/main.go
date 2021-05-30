package main

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	firstValue, secondValue, targetValue := 0, 0, 0
	for i := 0; i < len(firstWord); i++ {
		firstValue = firstValue*10 + int(firstWord[i]-'a')
	}
	for i := 0; i < len(secondWord); i++ {
		secondValue = secondValue*10 + int(secondWord[i]-'a')
	}
	for i := 0; i < len(targetWord); i++ {
		targetValue = targetValue*10 + int(targetWord[i]-'a')
	}
	return firstValue+secondValue == targetValue
}
