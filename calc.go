package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman = map[byte]int{
	105: 1,    //I
	118: 5,    //V
	120: 10,   //X
	108: 50,   //L
	99:  100,  //C
	100: 500,  //D
	109: 1000, //M
}

func toInt(x string) int {
	x = strings.ToLower(x)
	ans := 0
	if len(x) == 1 {
		return roman[x[0]]
	}
	for i := 0; i < len(x)-1; i++ {
		if roman[x[i]] < roman[x[i+1]] {
			ans -= roman[x[i]]
		} else {
			ans += roman[x[i]]
		}
	}
	return ans + roman[x[len(x)-1]]
}

func toRoman(x int) string {
	ones := [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hrns := [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	ths := [4]string{"", "M", "MM", "MMM"}
	return ths[x/1000] + hrns[(x%1000)/100] + tens[(x%100)/10] + ones[x%10]
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите выражение:")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		a := strings.Split(text, " ") //Массив вида: [число1, арифм. действие, число2]

		if a[0] == "0" || a[2] == "0" {
			panic("Калькулятор принемает только числа от 1 до 10 включительно")
		}

		num1, _ := strconv.Atoi(a[0]) //Преобразование к int. Если в a[0] и a[2] римские цифры, то в num1 и num2 записываются нули
		num2, _ := strconv.Atoi(a[2])

		if (num1 == 0 && num2 != 0) || (num1 != 0 && num2 == 0) {
			panic("Вы пытаетесь совместить арабскую и римскую систему счисления! Не надо так.")
		}

		var isRoman bool
		if num1 == 0 && num2 == 0 { //Калькулятор не принемает на вход нули. Если одно из чисел 0, значит пользователь вводит римские цифры.
			num1 = toInt(a[0])
			num2 = toInt(a[2])
			isRoman = true
		}

		//Ограничение возможностей калькулятора
		if num1 > 10 || num1 < 0 || num2 > 10 || num2 < 0 {
			panic("Калькулятор принемает только числа от 1 до 10 включительно")
		}
		if isRoman && a[1] == "-" && num1 <= num2 {
			panic("В римской системе счистления нет таких понятий как ноль или отрицательное число")
		}

		var ans int
		switch a[1] {
		case "+":
			ans = num1 + num2
		case "-":
			ans = num1 - num2
		case "*":
			ans = num1 * num2
		case "/":
			ans = num1 / num2
		default:
			panic("Калькулятор пока не умеет выполнять такое арифмитическое действие как " + a[1])
		}

		//Вывод ответа
		if isRoman {
			fmt.Println("Ответ:", toRoman(ans))
		} else {
			fmt.Println("Ответ:", ans)
		}
	}
}
