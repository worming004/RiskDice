package logic

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func max(values ...int) int {
	result := MinInt
	for _, value := range values {
		if value > result {
			result = value
		}
	}
	return result
}

func min(values ...int) int {
	result := MaxInt
	for _, value := range values {
		if value < result {
			result = value
		}
	}
	return result
}
