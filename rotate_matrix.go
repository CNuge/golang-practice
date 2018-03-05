/*
given an nxn matrix, write a method to rotate it 90 degrees.

can you do it in place?
*/

package main

import(
	"fmt"
	"strings"
	)


// make a structure that is a slice of slices

type Matrix struct{
	data [][]int
}

func (m Matrix) String() string {
	outstring := ""
	for _, i := range m.data {
		line_string := fmt.Sprintf("%v\n",  i)
		s := []string{outstring, line_string}
		outstring = strings.Join(s , "")
	}
	return outstring
}


func (m *Matrix) Fill(row []int) [][]int {
	m.data = append(m.data, row)
	return m.data
}


//rotate clockwise
func (m *Matrix) Rotate() [][]int {
	// size has the rows and columns flipped from the original
	new_mat := [len(m.data[0])][len(m.data)]int{}

// iterate through the reversed i, forward j positions
// to the new column eqaul to the old jth row, append the original i,j position

// essentially starting the columns sideways and then appeanding the data
// moving from bottom right to top left

// take the i, j coordinates of the matrix


}




func main(){
	//create a slice of slices
	test_arr := [][]int{}
	// append the rows into the slice
	row1 := []int{1,2,3,}
	row2 := []int{4,5,6,}
	row3 := []int{7,8,9,}

	test_arr = append(test_arr,row1)
	test_arr = append(test_arr,row2)
	test_arr = append(test_arr,row3)

	fmt.Println(test_arr)

	test_arr1 := Matrix{}
	test_arr1.Fill(row1)
	test_arr1.Fill(row2)
	test_arr1.Fill(row3)

	fmt.Println(test_arr1)
	fmt.Println(test_arr1.data[0])
}


