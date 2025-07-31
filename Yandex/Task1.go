package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	var str string
	var str1 string
	var str2 string
	var str3 string
	var floa1 float64
	var floa2 float64
	var floa3 float64

	fmt.Fscan(os.Stdin, &str)

	const layout = "02.01.2006"
	t, err := time.Parse(layout, str)
	if err != nil {
		log.Fatal(err)
	}

	t = t.AddDate(0, 0, 15)

	s := t.Format("02.01.2006")

	fmt.Fscan(os.Stdin, &str1)
	fmt.Fscan(os.Stdin, &str2)
	fmt.Fscan(os.Stdin, &str3)
	fmt.Fscan(os.Stdin, &floa1)
	fmt.Fscan(os.Stdin, &floa2)
	fmt.Fscan(os.Stdin, &floa3)
	total := floa1 + floa2 + floa3
	ruble := int(total)
	copeika := int(total - float64(ruble))

	result := fmt.Sprintf(`Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.
Дата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день. 
Общая сумма выплат составит %d руб. %d коп. 
С уважением, 
Гл. бух. Иванов А.Е.`, str2, str1, str3, s, ruble, copeika)
	fmt.Println(result)

}

// time.Now()
/*
Первая часть (2024-04-11 17:39:17.756388243) представляет собой дату и время в формате год-месяц-день час:минута:секунда.миллисекунда
Вторая часть (+0300) указывает смещение временной зоны относительно UTC (Всемирного координированного времени) – это основной стандарт времени, используемый в авиации, картах, планах полетов, прогнозах погоды и других областях.
Третья часть (MSK) указывает на часовой пояс. В примере это московское время.
Четвертая часть (m=+0.000053694) содержит вспомогательную информацию, предоставляемую самим языком. В данном случае m обозначает момент времени, а +0.000053694 представляет его как количество секунд с начала выполнения программы или другого опорного момента
*/

/*
Для Printf / Scanf/ Sprintf и все что оканчивается на f

%s — строка;
%d — целое число;
%f — число с плавающей точкой;
%t — булево значение;
%v — значение в стандартном формате (универсальный спецификатор);
%+v — значение с подробным выводом (например, с полями структуры).
*/

/*
Работа с датами

Пакет time
Создание даты
now := time.Now()

// Конкретная дата
specificDate := time.Date(2023, time.November, 15, 23, 0, 0, 0, time.UTC)

Форматирование даты (с использоваием эталонной даты)
Эталонная дата - Mon Jan 2 15:04:05 MST 2006

t := time.Now()
fmt.Println(t.Format("2006-01-02"))               // 2023-11-15
fmt.Println(t.Format("02.01.2006"))               // 15.11.2023
fmt.Println(t.Format("15:04:05"))                 // 23:00:00
fmt.Println(t.Format("Jan _2 2006"))              // Nov 15 2023
fmt.Println(t.Format("2006-01-02 15:04:05 MST"))  // 2023-11-15 23:00:00 UTC

Извлечение компонентов дат

year := now.Year()
month := now.Month()   // Возвращает time.Month (тип-перечисление)
day := now.Day()
hour := now.Hour()
minute := now.Minute()
second := now.Second()
weekday := now.Weekday() // Возвращает time.Weekday

Извлечение отдельных частей
// Только дата (без времени)
dateOnly := now.Format("2006-01-02")
// Только время (без даты)
timeOnly := now.Format("15:04:05")
// Только год
yearOnly := now.Format("2006")
// Только месяц (название)
monthName := now.Format("January")
// Только месяц (число)
monthNumber := now.Format("01")
// Только день
dayOnly := now.Format("02")
// Только день недели
weekdayOnly := now.Format("Monday")

Разница между датами

start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)

diff := end.Sub(start)
fmt.Println("Всего:", diff)
fmt.Println("Часы:", diff.Hours())
fmt.Println("Минуты:", diff.Minutes())
fmt.Println("Секунды:", diff.Seconds())

Парсинг дат из строк

const layout = "2006-01-02"
t, err := time.Parse(layout, "2023-11-15")
if err != nil {
    log.Fatal(err)
}


Операции с датами

// Добавление времени
future := now.Add(24 * time.Hour) // Через 24 часа
past := now.Add(-1 * time.Hour)   // Час назад

// Разница между датами
diff := future.Sub(past)


*/

/*
Работа со строками

// Конкатенация
s1 := "Hello"
s2 := "World"
result := s1 + " " + s2
// Длина строки
length := len(result)
// Доступ к символам (получаем байты!)
firstChar := result[0]
// Подстроки
sub := result[0:5] // "Hello"

Библиотека string

import "strings"
// Проверка наличия подстроки
contains := strings.Contains("hello world", "world")
// Разделение строки
parts := strings.Split("a,b,c", ",") // ["a", "b", "c"]
// Объединение
joined := strings.Join([]string{"a", "b", "c"}, "-") // "a-b-c"
// Замена
replaced := strings.Replace("foo bar baz", "bar", "qux", 1)
// Верхний/нижний регистр
upper := strings.ToUpper("test")
lower := strings.ToLower("TEST")
// Обрезка пробелов
trimmed := strings.TrimSpace("  hello  ")

Регулярные выражения

import "regexp"

matched, err := regexp.MatchString(`^[a-z]+\[[0-9]+\]$`, "adam[23]")
if err != nil {
    log.Fatal(err)
}

re := regexp.MustCompile(`(?P<first>\w+) (?P<last>\w+)`)
result := re.ReplaceAllString("John Smith", "${last}, ${first}")
// "Smith, John"



Перевод from int to string

import "strconv"

// Строка -> число
i, err := strconv.Atoi("123")
f, err := strconv.ParseFloat("3.14", 64)

// Число -> строка
s := strconv.Itoa(123)
s := strconv.FormatFloat(3.14, 'f', 2, 64)

*/

/*
Парсинг с помощью регулярных выражений

logRegex := regexp.MustCompile(`\[(?P<time>\d{2}:\d{2}:\d{2})\] (?P<level>\w+): (?P<msg>.+)`)
logLine := "[14:23:45] ERROR: File not found"

matches := logRegex.FindStringSubmatch(logLine)
if len(matches) > 0 {
    time := matches[logRegex.SubexpIndex("time")]
    level := matches[logRegex.SubexpIndex("level")]
    message := matches[logRegex.SubexpIndex("msg")]
    fmt.Printf("Time: %s, Level: %s, Message: %s\n", time, level, message)
}

*/

/*

Ввод и вывод

Считывние данных

С пробелами как разделители
var name string
var age int
fmt.Fscan(os.Stdin, &name, &age) // Читает до пробела/переноса строки

До переноса строки
fmt.Fscanln(os.Stdin, &name, &age) // Читает до переноса строки

Чтение с форматированием

fmt.Fscanf(os.Stdin, "Name: %s Age: %d", &name, &age)


Чтение строки целиком

reader := bufio.NewReader(os.Stdin)
input, _ := reader.ReadString('\n') // Читает до переноса строки
input = strings.TrimSpace(input)    // Удаляет \n в конце

Вывод

fmt.Print("Hello, ", name, "!") // "Hello, John!"

fmt.Println("Hello,", name, "!") // "Hello, John !\n"

fmt.Printf("Hello, %s! You are %d years old.\n", name, age)


Форматирование строк без вывода

message := fmt.Sprintf("%s is %d years old", name, age)

result := fmt.Sprint("The values are ", 10, " and ", 20.5)

*/
