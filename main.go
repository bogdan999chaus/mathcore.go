package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"mathcore.go/domain"
)

const totalPoints = 100
const pointsperQuestion = 20

var id uint64 = 1

func menu() {
	fmt.Println("1. Грати")
	fmt.Println("2. Рейтинг")
	fmt.Println("3. Вийти")
}

func main() {
	fmt.Println("Вітаю у грі Богдана Чауса \\m/")

	var users []domain.User

	for {
		menu()

		choice := ""
		fmt.Scan(&choice)

		switch choice {
		case "1":
			u := play()
			users = append(users, u)
		case "2":
			for _, u := range users {
				fmt.Printf("Id: %v Name: %s Time: %v\n",
					u.Id, u.Name, u.Time)

			}
		case "3":
			return
		default:
		}

	}

}

func play() domain.User {
	fmt.Println("Good luck!")
	myPoints := 0
	start := time.Now()

	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)

		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("НЕПРАВИЛЬНО!!!")
		} else {
			if ansInt == x+y {
				myPoints += pointsperQuestion
				fmt.Printf("Правильно! Кількість балів: %v\n", myPoints)
			} else {
				fmt.Println("Спробуй ще!")
			}
		}
	}

	end := time.Now()
	timeSpent := end.Sub(start)
	fmt.Printf("Вітаю! Ти впорався(лась) за %v!", timeSpent)
	fmt.Println("Введіть ім'я: ")
	name := ""

	fmt.Scan(&name)

	user := domain.User{
		Id:   id,
		Name: name,
		Time: timeSpent,
	}
	id++

	return user
}
