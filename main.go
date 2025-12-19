//A simple Golang calculator that performs calclations

package main

// Declares package name, tells Go this is an executable program

import "fmt"

// Imports Go's built in formatting package, fmt is used for printing output to terminal
//and reading input frm user.

func main() { //func declares function
	var num1, num2 float64 //var declares variable to store user input
	var operator string    //string and float are data types

	fmt.Println("===Simple Calculator ===") //Displays title, Println moves to a new line after printing

	fmt.Print("Enter first number: ") //first number input
	fmt.Scan(&num1)                   //fmt.Scan read user input

	fmt.Print("Enter operator (+, -, *, /): ") //operator
	fmt.Scan(&operator)                        //fmt.Print() shows prompt withpout moving to next line.

	fmt.Print("Enter second number: ") //second number input
	fmt.Scan(&num2)

	fmt.Println("\n--- Result ---") // \n inserts new line(adds space for readability)

	switch operator { //evaluates value of operator. cleaner than multiple if statements
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2) //addition
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2) //subtraction
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2) //Multiplication
	case "/":
		if num2 != 0 {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2) //Division
		} else {
			fmt.Println("Error: Cannot divide by zero!") //prints out the statement when a number is divided by 0
		}
	default:
		fmt.Println("Error: Invalid operator!") // prints out the statement when an
	} //operator other than those stated are used
}
