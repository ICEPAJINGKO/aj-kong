package main

import (
	"fmt"
	"gobasic/calculator"
)

// struct
type Product struct {
	name     string
	price    float64
	category string
	qty      int
}

func main() {
	welcomeMessage("ICEPAJINGKO")

	// manual type declaration
	// var name string = "ICEPAJINGKO"
	// var age int = 25
	// var score float32 = 3.14
	// var isPass bool = true

	// type inference
	// name := "ICEPAJINGKO"
	// age := 25
	// score := 3.14
	// isPass := true

	// constant declaration
	// const name string = "ICEPAJINGKO 2"
	// const name = "ICEPAJINGKO 2"

	// fmt.Println("Hello, my name is", name)
	// fmt.Println("I am", age, "years old")
	// fmt.Println("My score is", score)
	// fmt.Println("Am I pass?", isPass)

	// fmt.Printf("My name is %v\n", name)
	// fmt.Printf("Type of age is %T\n", age)
	// fmt.Printf("Type of score is %T\n", score)
	// fmt.Printf("Type of isPass is %T\n", isPass)

	// ----------------------------

	// var number1, number2 = 5, 2

	// fmt.Println("ผลบวก = ", number1+number2)
	// fmt.Println("ผลลบ = ", number1-number2)
	// fmt.Println("ผลคูณ = ", number1*number2)
	// fmt.Println("ผลหาร = ", number1/number2)
	// fmt.Println("ผลหารเอาเศษ = ", number1%number2)

	// ----------------------------

	// number1, number2 := 5, 2

	// fmt.Println("เท่ากันหรือไม่ = ", number1 == number2)
	// fmt.Println("ไม่เท่ากันหรือไม่ = ", number1 != number2)
	// fmt.Println("มากกว่าหรือไม่ = ", number1 > number2)
	// fmt.Println("น้อยกว่าหรือไม่ = ", number1 < number2)

	// ----------------------------

	// scanf

	// var name string
	// fmt.Print("Enter your name: ")
	// fmt.Scanf("%s", &name)
	// fmt.Println("สวัสดี ", name)

	// var age int
	// fmt.Print("Enter your age: ")
	// fmt.Scanf("%d", &age)
	// fmt.Println("อายุ ", age, " ปี")

	// var score float32
	// fmt.Print("Enter your score: ")
	// fmt.Scanf("%f", &score)
	// fmt.Println("คะแนน ", score)

	// ----------------------------

	// if else

	// var score int
	// fmt.Print("Enter your score: ")
	// fmt.Scanf("%d", &score)

	// if score >= 50 {
	// 	fmt.Println("ผ่าน")
	// } else {
	// 	fmt.Println("ไม่ผ่าน")
	// }

	// เลขตคู่หรือเลขคี่
	// var number int
	// fmt.Print("Enter your number: ")
	// fmt.Scanf("%d", &number)

	// if number%2 == 0 {
	// 	fmt.Println("เลขคู่")
	// } else {
	// 	fmt.Println("เลขคี่")
	// }

	// switch case
	// var number int = 2
	// switch number {
	// case 1:
	// 	fmt.Println("เลข 1")
	// case 2:
	// 	fmt.Println("เลข 2")
	// case 3:
	// 	fmt.Println("เลข 3")
	// default:
	// 	fmt.Println("ไม่รู้จักเลข")
	// }

	// ----------------------------

	// Array
	// var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(numbers[2])
	// fmt.Println(len(numbers))

	// score := [...]int{85, 90, 95}
	// fmt.Println(len(score))

	// fmt.Println("2:4", numbers[2:4])

	// slice
	// names := []string{"John0", "Jane1", "Joe2"}
	// fmt.Println(names)
	// names = append(names, "Ice3")
	// fmt.Println(names)
	// fmt.Println(":", names[:])
	// fmt.Println("1:", names[1:])
	// fmt.Println(":1", names[:1])
	// fmt.Println("2:3", names[2:3])
	// fmt.Printf("Type of names is %T\n", names)

	// Maps
	// map[KeyType]ValueType
	// country := map[string]string{}
	// country["th"] = "Thailand"
	// country["jp"] = "Japan"
	// country["kr"] = "Korea"
	// country["us"] = "USA"

	// value, isNotEmpty := country["jpg"]

	// fmt.Println(isNotEmpty)
	// fmt.Println(value)
	// if isNotEmpty {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Not found")
	// }

	// crypto := map[string]string{"btc": "Bitcoin", "eth": "Ethereum", "doge": "Dogecoin"}

	// fmt.Println(crypto["kr"])

	// ----------------------------

	// Loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// var province = []string{"Bangkok", "Chiangmai", "Phuket"}
	// for i := 0; i < len(province); i++ {
	// 	fmt.Println(province[i])
	// }

	// break and continue
	// for i := 1; i <= 5; i++ {
	// 	if i == 3 {
	// 		continue // skip
	// 	}
	// 	fmt.Println(i)
	// }

	// fmt.Println("=====")

	// for i := 1; i <= 5; i++ {
	// 	if i == 3 {
	// 		break // stop
	// 	}
	// 	fmt.Println(i)
	// }

	// range
	// number := []int{10, 20, 30, 40, 50}
	// for i, v := range number {
	// 	fmt.Println("index = ", i, "value = ", v)
	// }

	// for _, v := range number { // ignore index
	// 	fmt.Println("value = ", v)
	// }

	// languages := map[string]string{"TH": "Thai", "EN": "English", "JP": "Japan"}
	// for k, v := range languages {
	// 	fmt.Println("key = ", k, "value = ", v)
	// }

	// ----------------------------

	// Function
	// fmt.Println(sum(5, 2))
	// fmt.Println(sub(5, 2))
	// result1, result2 := sumamtion(5, 2)
	// fmt.Println(result1, result2)
	// fmt.Println(sumInfinity(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	// ----------------------------

	// Struct
	// lays := Product{name: "Lays", price: 10, category: "Snack", qty: 10}
	// fmt.Println(lays)
	// fmt.Println("ชื่อสินค้า", lays.name)
	// fmt.Println("ราคาสินค้า", lays.price)

	// ----------------------------

	// package
	fmt.Println(calculator.Add(5, 2))
	fmt.Println(calculator.Subtract(5, 2))
}

func welcomeMessage(projectName string) {
	fmt.Println("=====> Hello, World! <=====")
	fmt.Println("=====> Welcome to", projectName, "<=====")
}

func sum(number1 int, number2 int) int {
	return number1 + number2
}

func sub(number1, number2 int) int { // same type
	return number1 - number2
}

func sumamtion(number1, number2 int) (int, string) { // multiple return
	total := number1 + number2
	evenOdd := ""

	if total%2 == 0 {
		evenOdd = "even"
	} else {
		evenOdd = "odd"
	}

	return total, evenOdd
}

func sumInfinity(numbers ...int) int { // variadic function
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}
