/*

take an m x n matrix and if there is a zero in the matrix then change the entire 
row and column it is in to 0s

*/
package main

import(
	"fmt"
	"strings"
)

type Matrix [][]int


// represents the 2d matrix with each row on a separate line
func (m Matrix) String() string {
	outstring := ""
	for _, i := range m {
		line_string := fmt.Sprintf("%v\n", i)
		s := []string{outstring, line_string}
		outstring = strings.Join(s, "")
	}
	return outstring
}

func ZeroMatrix(mat Matrix) Matrix{
	
	z_i := []int{}
	z_j := []int{}

	for i, row := range mat{
		for j, val := range row{
			if val == 0{
				z_i = append(z_i, i)
				z_j = append(z_j, j)
			}
		}
	}

	for _, i := range z_i{
		for j:=0; j < len(mat[0]); j++{
			mat[i][j] = 0
		}
	}

	for _, j := range z_j{
		for i :=0 ; i < len(mat); i ++{
			mat[i][j] = 0
		}
	}

	return mat

}


func main(){
		t1 := Matrix{{0,1,1,},
					{1,1,1,},
					{1,0,1,},
					{1,1,1,},}


		o1 := Matrix{{0,0,0,},
					{0,0,1,},
					{0,0,0,},
					{0,0,1,},}

		fmt.Println("start 1:\n")
		fmt.Println(t1)
		fmt.Println("post 1:\n")
		fmt.Println(ZeroMatrix(t1))
		fmt.Println("compare 1:\n")
		fmt.Println(o1)


		t2 := Matrix{{1,1,1,1,},
					{1,1,1,1,},
					{1,1,1,0,},
					{1,1,1,1,},}

		o2 := Matrix{{1,1,1,0,},
					{1,1,1,0,},
					{0,0,0,0,},
					{1,1,1,0,},}


		fmt.Println("start 2:\n")
		fmt.Println(t2)
		fmt.Println("post 2:\n")
		fmt.Println(ZeroMatrix(t2))
		fmt.Println("compare 2:\n")
		fmt.Println(o2)
}