package main

import (
	"fmt"
	"sort"
)

// Somme des longueurs des chemins
// Utilisé pour déterminer si une combinaison est meilleure qu'une autre avec un même nombre de chemin
func SumOfPathLens(gp GoodPath) int {
	x := 0

	for _, y := range gp.distinctPaths {
		x += len(y)
	}

	return x
}

// true si x contient i
func Contains(x []int, i int) bool {
	for _, y := range x {
		if i == y {
			return true
		}
	}
	return false
}

// true si a et b ont un élément en commun
func HasConflict(a []int, b []int) bool {
	for i := 1; i < len(a)-1; i++ {
		if Contains(b, a[i]) {
			return true
		}
	}
	return false
}

// minimum de a et b
func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// range x en fonction de la taille de ses éléments
func SortByLen(x [][]int) {
	sort.Slice(x, func(i, j int) bool {
		return len(x[i]) < len(x[j])
	})
}

// imprime les informations de la salle
func PrintRoom(x Room) {
	fmt.Printf("Room name : %s \n", x.name)
	fmt.Print("Linked to : ")
	for _, y := range x.tunnels {
		fmt.Printf(" %s ", antHill.allRooms[y].name)
	}
	fmt.Println()
}

// imprime les informations de la fourmillière
func PrintAntHill() {
	fmt.Printf("Start is : %s\nEnd is : %s\n", antHill.allRooms[antHill.start].name, antHill.allRooms[antHill.end].name)
	fmt.Println()

	for _, x := range antHill.allRooms {
		PrintRoom(x)
		fmt.Println()
	}
}

// imprime les noms des salles correspondantes
func PrintPathRoomes(x []int) {
	for _, y := range x {
		fmt.Print(antHill.allRooms[y].name + " ")
	}
	fmt.Print("\n")
}

// imprime les chemins
func PrintGoodPaths(a []GoodPath) {
	for _, x := range a {
		for _, y := range x.distinctPaths {
			PrintPathRoomes(y)
		}
		fmt.Print("\n")
	}
}
