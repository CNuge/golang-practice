package main

import(
	"fmt"
)


/* run through the use of:
	1. structs x
	2. slices / arrays x
	3. maps x
	4. pointers x
*/





func main(){
	// we can make a struct that holds any number of named data fields
	// these fields can have different types

	type Player struct{
		lname string
		goals int
		assists int
		perc_TOI float64
	}
	// make a struct instance
	mitch := Player{"marner", 7, 14, 22.34}
	// the individual parts can be called
	fmt.Println(mitch.lname)
	fmt.Printf("goals: %d\n", mitch.goals)

	// here we make a pointer to mitch, and we can change both values
	mitchy := &mitch
	// mitchy scored!
	mitchy.goals++
	// mitch now contains 8 goals
	fmt.Printf("goals: %d\n", mitch.goals)
	
	// can only fill some fields, and rest are left as 0 or empyty strings
	josh := Player{lname: "levio", assists : 2}
	fmt.Println(josh.lname)
	fmt.Printf("goals: %d\n", josh.goals)


	// arrays are objects of fixed sizes, 

	var line1 [3]int
	line1[0] = 29
	line1[1] = 34
	line1[2] = 11

	fmt.Println(line1)

	//these can be initiated all at once via an array literal
	// array literals make the whole array at onc3
	// in practice slices and slice literals are more often used

	line2 := [3]int{16, 43, 12}

	fmt.Println(line2)


	// a slice is a reference to an unerlying array
	// they can be called on existing arrays
	mar_kad := line2[0:2]
	fmt.Println(mar_kad)

	// but more often the can be created via slice literals,
	// to create a flexibly sized array
	
	defencemen := []int{51,2,22,44,46,23,8,} // trailing comma allowed btw!
	fmt.Println(defencemen)

	// we can even make an array of structs!
	// this would be a lot prettier if the struct were defined beforehand
	// and then called where it says struct
	// note whitespace in the slice literal is okay!
	forwards := [] struct {
		name string
		number int
	}{
		{"Matthews", 34},
		{"Marleau", 12},
		{"Marner", 16},
	}
	fmt.Println(forwards)

	// the last data structure is a map, which maps one value to another in
	// a hash table setup

	type plr struct {
		fname, lname string
	}


	// start a map like this:
	roster := make(map[int]plr)
	//then add players
	roster[46] = plr{"Roman", "Polak"}
	
	fmt.Println(roster)
	
	// or use a map literal outside of a function body, see top
	// note the inner struct type doesn't need to be repeated in the map liters
	var players = map[int]plr{
		16 : {"Mitch", "Marner"},
		51 : {"Jake", "Gardiner"},
		43 : {"Nazem", "Kadri"},
	} 
	// The map can then be added to
	// note when you add to a map, you need to specify the struct type if needed
	players[12] = plr{"Patrick", "Marleau"}

	fmt.Println(players)

	// those seem to be all the main data types in go... very flexible!
	// using these simple structures we can probably accomplish a lot... so lets try
}

