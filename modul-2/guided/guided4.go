package main

import "fmt"

func main() {
	var celsius float64 

	fmt.Print("Masukkan suhu dalam derajat Celsius: ") 
	fmt.Scan(&celsius) 

	fahrenheit := (celsius * 9 / 5) + 32 
	reamur := celsius * 4 / 5 
	kelvin := celsius + 273.15 

	fmt.Printf("Temperatur Celsius: %.2f\n", celsius) 
	fmt.Printf("Derajat Reamur: %.2f\n", reamur) 
	fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit) 
	fmt.Printf("Derajat Kelvin: %.2f\n", kelvin) 
}