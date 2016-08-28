package main

import "fmt"

const (
	null = iota
	addition = iota
	subtraction = iota
	multiplication = iota
	division = iota
)

func main() {
	fmt.Println("Welcome to Sudhir Calculator")
	var quit int
	for {
		var datavalue int = printdata()
		switch datavalue{
		case null:
			quit = 1
		case addition:
			calculator(datavalue)
		case subtraction:
			calculator(datavalue)
		case multiplication:
			calculator(datavalue)
		case division:
			calculator(datavalue)
		default:
			fmt.Println("Unexpected input")
		}
		if quit == 1{
			break;
		}
	}
}

func calculator(data int){
	var value int
	var value2 int
	var result int
	fmt.Println("Please enter first value")
	fmt.Scanf("%var", &value)
	fmt.Println("Please enter secound value")
	fmt.Scanf("%var", &value2)

	switch data{
	case 1:
		result = value + value2
		fmt.Printf("Result of Addition %v\n", result)
	case 2:
		result = value - value2
		fmt.Printf("Result of Subtraction %v\n", result)
	case 3:
		result = value * value2
		fmt.Printf("Result of Multiplication %v\n", result)
	case 4:
		result = value/value2
		fmt.Printf("Result of Division %v\n", result)
	}


}

func printdata() int{
	var value int
	fmt.Println("Press 0 to EXIT: ")
	fmt.Println("Press 1 for Addition: ")
	fmt.Println("Press 2 for Subtraction: ")
	fmt.Println("Press 3 for Multiplication: ")
	fmt.Println("Press 4 for Division: ")
	fmt.Scanf("%v", &value)
	return value
}

