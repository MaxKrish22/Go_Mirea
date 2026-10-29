package main

import "fmt"

func main() {

	var name string = "Макс"
	var group string = "OOOO-05-26"
	const course string = "Бэкенд-разработка на языке Go"
	var num_works int = 100
	var interes bool = true

	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Printf("Курс: %s\n", course)
	fmt.Printf("Завершено работ: %d\n", num_works)
	fmt.Printf("Интерес к бэкенду: %t\n", interes)
}
