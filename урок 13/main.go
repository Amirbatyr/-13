
package main

import "fmt"

func main() {
	// Задание 1
	runningExercises := []string{"Спринт", "Бег на длинную дистанцию"}
	walkingExercises := []string{"Обычная ходьба", "Быстрая ходьба"}

	fmt.Println(len(runningExercises))
	fmt.Println(len(walkingExercises))

	// Задание 2
	subjectsList := []string{"Физика", "Химия", "География"}

	fmt.Println(subjectsList[0])
	fmt.Println(subjectsList[len(subjectsList)-1])

	subjectsList[1] = "Биология"
	fmt.Println(subjectsList)

	// Задание 3
	myList := []string{"Tom", "35", "New York"}

	name := myList[0]
	age := myList[1]
	home := myList[2]

	fmt.Println(name, age, home)

	// Задание 4
	numbersList := []int{1, 2, 3, 4, 5}
	found := false

	for _, number := range numbersList {
		if number == 3 {
			found = true
		}
	}

	if found {
		fmt.Println("Число 3 есть в массиве")
	} else {
		fmt.Println("Число 3 отсутствует в массиве")
	}

	// Задание 5
	friendsList := []string{"Tom", "Bekbolat", "Alex"}
	found = false

	for _, friend := range friendsList {
		if friend == "Bekbolat" {
			found = true
		}
	}

	if found {
		fmt.Println("Мне очень повезло")
	} else {
		fmt.Println("Мне не повезло")
	}

	// Задание 6
	firstList := []int{1, 2, 3}
	secondList := []int{1, 2, 4}
	equal := len(firstList) == len(secondList)

	if equal {
		for i := 0; i < len(firstList); i++ {
			if firstList[i] != secondList[i] {
				equal = false
				break
			}
		}
	}

	if equal {
		fmt.Println("Массивы равны")
	} else {
		fmt.Println("Массивы не равны")
	}

	// Задание 7
	myWishList := []string{"Компьютер", "Наушники"}
	friendsWishList := []string{"Книга", "Мяч"}
	resultList := []string{}

	for _, gift := range myWishList {
		resultList = append(resultList, gift)
	}

	for _, gift := range friendsWishList {
		resultList = append(resultList, gift)
	}

	fmt.Println(resultList)

	// Задание 8
	toyList := []string{"Car", "Doll", "Ball"}
	testToyList := make([]string, len(toyList))
	copy(testToyList, toyList)

	testToyList[1] = "Boat"

	fmt.Println(toyList)
	fmt.Println(testToyList)
}