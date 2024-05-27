package main

import (
	"fmt"
	"os"
)

// Construit tous les chemins allant de la première à la dernière salle
func BuildAllPaths() [][]int {
	var res [][]int // contient les chemins valides

	var current [][]int // contient les chemins en cours d'exploration

	var start []int
	start = append(start, antHill.start)
	current = append(current, start)

	for len(current) > 0 {
		// on prend le dernier chemin en cours d'exploration et on complète avec tous les tunnels disponibles
		x := current[len(current)-1]
		current = current[:len(current)-1]
		for _, i := range antHill.allRooms[x[len(x)-1]].tunnels {
			if i == antHill.end {
				// le chemin a atteint la fin
				x = append(x, i)
				res = append(res, x)
				break
			} else if i != antHill.end && !Contains(x, i) {
				// on duplique pour chaque tunnel
				y := make([]int, len(x))
				copy(y, x)
				y = append(y, i)
				current = append(current, y)
			}
		}
	}

	return res
}

// Renvoie les combinaisons de chemins utilisables
// Le résultat contient un chemin opti
func BuildSolutionPath() []GoodPath {
	var res []GoodPath

	allPaths := BuildAllPaths()
	if len(allPaths) == 0 {
		fmt.Println("Error : no path between start and end")
		os.Exit(0)
	}

	// On range les chemins par taille croissante
	SortByLen(allPaths)

	var onePath [][]int

	// la "combinaison" à 1 chemin ets juste le chemin le plus court
	onePath = append(onePath, allPaths[0])
	res = append(res, GoodPath{onePath})

	// nombre maximal de chemins concurents pouvant exister
	n := min(len(antHill.allRooms[antHill.start].tunnels), len(antHill.allRooms[antHill.end].tunnels))

	i := 2
	flag := true

	for flag && i <= n {
		// on cherche à construire une combinaison de i chemins
		var bestResult GoodPath
		sum := 10000000000
		for j, firstPath := range allPaths {
			var gp GoodPath
			gp.distinctPaths = append(gp.distinctPaths, firstPath)
			lastUsed := j
			for n := 0; n < i-1; n++ {
				lastUsed = TryAddPath(&gp, allPaths, lastUsed)
				if lastUsed == -1 {
					// aucun chemin n'a pu être ajouté.
					// on arrête la boucle
					n = i
				}
			}
			if len(gp.distinctPaths) == i {
				// la solution trouvée est valide
				n := SumOfPathLens(gp)
				if n < sum {
					bestResult = gp
					sum = n
				}
			}
		}
		res = append(res, bestResult)
		i++
	}

	return res
}

// Tente d'ajouter un chemin à la combinaison de chemins
// Les chemins d'indices inférieurs à lastUsed ont déjà été essayés
func TryAddPath(gp *GoodPath, allPath [][]int, lastUsed int) int {
	for k := lastUsed + 1; k < len(allPath); k++ {
		app := allPath[k]
		noConflict := true
		for _, gpp := range gp.distinctPaths {
			noConflict = noConflict && !HasConflict(app, gpp)

		}
		if noConflict {
			gp.distinctPaths = append(gp.distinctPaths, app)
			return k
		}
	}

	return -1
}
