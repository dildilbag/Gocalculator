package main

import "fmt"

func main() {

	var p, s = 35, 7 //p is int number 1 & s is int number 2

	fmt.Println("x + y = ", p+s)
	fmt.Printf("x - y = %d\n", p-s)
	fmt.Printf("x * y = %d\n", p*s)
	fmt.Printf("x / y = %d\n", p/s)

	///////////Tasks With Scan////////////
	/*
		var a, b int /// a is NUmber 1 & b is NUm 2 /
		fmt.Println("Give number Here for + : ")
		fmt.Scan(&a, &b)
		//fmt.Println("Test :", &a)
		var add = a + b
		fmt.Println("Totat is =", add)

		var d, e int
		fmt.Println("Give number Here for know to your Rest NUmber")
		fmt.Scan(&d, &e)
		var subtraction = d - e
		fmt.Println("Rest is =", subtraction)

		var v, x int
		fmt.Println("Give number Here for * : ")
		fmt.Scan(&v, &x)
		var mul = v * x
		fmt.Println("Your Multiplication   is =", mul)

		var h, j int
		fmt.Println("Give number Here for / : ")
		fmt.Scan(&h, &j)
		var division = h / j
		fmt.Println("Your Division is =", division)
	*/
	var num1 int
	var num2 int

	fmt.Println("Give your number Here :")
	fmt.Scan(&num1, &num2)

	addition := num1 + num2
	sub := num1 - num2
	multi := num1 * num2
	div := num1 / num2

	fmt.Println("Choose your Operator : +, -, *, / ")

	var operator string
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Println("Here is your Ans :", addition)
	case "-":
		fmt.Println("Here is your Ans :", sub)
	case "*":
		fmt.Println("Here is your Ans :", multi)
	case "/":
		fmt.Println("Here is your Ans :", div)

	default:
		fmt.Println("Give Please write number")
	}
	var opt string
	fmt.Println("Do you want Exit or continue")
	fmt.Scan(&opt)

}
