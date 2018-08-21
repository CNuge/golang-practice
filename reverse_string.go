package main


import(
	"fmt"
	)



func Reverse(s string) string{
	outstring := ""
	for i := (len(s) - 1) ; i >= 0 ; i-- {
		outstring = fmt.Sprintf("%v%v", outstring, string(s[i]))
	}
	return outstring
}



func main(){

	input := "forward esrever"


	fmt.Println(Reverse(input))

}