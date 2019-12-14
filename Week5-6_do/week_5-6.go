package week5_6

import (
	"fmt"
	"strconv"
)

//FizzWord รับพารามิเตอร์ 2 ตัว เมื่อนำตัวแรกมาหารด้วยตัวที่สองลงตัวจะคืนค่าเป็น "Fizz" หากไม่ลงตัวจะคืนค่าเป็นพารามิตเอร์ตัวแรกแต่มีไทป์เป็น string
func FizzWord(number int, mod int) string {
	if number%mod == 0 {
		return "Fizz"
	} else {
		return strconv.Itoa(number) //strconv
	}

}

//MultiplicationTable จะรับพารามิเตอร์ 1 ตัวแล้วคืนค่าเป็นจำนวนสูตรคูณตั้งแต่ 1-12 กลับไป เช่น รับ n มาเป็น 2 จะคืนค่า []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24}
func MultiplicationTable(n int) []int {
	sum := []int{}
	for xloob := 1; xloob <= 12; xloob++ {
		input := xloob * n
		sum = append(sum, input)
	}
	return sum
}

//StringMultiplicationTable เหมือนกับ MultiplicationTable ในรูปแบบ string เช่น 2 x 1 = 2 ...
func StringMultiplicationTable(n int) []string {
	b := []string{}
	for index := 1; index <= 12; index++ {
		a := fmt.Sprintf("%d x %d = %d", n, index, n*index)
		b = append(b, a)
	}

	return b
}

//DeleteIntItem รับพารามิเตอร์ 2 ตัว ตัวแรกเป็น array ตัวที่ 2 เป็นตัวเลข หากในอาเรย์มีตัวเลขเหมือนพารามิเตอร์ตัวที่ 2 ให้ลบทิ้งและคืนค่าที่ได้กลับไป
func DeleteIntItem(b []int, n int) []int {
	input := a
	if a != nil {
		for i := 0; i < len(input); i++ {
			if input[i] == y {
				input = append(input[:i], input[i+1]...)
				i--
			}
		}
	}
	return input
}

//Grade รับค่าเป็นคะแนน แล้วคืนค่าเป็นเกรด A,B,C,D,F โดยใช้เกณฑ์มาตฐาน A มากกว่าหรือเท่ากับ 80
func Grade(point float32) string {
	if point >= 80 {
		return "A"
	} 
	if point >= 70 {
		return "B"
	} 
	if point >= 60 {
		return "C"
	}
	if point >= 50 {
		return "D"
	} 
	}return "F"
}
