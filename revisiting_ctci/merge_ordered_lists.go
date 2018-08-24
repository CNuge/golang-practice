package main

import (
	"fmt"
)

type Ind struct {
	id    string
	score int
}

// take two sorted slices of structs with id names and scores (high to low)
// merge them into a single ordered list with no duplicates
func MergeLists(l1, l2 []Ind) []Ind {

	outlist := []Ind{}
	obvs_ids := make(map[string]int)

	for len(l1) > 0 && len(l2) > 0 {
		if len(l1) == 0 {
			//loop through 2, add all of it to out and return
			for _, i := range l2 {
				if obvs_ids[i.id] == 0 {
					obvs_ids[i.id] = 1
					outlist = append(outlist, i)
				}
			}
			return outlist
		} else if len(l2) == 0 {
			// loop through 1, add all of it to out and return
			for _, i := range l1 {
				if obvs_ids[i.id] == 0 {
					obvs_ids[i.id] = 1
					outlist = append(outlist, i)
				}
			}
			return outlist
		}

		if l1[0].id == l2[0].id {
			if obvs_ids[l1[0].id] == 0 {
				obvs_ids[l1[0].id] = 1

				if l1[0].score >= l2[0].score {
					outlist = append(outlist, l1[0])
					l1 = l1[1:]
					l2 = l2[1:]
					// else accounts for ties
				} else {
					outlist = append(outlist, l2[0])
					l1 = l1[1:]
					l2 = l2[1:]
				}
			}

		} else if l1[0].score > l2[0].score {
			if obvs_ids[l1[0].id] == 0 {
				obvs_ids[l1[0].id] = 1
				outlist = append(outlist, l1[0])
				l1 = l1[1:]
			}

		} else if l1[0].score < l2[0].score {
			if obvs_ids[l2[0].id] == 0 {
				obvs_ids[l2[0].id] = 1
				outlist = append(outlist, l2[0])
				l2 = l2[1:]
			}
		} else if l1[0].score == l2[0].score {
			if obvs_ids[l1[0].id] == 0 {
				obvs_ids[l1[0].id] = 1
				outlist = append(outlist, l1[0])
				l1 = l1[1:]
			}
			if obvs_ids[l2[0].id] == 0 {
				obvs_ids[l2[0].id] = 1
				outlist = append(outlist, l2[0])
				l2 = l2[1:]
			}
		}

	}
	return outlist
}

func main() {

	list1 := []Ind{
		Ind{"Dave", 99},
		Ind{"Jeff", 97},
		Ind{"Karl", 89},
		Ind{"Pete", 77},
	}
	list2 := []Ind{
		Ind{"Karl", 89},
		Ind{"Kevin", 82},
		Ind{"Doug", 79},
		Ind{"Jeff", 77},
	}

	fmt.Println(MergeLists(list1, list2))

}
