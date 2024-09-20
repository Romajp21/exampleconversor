package main

import "fmt"

func main() {
	fmt.Println("Bienvenido/a al conversor")
	fmt.Println("Selecciona una de las opciones:")
	fmt.Println("1. De Celsius a Fahrenheit")
	fmt.Println("2. De Fahrenheit a Celsius")
	fmt.Println("3. De kilómetros a millas")
	fmt.Println("4. De millas a kilómetros")
	fmt.Println("5. De pulgadas a centímetros")
	fmt.Println("6. De centímetros a pulgadas")
	fmt.Println("7. De pies a metros")
	fmt.Println("8. De metros a pies")

	var selection int
	fmt.Scanln(&selection)

	switch selection {
	case 1:
		fmt.Println("¿Cuántos grados Celsius?")
		var celsius float32
		fmt.Scanln(&celsius)
		fahrenheit := (celsius * 9 / 5) + 32
		fmt.Printf("%.2f grados Celsius son %.2f grados Fahrenheit\n", celsius, fahrenheit)
	case 2:
		fmt.Println("¿Cuántos grados Fahrenheit?")
		var fahrenheit float32
		fmt.Scanln(&fahrenheit)
		celsius := (fahrenheit - 32) * 5 / 9
		fmt.Printf("%.2f grados Fahrenheit son %.2f grados Celsius\n", fahrenheit, celsius)
	case 3:
		fmt.Println("¿Cuántos kilómetros?")
		var km float32
		fmt.Scanln(&km)
		miles := km / 1.60934
		fmt.Printf("%.2f kilómetros son %.2f millas\n", km, miles)
	case 4:
		fmt.Println("¿Cuántas millas?")
		var miles float32
		fmt.Scanln(&miles)
		km := miles * 1.60934
		fmt.Printf("%.2f millas son %.2f kilómetros\n", miles, km)
	case 5:
		fmt.Println("¿Cuántas pulgadas?")
		var inches float32
		fmt.Scanln(&inches)
		centimeters := inches * 2.54
		fmt.Printf("%.2f pulgadas son %.2f centímetros\n", inches, centimeters)
	case 6:
		fmt.Println("¿Cuántos centímetros?")
		var centimeters float32
		fmt.Scanln(&centimeters)
		inches := centimeters / 2.54
		fmt.Printf("%.2f centímetros son %.2f pulgadas\n", centimeters, inches)
	case 7:
		fmt.Println("¿Cuántos pies?")
		var feet float32
		fmt.Scanln(&feet)
		meters := feet * 0.3048
		fmt.Printf("%.2f pies son %.2f metros\n", feet, meters)
	case 8:
		fmt.Println("¿Cuántos metros?")
		var meters float32
		fmt.Scanln(&meters)
		feet := meters / 0.3048
		fmt.Printf("%.2f metros son %.2f pies\n", meters, feet)
	default:
		fmt.Println("Selecciona una opción correcta")
	}
}
