package main

import (
	"errors"
	"fmt"
)

const (
	minAllowedTemp  = 15
	maxAllowedTemp  = 30
	minAllowedCount = 1
	maxAllowedCount = 1000
)

func processRoom() error {
	var staffCount int

	if _, err := fmt.Scan(&staffCount); err != nil {
		return fmt.Errorf("reading staff count: %w", err)
	}

	if staffCount < minAllowedCount || staffCount > maxAllowedCount {
		return errors.New("invalid staff count")
	}

	currentMin := minAllowedTemp
	currentMax := maxAllowedTemp

	for range staffCount {
		var sign string

		var temp int

		if _, err := fmt.Scan(&sign, &temp); err != nil {
			return fmt.Errorf("reading temperature requirement: %w", err)
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
			return errors.New("invalid operator")
		}

		if currentMin > currentMax {
			fmt.Println(-1)
		} else {
			fmt.Println(currentMin)
		}
	}

	return nil
}

func main() {
	var roomsCount int

	if _, err := fmt.Scan(&roomsCount); err != nil {
		fmt.Printf("reading rooms count: %s\n", err.Error())

		return
	}

	if roomsCount < minAllowedCount || roomsCount > maxAllowedCount {
		fmt.Printf("invalid rooms count: %d\n", roomsCount)

		return
	}

	for range roomsCount {
		if err := processRoom(); err != nil {
			fmt.Println(err.Error())

			return
		}
	}
}
