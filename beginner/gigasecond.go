package main

import(
	"fmt"
	"math"
)


//Calculate the moment when someone has lived for 10^9 seconds.
//A gigasecond is 10^9 (1,000,000,000) seconds.


func Round(x float64) int {
	/*
	this function takes a floating point as an input and returns the 
	closes integer, 0.5+ rounds up
	*/
	var round float64 // this is our rounder number
	nearest := int(x)
	diff := x - float64(nearest)
	if diff >= 0.5 {
		round = math.Ceil(x)
	}
	if diff < 0.5 {
		round = math.Floor(x)
	}
	return int(round)
}

func main(){
	gs := 1000000000.0

	// take minutes, leave remainder
	mins := int(gs)/60
	rem_sec :=  (gs/60.0 - float64(mins)) * float64(60)

	// take hours, leave remainder
	hr := int(mins)/60
	rem_min := (float64(mins)/60.0 - float64(hr)) * float64(60)

	// take days, leave remainder
	days := int(hr)/24
	rem_hr := (float64(hr)/24.0 - float64(days)) * float64(24)

	// take years from days, leave remainder
	years := days/365
	rm_days := (float64(days)/365.0 - float64(years)) * 365.0

	// account for leap years... drop 1 from the day count for every 4 years
	// this will fail on the first couple days of a year :/
	// ^if clause below deals with that concern

	leap_days := float64(years)/4
	leap_days = float64(int(leap_days))

	if rm_days < leap_days {
		years = years - 1
		rm_days = rm_days +365
	}

	rm_days = rm_days - leap_days


	// round the remainders, and print the string 
	fmt.Println("A a person has been alive a gigasecond after:")
	fmt.Printf("%v years, %v days, %v hours, %v minutes, %v seconds\n", 
				years, Round(rm_days), Round(rem_hr), Round(rem_min), Round(rem_sec))
	fmt.Printf("Alson know as %v days, %v hours, %v minutes, %v seconds\nSince leaps years make things fickle!\n", 
				days, Round(rem_hr), Round(rem_min), Round(rem_sec))
}
