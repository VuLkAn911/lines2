package main

import (
	"fmt"
	"strings"
)

func main() {
	// sub 1

	str := "GoLang"
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))

	// sub 2

	str2 := " backend developer "
	fmt.Println(strings.Trim(" ", str2))

	//sub 3

	str3 := "I like java"

	if strings.Contains(str3, "java") {
		fmt.Println("Найдено")
	} else {
		fmt.Println("Не найдено")
	}
	// sub 4

	str4 := "one,two,three"

	new_str4 := strings.Replace(str4, ",", ";", 2)
	fmt.Println(new_str4)

	// sub 5

	fmt.Print("Введите строку:")
	var text string
	fmt.Scan(&text)

	if text == "admin" {
		fmt.Println("Доступ разрешен")
	} else {
		fmt.Println("Доступ запрещен")
	}

	// sub 6

	str5 := "Go is fast and Go is simple"
	fmt.Println(strings.Count(str5, "Go"))

	//sub 7

	fmt.Print("Введите строку: ")
	var text2 string
	fmt.Scan(&text2)

	reversed := ""

	for _, v := range text2 {
		reversed = string(v) + reversed

	}
	if strings.EqualFold(text2, reversed) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Не палиндром")
	}
	fmt.Println(reversed)

	// sub 8

	fmt.Print("Введите строку: ")
	var text3 string
	fmt.Scan(&text3)
	falsed := false

	for i := 0; i < len([]rune(text3))-2; i++ {
		if []rune(text3)[i] == []rune(text3)[i+1] && []rune(text3)[i+1] == []rune(text3)[i+2] {

			falsed = true
			break
		}

	}
	if falsed {
		fmt.Println("Одинаковые символы")
	} else {
		fmt.Println("Не одинаковые символы")
	}

	// sub 9

	fmt.Print("Введите строку: ")
	var text4 string
	fmt.Scan(&text4)

	if len([]rune(text4)) < 8 {
		fmt.Println("Слишком короткий пароль")
	} else if !(strings.ContainsAny(text4, "1234567890")) {
		fmt.Println("Пароль должен содержать цифру")
	} else if strings.ToLower(text4) == text4 {
		fmt.Println("Пароль должен содержать заглавную букву")
	} else {
		fmt.Println("Пароль корректный")
	}

	// sub 10

	fmt.Print("Введите строку: ")
	var text5 string
	fmt.Scan(&text5)

	if !(strings.ContainsAny(text5, "@")) {
		fmt.Println("Email должен содержать @")
	} else if !(strings.ContainsAny(text5, "@")) {
		fmt.Println("Email должен содержать точку после @")
	} else {
		fmt.Println("Email Корректен")
	}

	// sub 11

	fmt.Print("Введите строку: ")
	var text6 string
	fmt.Scan(&text6)

	if strings.ToLower(string([]rune(text6)[0])) == string([]rune(text6)[0]) {
		fmt.Println("Не с заглавной буквы")
	} else if !(strings.HasSuffix(text6, ".")) {
		fmt.Println("Не заканчивается точкой")
	} else {
		fmt.Println("Строка оформлена правильно")
	}

	// sub 12

	fmt.Print("Введите строку: ")
	var text7 string
	fmt.Scan(&text7)
	countA := 0
	countE := 0

	for _, v := range text7 {
		if v == 'e' || v == 'E' {
			countE++
		} else if v == 'a' || v == 'A' {
			countA++
		}

	}
	fmt.Println(countE)
	fmt.Println(countA)

}
