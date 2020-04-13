package main

import( "fmt"
	"sort"
)

// var numbers = []float64{15, 99, 40, 73, 2, 24, 36, 95, 9}
var numbers = []float64{15, 99, 40, 73, 2, 24, 95, 9}

func MyPrinter() {
		fmt.Println("Your numbers:", numbers)	
}


func MySorter() {
		s_numbers := numbers
		sort.Float64s(s_numbers)
		fmt.Println("Your sorted numbers:",s_numbers)			
}


func HowLong() {
		fmt.Println("Len of your numbers:", len(numbers))	
}


func MyMin() {
		s_numbers := numbers
		sort.Float64s(s_numbers)
		fmt.Println("Min:", s_numbers[0])	
}


func MyMax() {
		s_numbers := numbers
		lenghts := len(numbers)
		index := lenghts - 1
	    sort.Float64s(s_numbers)
		fmt.Println("Max:", s_numbers[index] )	
}


func EvenOdd() {
		num := len(numbers)
		if num % 2 == 0 {
		        fmt.Println("Len of your numbers is EVEN.")
		} else {
		        fmt.Println("Len of your numbers is ODD.")}
}

func Median() {
	s_numbers := numbers
	sort.Float64s(s_numbers)
	lenghts := len(s_numbers)
	index1 := lenghts / 2 - 1
	index2 := lenghts / 2 
	if lenghts % 2 == 0 {
		fmt.Println("Result:", (s_numbers[index1] + s_numbers[index2])/2)
		} else {
		fmt.Println("Result:", s_numbers[(lenghts -1) / 2])
		}

	
}

func main() {
		MyPrinter()
		MySorter()
		HowLong()
		MyMin()
		MyMax()
		EvenOdd()
		Median()

		
}