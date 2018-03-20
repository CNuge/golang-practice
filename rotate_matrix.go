/*
given an nxn matrix, write a method to rotate it 90 degrees to the right.
*/

package main

import(
	"fmt"
	"strings"
	)


// a structure that is a slice of slices
// it stores the two dimensional matrix
type Matrix struct {
	data [][]int
}

// represents the 2d matrix with each row on a separate line
func (m Matrix) String() string {
	outstring := ""
	for _, i := range m.data {
		line_string := fmt.Sprintf("%v\n",  i)
		s := []string{outstring, line_string}
		outstring = strings.Join(s , "")
	}
	return outstring
}

// add a new row to the matrix
func (m *Matrix) Fill(row []int) [][]int {
	m.data = append(m.data, row)
	return m.data
}


//rotate the matrix clockwise
func (m *Matrix) Rotate() [][]int {
	// initiate a new matrix
	new_mat := [][]int{}
	// initiate the necessary # of rows == num cols from original
	for i := 0; i < len(m.data[0]); i++ {
		new_mat = append(new_mat, []int{})
	}
	// iterate through the reversed i, forward j positions
	// essentially starting the columns sideways on the left and appending to the rows
	for i := len(m.data) -1 ; i >=0 ; i-- {
		for j, dat := range m.data[i] {
			new_mat[j] =  append(new_mat[j], dat)
		}
	}
	m.data = new_mat
	return m.data
}




func main(){
	// append the rows into the slice
	row1 := []int{1,2,3,}
	row2 := []int{4,5,6,}
	row3 := []int{7,8,9,}
	row4 := []int{10,11,12,}

	test_arr1 := Matrix{}
	test_arr1.Fill(row1)
	test_arr1.Fill(row2)
	test_arr1.Fill(row3)
	test_arr1.Fill(row4)

	fmt.Println(test_arr1)

	test_arr1.Rotate()
	fmt.Println(test_arr1)

	test_arr2 := Matrix{ 
	[][]int{{1,2,3,}, {4,5,6,}, {7,8,9,}, {10,11,12,}} }

	fmt.Println(test_arr2)

	test_arr2.Rotate()
	fmt.Println(test_arr2)

}
