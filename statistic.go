package piscine

import (
	"bufio"
	"fmt"
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

Average: 35
Median: 4
Variance: 5
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

// ImportNums import the numbers from the file and convert it to a slice of strings
func ImportNums(numberFile string) []string {
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
	for scanned.Scan() {
		source = append(source, scanned.Text())
	}
	return source
}

func main() {

}
