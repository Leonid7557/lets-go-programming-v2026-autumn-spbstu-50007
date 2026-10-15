package main

import "fmt"

const (
	minAllowedTemp = 15
	maxAllowedTemp = 30
)

func processRoom() {
	var staffCount int

	if _, err := fmt.Scan(&staffCount); err != nil {
		fmt.Println("Invalid input")
		return
	}

	currentMin := minAllowedTemp
	currentMax := maxAllowedTemp

	for i := 0; i < staffCount; i++ {
		var sign string
		var temp int

		if _, err := fmt.Scan(&sign, &temp); err != nil {
			fmt.Println("Invalid input")
			return
		}

		switch sign {
		case ">=":
			if temp > currentMin {
				currentMin = temp
			}
		case "<=":
			if temp < currentMax {
				currentMax = temp
			}
		default:
			fmt.Println("Invalid operator")
			return
		}

		if currentMin > currentMax {
			fmt.Println(-1)
		} else {
			fmt.Println(currentMin)
		}
	}
}

func main() {
	var roomsCount int

	if _, err := fmt.Scan(&roomsCount); err != nil {
		fmt.Println("Invalid input")
		return
	}

	for i := 0; i < roomsCount; i++ {
		processRoom()
	}
}
