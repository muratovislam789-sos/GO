package main

import "fmt"

func main() {
	var age int
	var name string

	fmt.Print("введите ваше имя:")
	fmt.Scan(&name)

	fmt.Print("введите ваш возвраст:")
	fmt.Scan(&age)

	fmt.Print(("вас зовут "), name, (" ваш возвраст "), age)
}

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	sum := 0

// 	for i := 0; i < n; i++ {
// 		var x int
// 		fmt.Scan(&x)
// 		sum += x
// 	}

// 	a := float64(sum) / float64(n)

// 	fmt.Println(sum)
// 	fmt.Printf("%.2f\n", a)
// }

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	var x int
// 	fmt.Scan(&x)

// 	max := x
// 	min := x

// 	for i := 1; i < n; i++ {
// 		fmt.Scan(&x)
// 		if x > max {
// 			max = x
// 		}
// 		if x < min {
// 			min = x
// 		}
// 	}

// 	fmt.Println("max:", max)
// 	fmt.Println("min:", min)
// }

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	a := 0
// 	b := 0

// 	for i := 0; i < n; i++ {
// 		var x int
// 		fmt.Scan(&x)

// 		if x%2 == 0 {
// 			a++
// 		} else {
// 			b++
// 		}
// 	}

// 	fmt.Println("четный:", a)
// 	fmt.Println("нечетный:", b)
// }

// ПЕРВОЕ ЗАДАНИЕ

// func add(a, b int) int {
// 	return a + b
// }

// func sub(a, b int) int {
// 	return a - b
// }

// func mul(a, b int) int {
// 	return a * b
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// ВТОРОЕ ЗАДАНИЕ

// func isEven(n int) bool {
// 	return n%2 == 0
// }

// func isPositive(n int) bool {
// 	return n > 0
// }

// func abs(n int) int {
// 	if n < 0 {
// 		return -n
// 	}
// 	return n
// }

// func sumToN(n int) int {
// 	sum := 0
// 	for i := 1; i <= n; i++ {
// 		sum += i
// 	}
// 	return sum
// }

// func factorial(n int) int {
// 	if n == 0 || n == 1 {
// 		return 1
// 	}
// 	result := 1
// 	for i := 2; i <= n; i++ {
// 		result *= i
// 	}
// 	return result
// }

// func sumDigits(n int) int {
// 	if n < 0 {
// 		n = -n
// 	}

// 	sum := 0
// 	for n > 0 {
// 		sum += n % 10
// 		n /= 10
// 	}
// 	return sum
// }

// func reverseNumber(n int) int {
// 	sign := 1
// 	if n < 0 {
// 		sign = -1
// 		n = -n
// 	}

// 	reversed := 0
// 	for n > 0 {
// 		reversed = reversed*10 + n%10
// 		n /= 10
// 	}
// 	return reversed * sign
// }

// func countDigits(n int) int {
// 	if n == 0 {
// 		return 1
// 	}

// 	if n < 0 {
// 		n = -n
// 	}

// 	count := 0
// 	for n > 0 {
// 		count++
// 		n /= 10
// 	}
// 	return count
// }

// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	if n == 2 {
// 		return true
// 	}
// 	if n%2 == 0 {
// 		return false
// 	}

// 	for i := 3; i*i <= n; i += 2 {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func divide(a, b int) (int, bool) {
// 	if b == 0 {
// 		return 0, false
// 	}
// 	return a / b, true
// }
