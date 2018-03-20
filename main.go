package main

import "fmt"

func sumOfSquares(numInt int) int{
	if numInt == 0{
		return 0
	}

	var num int 					//integers that range from -100 to 100 to be squared and summed up
	fmt.Scanf("%d",&num)

	if num < 0 {					//only sums positive numbers
		return sumOfSquares(numInt -1)
	} else {
		return sumOfSquares(numInt -1) + num*num
	}
}

func operation(numCases int, sums map[int]int) map[int]int{
	if numCases == 0{
		return sums
	}

	var numInt int         			//number of integers in the following test case ranging from 0 to 100
	fmt.Scanf("%d",&numInt)
	sums[numCases] = sumOfSquares(numInt)
	return operation(numCases-1, sums)
}

func printSums(numCases int, sums map[int]int) {
	if numCases == 0 {
		return
	}

	temp := sums[numCases]
	fmt.Println(temp)

	printSums(numCases-1, sums)
}

func main() {
	var numCases int 				// number of test cases (1 <= numCases <= 100)
	var sums map[int]int

	sums = make(map[int]int)		//initializing map to store calculated sum of squared values

	fmt.Scanf("%d", &numCases)
	result := operation(numCases, sums)  //storing map of the calculated values to pass onto the printing function
	printSums(numCases, result)
}




