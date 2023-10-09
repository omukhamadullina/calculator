package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	a      int
	err    error
	b      int
	SysSch int
	rezalt int
)

func main() {

	for {
		printer()
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		//ждем ввода примера в формате строки

		text = strings.TrimSpace(text)
		//очистим все пробелы, табуляции - заменим все пробелы пустой строкой

		var example = strings.Fields(text)
		//разбиваем строку текст на подстроки

		if len(example) != 3 {
			fmt.Println(errors.New("error: incorrect data entry"))
			return
		}

		var roman = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
		//срез с римскими основными буквами

		SysSch := chekSysSch(example, roman)
		//проверка на римские числа, введенные большими латинскими буквами
		//если метод равен 0, то оба числа арабские

		if len(example) == 3 && SysSch == 0 {
			a, err = strconv.Atoi(example[0])
			if err != nil {
				fmt.Println(errors.New("error: the first number was entered incorrectly"))
				return
			}
			b, err = strconv.Atoi(example[2])
			if err != nil {
				fmt.Println(errors.New("error: the second number was entered incorrectly"))
				return
			}
			// преобразуем первую и третью подстроки в числа а и б

			if (a > 10) || (b > 10) {
				fmt.Println(errors.New("error: calculator processes numbers from 0 to 10"))
				return
			}
			//проверка вводимых чисел

			rezalt := calc(&a, &b, example)
			fmt.Println(text, " = ", rezalt)
		} // вычисления по арабским числам

		// ниже вычисления по римским числам
		if len(example) == 3 && SysSch != 0 {
			a = chekA(roman, example)
			b = chekB(roman, example)
			//перевели оба римских числа в арабсике

			if (a == 0) || (b == 0) {
				fmt.Println(errors.New("error: you cannot use different number systems at the same time"))
				return
			}
			//если только одно число римское

			if (a < b) && (a == b) {
				fmt.Println(errors.New("error: there are no zero and negative numbers in the Roman system"))
				return
			}

			rezalt := calc(&a, &b, example)
			//вычисляем пример, написанный римскими цифрами, в арабских числах

			rezaltRoman(rezalt, text, roman)
			//результат = арабское число, записываем римкими цифрами
		}
	}
}

func printer() {
	fmt.Println("= = = Это простой калькулятор = = =")
	fmt.Println("Обрабатывает ДВА числа, каждое от 1 до 10 включительно")
	fmt.Println("Введите исходные данные в строку, через ПРОБЕЛ")
	fmt.Println("= = = Можно использовать римские цифры = = =")
	//подсказка пользователю
}

func calc(a *int, b *int, example []string) int {

	switch example[1] {
	case "+":
		rezalt = *a + *b
	case "-":
		rezalt = *a - *b
	case "*":
		rezalt = *a * *b
	case "/":
		if *b != 0 {
			rezalt = *a / *b
		} else {
			panic(errors.New("error: dividing by zero is not correct"))
		}
	default:
		panic(errors.New("error: invalid input format"))
	}
	return rezalt
	//выбор выполнения арифметической операции по второй подстроке среза
}

func chekSysSch(example []string, roman []string) int {
	for i := 0; i < len(example); i++ {
		for j := 0; j < len(roman); j++ {
			if example[i] == roman[j] {
				SysSch = SysSch + 1
			}
		}
	}
	return SysSch
	// выбор системы счисления, если = 0, то арабская
}

func chekA(roman []string, example []string) int {
	for i := 0; i < len(roman); i++ {
		if roman[i] == example[0] {
			a = i + 1
		}
	}
	return a
}

func chekB(roman []string, example []string) int {
	for i := 0; i < len(roman); i++ {
		if roman[i] == example[2] {
			b = i + 1
		}
	}
	return b
}

func rezaltRoman(rezalt int, text string, roman []string) {
	fmt.Print(text, " = ")
	switch {
	case rezalt > 0 && rezalt < 11:
		for i := 0; i < len(roman); i++ {
			if rezalt-1 == i {
				fmt.Println(roman[i])
			}
		}
	case rezalt > 10 && rezalt < 21:
		fmt.Print("X")
		for i := 0; i < len(roman); i++ {
			if rezalt-11 == i {
				fmt.Println(roman[i])
			}
		}
	case rezalt > 20 && rezalt < 31:
		fmt.Print("XX")
		for i := 0; i < len(roman); i++ {
			if rezalt-21 == i {
				fmt.Println(roman[i])
			}
		}
	case rezalt > 30 && rezalt < 40:
		fmt.Print("XXX")
		for i := 0; i < len(roman); i++ {
			if rezalt-31 == i {
				fmt.Println(roman[i])
			}
		}
	case rezalt == 40:
		fmt.Println("XL")

	case rezalt > 40 && rezalt < 50:
		fmt.Print("XL")
		for i := 0; i < len(roman); i++ {
			if rezalt-41 == i {
				fmt.Println(roman[i])
			}
		}
	case rezalt == 50:
		fmt.Println("L")

	case rezalt > 50 && rezalt < 61:
		fmt.Print("L")
		for i := 0; i < len(roman); i++ {
			if rezalt-51 == i {
				fmt.Println(roman[i])
			}
		}

	case rezalt > 60 && rezalt < 71:
		fmt.Print("LX")
		for i := 0; i < len(roman); i++ {
			if rezalt-61 == i {
				fmt.Println(roman[i])
			}
		}
	case rezalt > 70 && rezalt < 81:
		fmt.Print("LXX")
		for i := 0; i < len(roman); i++ {
			if rezalt-71 == i {
				fmt.Println(roman[i])
			}
		}
	case rezalt > 79 && rezalt < 90:
		fmt.Print("LXXX")
		for i := 0; i < len(roman); i++ {
			if rezalt-81 == i {
				fmt.Println(roman[i])
			}
		}

	case rezalt == 90:
		fmt.Println("XC")

	case rezalt > 90 && rezalt < 100:
		fmt.Print("XC")
		for i := 0; i < len(roman); i++ {
			if rezalt-91 == i {
				fmt.Println(roman[i])
			}
		}
	case rezalt == 100:
		fmt.Println("C")
	default:
		fmt.Println(errors.New("error"))
		return
	}
	//вернули результату римские цифры
}
