package main

import "fmt"

type Player struct {
	firstName string
	lastName string
	team string
}

func commonPlayers(slice1, slice2 []*Player) (common []string) {
	hash := make(map[string]bool)

	for _, player := range slice1 {
		hash[player.firstName + " " + player.lastName] = true
	}

	for _, player := range slice2 {
		fullName := player.firstName + " " + player.lastName
		if hash[fullName] {
			common = append(common, fullName)
		}
	}

	return
}

func main() {
	basketballPlayers := []*Player{
		{firstName: "Jill", lastName: "Huang", team: "Gators"},
		{firstName: "Janko", lastName: "Barton", team: "Sharks"},
		{firstName: "Wanda", lastName: "Vakulskas", team: "Sharks"},
		{firstName: "Jill", lastName: "Moloney", team: "Gators"},
		{firstName: "Luuk", lastName: "Watkins", team: "Gators"},
	}

	footballPlayers := []*Player{
		{firstName: "Hanzla", lastName: "Radosti", team: "32ers"},
		{firstName: "Tina", lastName: "Watkins", team: "Barleycorns"},
		{firstName: "Alex", lastName: "Patel", team: "32ers"},
		{firstName: "Jill", lastName: "Huang", team: "Barleycorns"},
		{firstName: "Wanda", lastName: "Vakulskas", team: "Barleycorns"},
	}

	fmt.Println(commonPlayers(basketballPlayers, footballPlayers))
}