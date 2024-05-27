package main

var antHill AntHill

type AntHill struct {
	startAnts int    // nombre de fourmis au départ
	start     int    // indice de la salle de départ
	end       int    // indice de la salle d'arrivée
	allRooms  []Room // tableaux contenant toutes les salles
}

type Room struct {
	id      int
	name    string
	tunnels []int // indice des salles liées par un tunnel
}

type GoodPath struct {
	distinctPaths [][]int // combinaison de chemins
}

type Ant struct {
	id          int
	pathIndice  int
	currentRoom int
}
