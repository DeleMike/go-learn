package main

import (
	"fmt"
)

func main() {
	var price int = 100
	fmt.Println("Price is", price, "dollars.")

	taxRate := 0.08
	tax := taxRate * float64(price)
	fmt.Println("Tax is", tax, "dollars.")

	total := float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")

	availableFunds := float64(120)
	fmt.Println(availableFunds, "available.")
	fmt.Println("Within budget:", (total <= availableFunds))

	// accountAge := 2.6
	// fmt.Println("Account Age is", reflect.TypeOf(accountAge))
	// accountAgeInt := int(accountAge)

	// fmt.Println("Your account has existed for", accountAgeInt, "years")
	// fmt.Printf("I am %v years old", 10)

	// playWithFormatting()
	// ifConditionals()
	// fmt.Println(getCoo())
	// testStruct()

	// rectangle := Rect{width: 4, height: 5}
	// area := fmt.Sprintf("The area of the rectangele is %.1f", (rectangle.area()))
	// perimeter := fmt.Sprintf("The perimeter of the rectangele is %.1f", (rectangle.perimeter()))

	// fmt.Println(area)
	// fmt.Println(perimeter)

	// empType1 := fulltime{name: "Akindele@Google", salary: 100000}
	// empType2 := contractor{name: "Akindele@AIResearchUni", hourlyPay: 500, hoursPerYear: 1000}
	// empType3 := contractor{name: "Akindele@GHCommunity", hourlyPay: 300, hoursPerYear: 500}

	// fmt.Printf("%s is earning %.2f\n", (empType1.name), float64(empType1.getSalary()))
	// fmt.Printf("%s is earning %.2f\n", (empType2.name), float64(empType2.getSalary()))
	// fmt.Printf("%s is earning %.2f\n", (empType3.name), float64(empType3.getSalary()))

	// var i interface{} = "Hello, there!"
	// emp, ok := i.(string)

	// if ok {
	// 	fmt.Println("I am an employee and my name is: ", emp)
	// } else {
	// 	fmt.Println("This is a lie.")
	// }

	// printNumericValue("empType1")

	// i, err := strconv.Atoi("44n")
	// if err != nil {
	// 	error1 := errors.New("You cannot do this nau!")
	// 	fmt.Println(error1)
	// }else {
	// 	fmt.Println(i)
	// }

	// count:= 0
	// for count <= 5 {
	// 	fmt.Println(count)
	// 	count++
	// }

	// fizzBuzz(15)
	// mapPlay()
}

type Key struct {
	providerName string
	providerCode int
}

func mapPlay() {
	areaCodes := make(map[string]string)
	areaCodes["nigeria"] = "+234"
	fmt.Println("Area code is, ", areaCodes)
	nijaKey, ok := areaCodes["nigerhia"]
	fmt.Println("okay status = ", ok)
	fmt.Println("9ja key is = ", nijaKey)

	key := Key{providerName: "Google", providerCode: 2}
	fmt.Println("Struct is: ", key)
}

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

type userError struct {
	name string
	age  int
}

func (e userError) Error() string {
	return fmt.Sprintf("%v has a problem with their account", e.name)
}

func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case float64:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func playWithFormatting() {

	const name = "Andrew ng"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)

	fmt.Printf(msg)
}

func ifConditionals() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

	if length := -1; length < 1 {
		fmt.Println("Message code is invalid")
	}
}

func getCoo() (x, y int) {
	x, y = 2, 3
	return
}

func testStruct() {
	type Car struct {
		Make  string
		Model string
		age   float64
		price float64
		Wheel struct {
			Radius   float64
			Material string
		}
	}

	car := Car{Make: "Mercedes-Benz", Model: "gx-500", age: 0.2, price: 350000}
	car.Wheel.Material = "Black Puma"
	car.Wheel.Radius = 30.5
	fmt.Println("Car is ", car)

	// anonymous structs
	user := struct {
		name string
		age  int
	}{
		name: "Ada Lovelace",
		age:  70,
	}

	fmt.Println("User is ", user)

}

type Rect struct {
	width  float64
	height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * (r.height + r.width)
}

type shape interface {
	area() float64
	perimeter() float64
}

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fulltime struct {
	name   string
	salary int
}

func (ft fulltime) getSalary() int {
	return ft.salary
}

func (ft fulltime) getName() string {
	return ft.name
}

type User struct {
	name string
	age  int
}

func getUser() (User, error) {
	user := User{name: "Akindele", age: 23}
	return user, nil
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	costForCustomer, error1 := sendSMS(msgToCustomer)
	if error1 != nil {
		return 0, error1
	}

	costForCustomerSpouse, error2 := sendSMS(msgToSpouse)
	if error2 != nil {
		return 0, error2
	}

	return costForCustomer + costForCustomerSpouse, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
