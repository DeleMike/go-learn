package main

import (
	"example.com/gadget"
	"fmt"
)

func main() {
	//player := gadget.TapePlayer{}
	//mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	//playList(player, mixtape)

	//var value gadget.MyInterface
	//value = gadget.MyType(5)
	//value.MethodWithParameters(127.3)
	//fmt.Println(value.MethodWithReturnValues())

	//var toy NoiseMaker
	//toy = Whistle("Toyco Canary")
	//toy.MakeSound()
	//toy = Horn("Trumpet")
	//toy.MakeSound()

	//play(Whistle("Toyco Canary"))
	//play(Horn("Trumpet"))

	//var err error = checkTemperature(121.379, 100.0)
	//if err != nil {
	//	log.Fatal(err)
	//}

	AcceptAnything(3.4)
	AcceptAnything("yiu")
	AcceptAnything(true)
	AcceptAnything(Horn("Trumpet"))

}

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}

	device.Stop()

}

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
}

type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("OverheatError: %0.2f", o)
}

func checkTemperature(actual, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

type Anything interface {
}

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
}
