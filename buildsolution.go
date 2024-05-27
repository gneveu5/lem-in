package main

import "fmt"

// calcule le nombre de tours nécessaires pour faire passer n fourmis dans gp
func CalculateCost(gp GoodPath, n int) int {
	max := len(gp.distinctPaths[len(gp.distinctPaths)-1]) - 1
	p := len(gp.distinctPaths)

	res := 0

	for _, x := range gp.distinctPaths {
		res += max - (len(x) - 1)
	}

	if res >= n {
		return n + max - 1
	}

	if (n-res)%p == 0 {
		return (n-res)/p + max - 1
	} else {
		return (n-res)/p + max
	}
}

// renvoie l'indice de la solution à utiliser dans solutions
func PrintSolutionOnBestPath(solutions []GoodPath, n int) {
	bestSolCount := CalculateCost(solutions[0], n)
	bestSolIndice := 0
	// fmt.Printf("solution with 1 path costs %d\n", bestSolCount)

	for i := 1; i < len(solutions); i++ {
		solCount := CalculateCost(solutions[i], n)
		// fmt.Printf("solution with %d path costs %d \n", len(solutions[i].distinctPaths), solCount)
		if solCount < bestSolCount {
			bestSolCount = solCount
			bestSolIndice = i
		}
	}

	// fmt.Printf("Picked solution with %d paths \n", len(solutions[bestSolIndice].distinctPaths))

	PrintSolution(solutions[bestSolIndice], n)
}

func PrintSolution(gp GoodPath, n int) {
	waitingAnts := make([][]Ant, len(gp.distinctPaths))

	var runningAnts []Ant

	for i := 0; i < n; i++ {
		min := len(waitingAnts[0]) + len(gp.distinctPaths[0])
		minIndice := 0
		for j := 1; j < len(waitingAnts); j++ {
			p := len(waitingAnts[j]) + len(gp.distinctPaths[j])
			if p < min {
				min = p
				minIndice = j
			}
		}
		waitingAnts[minIndice] = append(waitingAnts[minIndice], Ant{id: i + 1, currentRoom: 0, pathIndice: minIndice})
	}

	flag := true

	for flag {
		flag = false
		for i := 0; i < len(waitingAnts); i++ {
			if len(waitingAnts[i]) > 0 {
				flag = true
				runningAnts = append(runningAnts, waitingAnts[i][0])
				waitingAnts[i] = waitingAnts[i][1:]
			}
		}
		var temp []Ant
		for _, x := range runningAnts {
			x.currentRoom++
			roomIndice := gp.distinctPaths[x.pathIndice][x.currentRoom]
			fmt.Printf("L%d-%s ", x.id, antHill.allRooms[roomIndice].name)
			if roomIndice != antHill.end {
				flag = true
				temp = append(temp, x)
			}
		}
		if flag {
			fmt.Print("\n")
		}
		runningAnts = temp
	}
	fmt.Print("\n")
}
