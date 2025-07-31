//package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Quntity(x int, y int) string {
	return "Осталось свободных мест: " + strconv.Itoa(x) + "\n" + "Всего человек в очереди: " + strconv.Itoa(y)
}

func Sup(final_res string, s_num string, number int, arr [5]int, flag bool, change_str string, count int) (string, string, [5]int, int) {
	if flag {
		final_res = final_res + "\n" + "Запись на место номер " + strconv.Itoa(number) + " невозможна: очередь переполнена"
	} else if arr[number-1] == 1 {
		final_res = final_res + "\n" + "Запись на место номер " + strconv.Itoa(number) + " невозможна: место уже занято"
	} else {
		switch number {
		case 1:
			s_num = "1. " + change_str
			arr[number-1] = 1
		case 2:
			s_num = "2. " + change_str
			arr[number-1] = 1
		case 3:
			s_num = "3. " + change_str
			arr[number-1] = 1
		case 4:
			s_num = "4. " + change_str
			arr[number-1] = 1
		case 5:
			s_num = "5. " + change_str
			arr[number-1] = 1
		}
		count++
	}

	return final_res, s_num, arr, count
}

func main() {

	var final_res string
	re := regexp.MustCompile(`-?\d+\.?\d*`)

	s1 := "1. -"
	s2 := "2. -"
	s3 := "3. -"
	s4 := "4. -"
	s5 := "5. -"
	var arr = [5]int{0, 0, 0, 0, 0}
	flag := false
	count := 0

	var current_mood_of_queue string

	scanner := bufio.NewScanner(os.Stdin)
	var input strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "конец" {
			break
		}
		input.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
		return
	}

	sss := strings.Split(input.String(), "\n")
	sst := sss[:len(sss)-1]
	for _, el := range sst {
		//Считывем текущую строку

		current_str := strings.TrimSpace(el) // Удаляет \n и другие пробельные символы

		//Считаем текущее положение очереди
		current_mood_of_queue = s1 + "\n" + s2 + "\n" + s3 + "\n" + s4 + "\n" + s5

		//Считаем количество занятых мест

		if count == 5 {
			flag = true
		}
		//Начинаем обрабатывать

		if current_str == "конец" {
			break
		} else if current_str == "очередь" {
			final_res = final_res + "\n" + current_mood_of_queue
		} else if current_str == "количество" {
			final_res = final_res + "\n" + Quntity(5-count, count)

		} else if current_str == "" {
			continue
		} else if !re.MatchString(current_str) {
			final_res = final_res + "\n" + "некорректный ввод"
		} else if re.MatchString(current_str) {

			num_in_str := re.FindString(current_str)
			num_in_int, _ := strconv.ParseFloat(num_in_str, 64)

			if num_in_int >= 1 && num_in_int <= 5 && float64(int(num_in_int)) == num_in_int {
				change_str := strings.Trim(current_str, " ")
				change_str = re.ReplaceAllString(change_str, "")

				switch num_in_int {
				case 1:
					final_res, s1, arr, count = Sup(final_res, s1, 1, arr, flag, change_str, count)
				case 2:
					final_res, s2, arr, count = Sup(final_res, s2, 2, arr, flag, change_str, count)
				case 3:
					final_res, s3, arr, count = Sup(final_res, s3, 3, arr, flag, change_str, count)
				case 4:
					final_res, s4, arr, count = Sup(final_res, s4, 4, arr, flag, change_str, count)
				case 5:
					final_res, s5, arr, count = Sup(final_res, s5, 5, arr, flag, change_str, count)
				}
			} else {
				final_res = final_res + "\n" + "Запись на место номер " + strconv.Itoa(int(num_in_int)) + " невозможна: некорректный ввод"
			}

		}

	}
	final_res = final_res[1:] + "\n" + s1 + "\n" + s2 + "\n" + s3 + "\n" + s4 + "\n" + s5
	fmt.Print(final_res)

}

/*

Математические операции

число с плавающей точкой float64(a) / float64(b)

import "math"

// Возведение в степень (float64)
result := math.Pow(2, 3) // 8.0

// Квадратный корень
sqrt := math.Sqrt(16) // 4.0

// Кубический корень
cbrt := math.Cbrt(27) // 3.0

angle := math.Pi / 2 // 90 градусов в радианах

// Основные тригонометрические функции
sin := math.Sin(angle)  // 1.0
cos := math.Cos(angle)  // ~0
tan := math.Tan(angle)  // очень большое число

// Обратные тригонометрические функции
asin := math.Asin(1)   // π/2
acos := math.Acos(0)   // π/2
atan := math.Atan(1)   // π/4

// Натуральный логарифм
ln := math.Log(math.E) // 1.0

// Десятичный логарифм
log10 := math.Log10(100) // 2.0

// Логарифм по основанию 2
log2 := math.Log2(8) // 3.0

x := 3.6

// Округление до ближайшего целого
round := math.Round(x) // 4.0

// Округление в меньшую сторону
floor := math.Floor(x) // 3.0

// Округление в большую сторону
ceil := math.Ceil(x) // 4.0

// Отбрасывание дробной части
trunc := math.Trunc(x) // 3.0

import "complex"

// Создание комплексного числа
z := complex(3, 4) // 3 + 4i

// Получение действительной и мнимой частей
realPart := real(z) // 3.0
imagPart := imag(z) // 4.0

// Модуль комплексного числа
modulus := cmplx.Abs(z) // 5.0

import "math/big"

// Большие целые числа
bigInt := new(big.Int)
bigInt.SetString("12345678901234567890", 10)

// Большие числа с плавающей точкой
bigFloat := new(big.Float)
bigFloat.SetString("123.456e+20")

// Рациональные числа
rat := new(big.Rat)
rat.SetString("3/4")
*/
