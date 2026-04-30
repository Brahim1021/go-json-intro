package main

import (
	"encoding/json"
	"fmt"
	// "log"
	// "os" // Импортируем пакет для работы с ОС
)

// type Human struct {
// 	Name   string
// 	Age    int
// 	Gender string
// }

//type Создай файл coffee.json и накидай в нём пример JSON данных, которые будут соответствовать следующей структуре:

type CoffeeShop struct {
	Name    string   `json:"name"`
	City    string   `json:"city"`
	Items   []string `json:"items"`
	OpenNow bool     `json:"open_now"`
}

type Group struct {
	Title    string `json:"title"`
	Students int    `json:"students"`
	Remote   bool   `json:"remote"`
}

type School struct {
	School   string  `json:"school"`
	Location string  `json:"location"`
	Groups   []Group `json:"groups"`
}

type WorkDay struct {
	Title    string   `json:"title"`
	City     string   `json:"city"`
	Address  string   `json:"address"`
	IsOnline bool     `json:"is_online"`
	Seats    int      `json:"seats"`
	Mentors  []string `json:"mentors"`
}

type Student struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	City      string `json:"city"`
	WantsGo   bool   `json:"wants_go"`
}

type Lesson struct {
	Topic   string `json:"topic"`
	Minutes int    `json:"minutes"`
}

type Meetup struct {
    Title   string `json:"title"`
    Minutes string    `json:"minutes"`
	}

func main() {
	
	workDay := WorkDay{
		Title:    "Первый день в интукод",
		City:     "Грозный",
		Address:  "улица Дади Айбики",
		IsOnline: true,
		Seats:    20,
		Mentors: []string{
			"Ахмад Кудузов",
			"Расул Назиров",
		},
	}

	jsonBytes, err := json.Marshal(workDay)
	if err != nil {
		fmt.Println("Ошибка")
		return
	}

	fmt.Println(string(jsonBytes))

	studentJSON := `{
	"first_name": "Иса",
	"last_name": "Исаев",
	"age": 18,
	"city": "Грозный",
	"wants_go": true
	}`
	fmt.Println(studentJSON)

	student := Student{}
	json.Unmarshal([]byte(studentJSON), &student)
	fmt.Println(student)

	if student.WantsGo == true && student.Age > 16 {
		fmt.Printf("%s %s готов к буткэмпу в Грозном\n", student.FirstName, student.LastName)
	} else {
		fmt.Println("Слишком маленький возраст")
	}
	
	//////////задание 4
	data := []byte(`{"topic":"Введение в JSON","minutes":45}`)
	var l Lesson
	fmt.Println(l)
	err = json.Unmarshal(data, &l)

	fmt.Println("\nОшибка:", err)
	fmt.Printf("Результат: %+v\n", l)

	///////////////задание 5
	newData := []byte(`{ "title": "Разбор задач", "minutes": "60" }`)
	var d Meetup
	errData := json.Unmarshal(newData, &d)

	fmt.Println("Ошибка: ", errData)

	fmt.Println("Результат:  ", d)
	//ошибка cannot unmarshal string into Go struct field Meetup.minutes of type int
	//меняю тип данных для поля Minutes структуры Meetup

}
