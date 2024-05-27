package main

func main() {
	ReadFile()

	solutionsPath := BuildSolutionPath()

	// PrintGoodPaths(solutionsPath)

	PrintSolutionOnBestPath(solutionsPath, antHill.startAnts)
}
