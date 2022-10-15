package main

import (
	"fmt"
	"golang-dasar/materi"
)

func main() {
	fmt.Println("Hello World")

	//Mengimport fungsi dari file lain di package lain
	// quiz1 := quiz.Multiply(1, 3)
	// fmt.Println("Quiz 1 : ", quiz1)

	//For loop sentences
	// quiz2 := quiz.PrintForLoop("Golang is the best language")
	// fmt.Print(quiz2)

	// x := []map[string]string{}
	// x = append(x, map[string]string{"halo": "halo"})
	// fmt.Println(x)

	// Average
	// numbers := []int{100, 80, 75, 92, 70, 93, 88, 67}
	// quiz3_1 := quiz.Average(numbers)
	// quiz3_2 := quiz.GetBiggerNumber(numbers)

	// fmt.Println(quiz3_1)
	// fmt.Println(quiz3_2)

	// Sum & Average
	// numbers := []int{100, 80, 75, 92, 70, 93, 88, 67}
	// sum := quiz.Sum(numbers)
	// result, err := quiz.Calculate(1, 1, "==")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Print(result)

	//Pointer
	// gamer := quiz.Gamer{Name: "GTA V"}
	// gamer.AddGame("FIFA 22")
	// gamer.AddGame("Super Mario")

	// for _, game := range gamer.Games {
	// 	fmt.Println(game)
	// }

	//Pointer Materi
	// var number1 int = 5
	// number2 := &number1

	// fmt.Println(number1)
	// fmt.Println(number2)
	// *number2 = 3
	// fmt.Println(*number2)
	// fmt.Println(number1)

	// number1 := 5
	// number1 = number1 + 1
	// fmt.Print(number1)

	//Interface
	hitungLuasPersegi := materi.SeberapaLuas(materi.Persegi{Sisi: 4})
	hitungLuasPersegiPanjang := materi.SeberapaLuas(materi.PersegiPanjang{Panjang: 6, Lebar: 4})

	fmt.Printf("Luas Persegi %v", hitungLuasPersegi)
	fmt.Println()
	fmt.Printf("Luas Persegi Panjang %v", hitungLuasPersegiPanjang)
}
