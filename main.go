package main

// "learnpackage/simpleinterest"

// var p, r, t = -5000.0, 10.0, 1.0

// func init() {
// 	// fmt.Println("Simple interest package initialized")
// 	println("Main package intitialized")
// 	if p < 0 {
// 		log.Fatal("Principal is less the zero")
// 	}
// 	if r < 0 {
// 		log.Fatal("Rate of interest is less the zero")
// 	}
// 	if t < 0 {
// 		log.Fatal("Duration is less the zero")
// 	}
// }
// func number() int {
// 	num := 15 * 5
// 	return num
// }

// func changeLocal(num [5]int) {
// 	num[0] = 55
// 	fmt.Println("inside funtion", num)
// }

// func printarray(a [3][2]string) {
// 	for _, v1 := range a {
// 		for _, v2 := range v1 {
// 			fmt.Printf("%s ", v2)
// 		}
// 		fmt.Printf("\n")
// 	}
// }

// func subtactOne(numbers []int) {
// 	for i := range numbers {
// 		numbers[i] -= 2
// 	}
// }

// func countries() []string {
// 	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
// 	neededCountries := countries[:len(countries)-2]
// 	countriesCpy := make([]string, len(neededCountries))
// 	copy(countriesCpy, neededCountries)
// 	return countriesCpy
// }

// func hello(b ...int, a int) {

// }

// func find(num int, nums []int) {
// func find(num int, nums ...int) {
// 	fmt.Printf("type of nums is %T\n", nums)
// 	found := false
// 	for i, v := range nums {
// 		if v == num {
// 			fmt.Println(num, "found at index", i, "in", nums)
// 			found = true
// 		}
// 	}

// 	if !found {
// 		fmt.Println(num, "not found in ", nums)
// 	}
// 	fmt.Printf("\n")
// }

// func change(s ...string) {
// 	s[0] = "Go"
// 	s = append(s, "playground")
// 	fmt.Println(s)
// }

// type employee struct {
// 	salary  int
// 	country string
// }

func main() {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	map2 := map1
	if map2 == map1 {
	}

	// emp1 := employee{
	// 	salary:  1200,
	// 	country: "USA",
	// }

	// emp2 := employee{
	// 	salary:  1400,
	// 	country: "Canada",
	// }

	// emp3 := employee{
	// 	salary:  1300,
	// 	country: "India",
	// }

	// employeeInfo := map[string]employee{
	// 	"Steve": emp1,
	// 	"Mike":  emp2,
	// 	"Joe":   emp3,
	// }
	// for name, info := range employeeInfo {
	// 	fmt.Printf("Employee: %s Salary:$%d Country: %s\n", name, info.salary, info.country)
	// }

	// fmt.Println("length is", len(employeeInfo))

	// var employeeSalary map[string]int
	// employeeSalary["Steve"] = 1200
	// fmt.Println(employeeSalary)

	// employeeSalary := map[string]int{
	// 	"akhmad": 5000000,
	// 	"dika":   3500000,
	// 	"maya":   4000,
	// }

	// fmt.Println("Original employee salary", employeeSalary)
	// modified := employeeSalary
	// modified["maya"] := 5000000
	// fmt.Println("Employee salary changed", employeeSalary)

	// fmt.Println("map before deletion", employeeSalary)
	// delete(employeeSalary, "maya")
	// fmt.Println("map after deletion", employeeSalary)

	// fmt.Println("Content of the map")
	// for key, value := range employeeSalary {
	// 	fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	// }

	// employee := "dika"
	// salary := employeeSalary[employee]
	// fmt.Println("Salary of", employee, "is", salary)

	// employeeSalary := map[string]int{
	// 	"Akhmad": 5000000,
	// 	"Dika":   3500000,
	// }
	// fmt.Println("Salary of joe is", employeeSalary["joe"])
	// newEmp := "Akhmad"
	// value, ok := employeeSalary[newEmp]
	// if ok == true {
	// 	fmt.Println("Salary of", newEmp, "is", value)
	// 	return
	// }
	// fmt.Println(newEmp, "not found ")

	// Version one
	// employeeSalary := make(map[string]int)
	// employeeSalary["Akhmad"] = 5000000
	// employeeSalary["Dika"] = 3500000
	// employeeSalary["Arif"] = 4000000
	// fmt.Println("employeeSalary map contents: ", employeeSalary)
	// fmt.Println(employeeSalary)

	// welcome := []string{"hello", "world"}
	// change(welcome...)
	// fmt.Println(welcome)

	// nums := []int{89, 90, 95}
	// find(89, nums...)

	// find(89, []int{89, 89, 90, 95})
	// find(45, []int{45, 56, 67, 45, 90, 109})
	// find(78, []int{78, 38, 56, 98})
	// find(87, []int{})

	// r := hello(8, 1)

	// countriesNeeded := countries()
	// fmt.Println(countriesNeeded)
	// nos := []int{8, 7, 6}
	// fmt.Println("slice before function call", nos)
	// subtactOne(nos)
	// fmt.Println("slice after function call", nos)

	// pls := [][]string{
	// 	{"C#", "c++"},
	// 	{"Javascript"},
	// 	{"Go", "Rust"},
	// }

	// for _, v1 := range pls {
	// 	for _, v2 := range v1 {
	// 		fmt.Printf("%s ", v2)
	// 	}
	// 	fmt.Printf("\n")
	// }

	// veggies := []string{"Potatoes", "Tomatoes", "Brinjai"}
	// fruits := []string{"Orange", "Apples"}
	// food := append(veggies, fruits...)
	// fmt.Println("Food:", food)
	// fmt.Println("Food:", len(food), cap(food))

	// var names []string
	// if names == nil {
	// 	fmt.Println("slice is nil going to append")
	// 	names = append(names, "Akhmad", "Arif", "Anggi")
	// 	fmt.Println("name contents:", names)
	// }

	// cars := []string{"Ferrari", "Honda", "Ford"}
	// fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	// cars = append(cars, "Toyota")
	// cars = append(cars, "BMW")
	// cars = append(cars, "Porsche")
	// fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))
	// i := make([]int, 5, 5)
	// fmt.Println(i)

	// fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	// fruitslice := fruitarray[1:3]
	// fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))
	// fruitslice = fruitslice[:cap(fruitslice)]
	// fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))

	// numa := [3]int{78, 79, 80}
	// nums1 := numa[:]
	// nums2 := numa[:]
	// fmt.Println("array before change 1", numa)
	// nums1[0] = 100
	// fmt.Println("array after modification to slice nums1", numa)
	// nums2[1] = 101
	// fmt.Println("array after modification to slice nums2", numa)

	// darr := [...]int{57, 87, 90, 82, 100, 78, 67, 69, 59}
	// dslice := darr[2:5]
	// fmt.Println("array before", darr)
	// for i := range dslice {
	// 	dslice[i]++
	// }
	// fmt.Println("array after", darr)

	// c := []int{6, 7, 8}
	// fmt.Println(c)

	// a := [5]int{76, 77, 78, 79, 80}
	// var b []int = a[1:4]
	// fmt.Println(b)

	// a := [3][2]string{
	// 	{"lion", "tiger"},
	// 	{"cat", "dog"},
	// 	{"piegon", "peacock"},
	// }
	// printarray(a)
	// var b [3][2]string
	// b[0][0] = "apple"
	// b[0][1] = "samsung"
	// b[1][0] = "microsoft"
	// b[1][1] = "google"
	// b[2][0] = "AT&T"
	// b[2][1] = "T-Mobile"
	// fmt.Printf("\n")
	// printarray(b)

	// fmt.Println("Simple interest calculation")
	// p := 5000.0
	// r := 10.0
	// t := 1.0
	// si := simpleinterest.Calculate(p, r, t)

	// fmt.Println("Simple interest is", si)

	// If statement
	// num := 102
	// if num%2 == 0 {
	// 	fmt.Println("The number", num, "is even")
	// 	return
	// }
	// fmt.Println("The number", num, "is odd")
	// } else {
	// 	fmt.Println("The number", num, "is odd")
	// }
	// if num <= 50 {
	// 	fmt.Println(num, "is less than or equal to 50")
	// } else if num >= 51 && num <= 100 {
	// 	fmt.Println(num, "is between 51 and 100")
	// } else {
	// 	fmt.Println(num, "is greater than 100")
	// }
	// if num := 10; num%2 == 0 {
	// 	fmt.Println(num, "is even")
	// } else {
	// 	fmt.Println(num, "is odd")
	// }
	// num := 10
	// if num%2 == 0 {
	// 	fmt.Println("The number is even")
	// } else {
	// 	fmt.Println("The number is odd")
	// }

	// loops
	// exlude do while and while
	// for i := 1; i <= 10; i++ {
	// if i > 5 {
	// 	break
	// }
	// if i%2 == 0 {
	// 	continue
	// }
	// fmt.Printf(" %d", i)
	// }

	// fmt.Printf("\nline after for loop")
	// n := 5
	// for i := 0; i < n; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
	// outer:
	// 	for i := 0; i < 3; i++ {
	// 		for j := 1; j < 4; j++ {
	// 			fmt.Printf("i = %d, j = %d\n", i, j)
	// 			if i == j {
	// 				// break
	// 				break outer
	// 			}
	// 		}
	// 	}

	// i := 0

	// for i <= 10 {
	// 	fmt.Printf("%d ", i)
	// 	i += 2
	// }

	// for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
	// 	fmt.Printf("%d * %d = %d\n", no, i, no*i)
	// }

	// for {
	// 	fmt.Println("Hello World")
	// }

	// finger := 9
	// fmt.Printf("Finger %d is ", finger)
	// switch finger {
	// case 1:
	// 	fmt.Println("Thumb")
	// case 2:
	// 	fmt.Println("Index")
	// case 3:
	// 	fmt.Println("Middle")
	// case 4:
	// 	fmt.Println("Ring")
	// case 5:
	// 	fmt.Println("Pinky")
	// default:
	// 	fmt.Println("incorrect finger number")
	// }

	// letter := "A"
	// fmt.Printf("Letter %s is a ", letter)
	// switch letter {
	// case "a", "e", "i", "o", "u":
	// 	fmt.Println("vowel")
	// default:
	// 	fmt.Println(" not a vowel")
	// }

	// num := 75
	// switch {
	// case num >= 0 && num <= 50:
	// 	fmt.Printf("%d is greather than 0 and less than 50", num)
	// case num >= 51 && num <= 100:
	// 	fmt.Printf("%d is greather than 51 and less than 100", num)
	// case num >= 101:
	// 	fmt.Printf("%d is greather than 100", num)
	// }

	// switch num := number(); {
	// switch num := -25; {
	// case num < 50:
	// 	if num < 0 {
	// 		break
	// 	}
	// 	fmt.Printf("%d is lesser than 50\n", num)
	// 	fallthrough
	// case num < 100:
	// 	fmt.Printf("%d is lesser than 100\n", num)
	// 	fallthrough
	// case num < 200:
	// 	fmt.Printf("%d is lesser than 200\n", num)
	// }
	// randloop:
	// 	for {
	// 		switch i := rand.Intn(100); {
	// 		case i%2 == 0:
	// 			fmt.Printf("Generated even number %d", i)
	// 			break randloop
	// 		}
	// 	}

	// var a [3]int
	// a[0] = 12
	// a[1] = 78
	// a[2] = 60

	// a := [3]int{12, 78, 50}
	// a := [3]int{12, 78, 50, 30}

	// a := [3]int{12}

	// a := [...]int{12, 78, 50}
	// a := [3]int{12, 78, 50}
	// var b [5]int
	// b = a
	// fmt.Println(a)

	// a := [...]string{"USA", "Chine", "India", "Germany", "France"}
	// b := a
	// b[0] = "Singapore"
	// fmt.Println("a is ", a)
	// fmt.Println("b is ", b)

	// num := [...]int{5, 6, 7, 8, 8}
	// fmt.Println("before passing to function", num)
	// changeLocal(num)
	// fmt.Println("After passing to function", num)

	// a := [...]float64{67.7, 89.8, 21, 78}
	// fmt.Println("length of a is", len(a))
	// for i := 0; i < len(a); i++ {
	// 	fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	// }
	// sum := float64(0)
	// for i, v := range a {
	// for _, v := range a {
	// 	fmt.Printf("%d  the element of  a is %.2f\n", i, v)
	// 	sum += v
	// }
	// fmt.Println("\nsum of all elements of a", sum)
}
