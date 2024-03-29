package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

/*
Your program must be able to read from a file and print the result of each statistic mentioned above. In other words, your program must be able to read the data present in the path passed as argument. The data in the file will be presented as the following example:

189
113
121
114
145
110
...

This data represents a statistical population: each line contains one value.

To run your program a command similar to this one will be used if your project is made in Go:

$> go run your-program.go data.txt

After reading the file, your program must execute each of the calculations asked above and print the results in the following manner (the following numbers are only examples):

Average: 132
Median: 150
Variance: 784.66667
Standard Deviation: 65

Please note that the values are rounded integers.
*/

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                int
}

// BTreeInsertInt insert the data into TreeNode
func BTreeInsertInt(root *TreeNode, data int) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertInt(root.Left, data)
		if root.Left != nil {
			root.Left.Parent = root
		}
	} else {
		root.Right = BTreeInsertInt(root.Right, data)
		if root.Right != nil {
			root.Right.Parent = root
		}
	}
	return root
}

// Atoi converts a string to a number
func Atoi(s string) int {
	var b int
	isNeg := false
	for i := 0; i <= len(s)-1; i++ {
		// check for double pos or neg
		if len(s) > 1 {
			if s[i] == '+' && (s[i+1] == '+' || s[i+1] == '-') || s[i] == '-' && (s[i+1] == '-' || s[i+1] == '+') {
				b = 0
				break
			}
		}
		// check for space
		if s[i] == ' ' {
			b = 0
			break
		}

		// check for digits
		if (s[i] >= '0' && s[i] <= '9') || s[i] == '+' || s[i] == '-' {
			if s[i] == '+' {
			} else if s[i] == '-' {
				isNeg = true
			} else {
				b += int(s[i] - 48)
				if i <= len(s)-2 {
					b *= 10
				}
			}
		} else {
			b = 0
			break
		}
	}
	if isNeg {
		return -(b)
	} else {
		return b
	}
}

// ImportNums import the numbers from the file and convert it to a slice of strings
func ImportNums(numberFile string) []int {
	file, err := os.Open(numberFile)
	if err != nil {
		fmt.Println("Error opening the file:", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing the file:", err)
		}
	}(file)
	scanned := bufio.NewScanner(file)
	scanned.Split(bufio.ScanLines)
	var source []string
	var digits []int
	for scanned.Scan() {
		source = append(source, scanned.Text())
	}
	for _, number := range source {
		digits = append(digits, Atoi(number))
	}
	return digits
}

func Average(numbers []int) float64 {
	total := 0
	counter := 0
	for _, number := range numbers {
		total += number
		counter++
	}
	return float64(total / counter)
}

func Median(numbers []int) float64 {
	lowest := numbers[0]
	highest := numbers[0]
	for _, number := range numbers {
		if number < lowest {
			lowest = number
		}
		if number > highest {
			highest = number
		}
	}
	return float64(highest - (highest-lowest)/2)
}

//func TwoDigitSquare(digits int) int {
//	a := digits % 10
//	b := digits - a
//	var square int
//	square = (a*a)*100 + (2*a*b)*10 +
//}

func Variance(numbers []int) float64 {
	var variance float64
	var total float64
	var counter float64
	for _, number := range numbers {
		variance = (float64(number) - Average(numbers)) * (float64(number) - Average(numbers))
		total += variance
		counter++
	}
	return total / counter
}

func main() {
	numbers := ImportNums(os.Args[1])
	fmt.Printf("Average: %v\n", math.Round(Average(numbers)))
	fmt.Printf("Median: %v\n", math.Round(Median(numbers)))
	fmt.Printf("Variance: %v\n", math.Round(Variance(numbers)))
	fmt.Printf("Standard Deviation: %v\n", math.Round(math.Sqrt(Variance(numbers))))
}
