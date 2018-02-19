package main

import(
	"fmt"
)


func main(){
	// make a slice named nhl teams, which corresponds to the array given
	var nhl_teams = []string{"AFM","ANA","ARI","ATL","BOS","BRK","BUF","CAR","CGS","CGY","CHI","CBJ","CLE","CLR","COL","DAL","DFL","DCG","DET","EDM","FLA","HAM","HFD","KCS","LAK","MIN","MMR","MNS","MTL","MWN","NSH","NJD","NYA","NYI","NYR","OAK","OTT","PHI","PHX","PIR","PIT","QUA","QUE","QBD","SEN"}
	fmt.Printf("len=%d, capacity=%d\n",len(nhl_teams), cap(nhl_teams))

	for i, val := range nhl_teams{
		fmt.Printf("team_num=%d, team_name=%v\n", i, val)
	}
}