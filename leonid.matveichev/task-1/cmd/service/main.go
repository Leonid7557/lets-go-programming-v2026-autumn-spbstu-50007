package main

import "fmt"

func main() {
	var first_number, second_number int
	var mathematical_operation string

	if _, err := fmt.Scan(&first_number); err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if _, err := fmt.Scan(&second_number); err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if _, err := fmt.Scan(&mathematical_operation); err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch mathematical_operation {
	case "+":
		fmt.Println(first_number + second_number)
	case "-":
		fmt.Println(first_number - second_number)
	case "*":
		fmt.Println(first_number * second_number)
	case "/":
		if second_number == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(first_number / second_number)
	default:
		fmt.Println("Invalid operation")
	}
}
