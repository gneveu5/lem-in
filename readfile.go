package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Lit le fichier d'entrée
func ReadFile() {
	if len(os.Args) != 2 {
		fmt.Println("Error : wrong number of arguments")
		os.Exit(0)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	n, err := strconv.Atoi(fileScanner.Text())
	if err != nil {
		fmt.Println("Error : can't read number of ants")
		os.Exit(0)
	}

	if n == 0 {
		fmt.Println("Error : number of ants is 0")
		os.Exit(0)
	}

	antHill.startAnts = n
	antHill.start = -1
	antHill.end = -1

	readStart := false
	readEnd := false
	readingRooms := true

	for readingRooms {
		fileScanner.Scan()
		switch fileScanner.Text() {
		case "##start":
			readStart = true
		case "##end":
			readEnd = true
		default:
			if fileScanner.Text()[0] == '#' {
				break
			}
			splitted := strings.Split(fileScanner.Text(), " ")
			if len(splitted) == 3 {
				room := Room{name: splitted[0], id: len(antHill.allRooms)}
				antHill.allRooms = append(antHill.allRooms, room)
				if readStart {
					if antHill.start != -1 {
						fmt.Println("Error : 2 starts")
						os.Exit(0)
					}
					antHill.start = room.id
					readStart = false
				}
				if readEnd {
					if antHill.end != -1 {
						fmt.Println("Error : 2 ends")
						os.Exit(0)
					}
					antHill.end = room.id
					readEnd = false
				}
			} else {
				BuildTunnel(fileScanner.Text())
				readingRooms = false
			}
		}
	}

	if antHill.start == -1 {
		fmt.Println("No start room")
		os.Exit(0)
	}

	if antHill.end == -1 {
		fmt.Println("No end room")
		os.Exit(0)
	}

	for fileScanner.Scan() {
		BuildTunnel(fileScanner.Text())
	}

}

func BuildTunnel(x string) {
	splitted := strings.Split(x, "-")
	if len(splitted) != 2 {
		fmt.Println("Error : bad format reading rooms")
		os.Exit(0)
	}

	var rooms []int

	for _, room := range antHill.allRooms {
		if room.name == splitted[0] || room.name == splitted[1] {
			rooms = append(rooms, room.id)
		}
	}

	if len(rooms) != 2 {
		fmt.Printf("Error : link %s is invalid\n", x)
		os.Exit(0)
	}

	antHill.allRooms[rooms[0]].tunnels = append(antHill.allRooms[rooms[0]].tunnels, rooms[1])
	antHill.allRooms[rooms[1]].tunnels = append(antHill.allRooms[rooms[1]].tunnels, rooms[0])
}
