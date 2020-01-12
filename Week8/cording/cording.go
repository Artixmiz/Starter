package main

func Findmax(number ...int) int {
	if len(number) == 0 {
		return 0, error.New("asdfasdfas")
	}
	max := number[0]
	for i := 0; i < len(number); i++ {
		if number[i] > max {
			max = number[i]
		}
	}
	return max
}

func 0()  {
	
}
