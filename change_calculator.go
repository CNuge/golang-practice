package main

import(
	"fmt"
	"math"
	"sort"
	"strings"
	)

type Transaction struct{
	cost float64
	paid float64
	change_amount float64
}

func (tr Transaction) String() string {
	change_amount :=  tr.paid - tr.cost 
	if change_amount < -0.02{
		return fmt.Sprintf("They have underpaid you by: $%.2f", tr.cost - tr.paid)
	} else if change_amount <= 0.02{
		return "Change due: $0.00"
	} else {
		return fmt.Sprintf("Change due: $%.2f", change_amount)
	}
}


func key_sort(in_map map[float64]int) []float64 {
	var keys []float64
    for k := range in_map {
        keys = append(keys, k)
    }
    sort.Float64s(keys)
    return keys
}


func bill_repr(bill_map map[float64]int) string {
	keys := key_sort(bill_map)
	outdat := []string{"Hand the customer:\n"}
	for i := len(keys)-1; i >= 0; i-- {
		denom := keys[i]
		if bill_map[denom] != 0 {
			if denom >= 5.0{
				addition := fmt.Sprintf("%v $%v bill", bill_map[denom], denom )
				outdat = append(outdat, addition)
			} else {
				addition := fmt.Sprintf("%v $%v coins", bill_map[denom], denom )
				outdat = append(outdat, addition)
			}
			
			if bill_map[denom] > 1 {
				outdat = append(outdat, "s")
			} 
			outdat = append(outdat,"\n")
		}
	}
	return strings.Join(outdat, "")
}


func (tr *Transaction) ChangeBreakdown() string {
	tr.change_amount =  tr.paid - tr.cost 
	var denominations = map[float64]int {100.0 : 0, 50.0 : 0, 20.0 : 0, 10.0 : 0, 5.0 : 0,
											2.0 : 0, 1.0 : 0, 0.25 : 0, 0.1 : 0, 0.05 : 0}
	if tr.change_amount < -0.02 {
		return fmt.Sprintf("They have underpaid you by: $%.2f", (tr.cost - tr.paid))
	}
	// sort the keys
	keys := key_sort(denominations)
	// pull the amount to break down
    total_breakdown := tr.change_amount
	for i := len(keys)-1; i >= 0; i-- {
		coin := float64(keys[i])
		num_denom := math.Floor((total_breakdown / coin) )
		if num_denom != 0 {
			denominations[coin] = int(num_denom)
			total_breakdown = total_breakdown - (num_denom * coin)
		}
	}
	return fmt.Sprintf("%v", bill_repr(denominations)) //, denominations)
}


func main(){
	test1 := Transaction{ cost: 1.75, paid: 5.00}
	fmt.Println(test1)
	fmt.Println(test1.ChangeBreakdown())
	fmt.Println(math.Floor(3.25 / 2))
}