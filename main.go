package main

import (
	"fmt"
)

func main() {
	fmt.Println("Проверим ваши умствен...навыки сложения.")

	fmt.Println("Введите первое число:")
	var firstNumber int
	fmt.Scan(&firstNumber)

	fmt.Println("Введите второе число:")
	var secondNumber int
	fmt.Scan(&secondNumber)

	correct := firstNumber + secondNumber

	fmt.Println("А теперь сложите эти числа в уме и напишите ответ:")
	var answer int
	fmt.Scan(&answer)

	if correct == answer {
		fmt.Println("А вы молодец, все верно!")
	} else if correct != answer {
		fmt.Println("Неверно! Правильный ответ:", correct)
	}
}
