package arrays

func Sum(numbers [5]int){
  var summed int
	for i := 0; i < 5; i++ {
		summed = summed +numbers[i]
	}

	return summed
}