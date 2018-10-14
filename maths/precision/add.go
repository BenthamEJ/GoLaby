package precision

func Add(a [12]int64, b [12]int64) []int64 {
	answer := make([]int64, 12)
	carry := make([]int64, 12)
	excess := make([]int64, 12)

	var quintillion int64 = 1000000000000000000

	for i := 11; i > 0; i-- {
		carry[i] = (a[i] + b[i]) / quintillion
		excess[i] = carry[i] * quintillion
		answer[i] = answer[i] + a[i] + b[i] - excess[i]
		answer[i-1] += carry[i]
	}

	if (answer[0] + a[0] + b[0]) < quintillion {
		answer[0] = answer[0] + a[0] + b[0]
	} else {
		// result out of bounds
	}

	return answer
}
