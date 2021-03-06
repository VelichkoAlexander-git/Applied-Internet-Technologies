Applied-Internet-Technologies(GOlang)
======
### Table of Contents

**[1. Ваша первая программа](#ваша-первая-программа)**<br>
**[2. Типы](#типы)**<br>
**[3. Переменные](#переменные)**<br>
**[4. Управление потоком](#управление-потоком)**<br>
**[5. Массивы, срезы, карты](#массивы-срезы-карты)**<br>
**[6. Функции](#функции)**<br>
**[7. Указатели](#указатели)**<br>
**[8. Структуры и интерфейсы](#структуры-и-интерфейсы)**<br>
**[9. Многопоточность](#многопоточность)**<br>
**[10. Пакеты и повтореное использование кода](#пакеты-и-повторение-использование-кода)**<br>
**[11. Тестирование](#тестирование)**<br>

***

## Ваша первая программа

### 1.  Что такое разделитель?

> Символы новой строки, пробелы и символы табуляции называются разделителями.

### 2.  Что такое комментарий? Назовите два способа записи комментариев.

> Строка, начинающаяся с `//`, является комментарием. Go поддерживает два вида комментариев: `//` превращает в комментарий весь текст до конца строки и `/* */`, где комментарием является всё, что содержится между символами `*` (включая переносы строк).

### 3.  Наша программа начиналась с *package main*. С чего начинаются файлы в пакете *fmt*?

```go
package fmt
```

### 4.  Мы использовали функцию *Println* из пакета *fmt*. Если бы мы хотели использовать функцию *Exit* из пакета *os*, что бы для этого потребовалось сделать?

> Подключить библиотеку командой: `import “os”`

### 5.  Измените написанную программу так, чтобы вместо *Hello World* она выводила *Hello, my name is* вместе с вашем именем.

```Go
package main

import "fmt"

func main() {
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello, my name is" + " " + name)
}
```

## Типы

### 1.  Как хранятся числа в компьютере?

> В отличие от десятичного представления чисел, которое используем мы, компьютеры используют двоичное представление.

### 2.  Мы знаем, что в десятичной системе самое большое число из одной цифры - это 9, а из двух - 99. В бинарной системе самое большое число из двух цифр это 11 (3), самое большое число из трех цифр это 111 (7) и самое большое число из 4 цифр это 1111 (15). Вопрос: каково самое большое число из 8 цифр? (Подсказка: 101-1=9 и 102-1=99)

> 10^1-1=9  
10^2-1=99  
…  
10^8-1=99999999  
В дес. сист.  
0101 1111 0101 1110 0000 1111 1111 (99999999)  
В бинар. сист.  
1111 1111 (255)


### 3.  В зависимости от задачи вы можете использовать Go как калькулятор. Напишите программу, которая вычисляет *32132 × 42452* и печатает это в терминал (используйте оператор * для умножения).

```Go
package main

import "fmt"

func main() {
    fmt.Println("32132 x 42452 =", 32132*42452)
}
```
> 32132 x 42452 = 1364067664

### 4.  Что такое строка? Как найти её длину?

> Строка — это последовательность символов определенной длины, используемая для представления текста.  
Нахождение длины строки `len("Hello World")`

### 5.  Какое значение примет выражение *(true && false) || (false && true) || !(false && false)?*

> Ответ: True

## Переменные

### 1.  Существуют два способа для создания новой переменной. Какие?

> Переменные в Go создаются с помощью ключевого слова `var`, за которым следуют имя переменной (`x`), тип (`string`) и присваиваемое значение (*Hello World*).

```Go
var x string = "Hello World"
```

> Если мы хотим присвоить значение переменной при её создании, то можем использовать сокращенную запись

```Go
x := "Hello World"
```

### 2.  Какое будет значение у *x* после выполнения *x := 5; x += 1*?

> Ответ: 6

### 3.  Что такое область видимости и как определяется область видимости переменной в Go?

> Места, где может использоваться переменная `x`, называются областью видимости переменной.  
Согласно спецификации «в Go область видимости ограничена блоками». В основном это значит, что переменные существуют только внутри текущих фигурных скобок `{ }` (в блоке), включая все вложенные скобки (блоки).

### 4.  В чем отличие *var* от *const*?

> Отличие в том, что переменные образованный от *const* не могут быть изменены после инициализации, а переменные образованные от *var* могут.

### 5.  Используя пример программы выше напишите программу, переводящую температуру из градусов Фаренгейта в градусы Цельсия. (*C = (F - 32) * 5/9*)

```Go
package main

import "fmt"

func main() {
	fmt.Print("Enter a Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32.0) * 5 / 9

	fmt.Println("Celsius: ", output)
}
```

### 6.  Напишите другую программу для перевода футов в метры (1 фут = 0.3048 метр).

```Go
package main

import "fmt"

func main() {
	fmt.Print("Enter a feet: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 0.3048

	fmt.Println("meters: ", output)
}
```

## Управление потоком

### 1.  Что делает следующий код?

```Go
i := 10

if i > 10 {
    fmt.Println("Big")
} else {
    fmt.Println("Small")
}
```

> •	создать переменную i со значением 10;  
•	i больше 10? нет;  
•	вывести “Small”;  
•	выходим.

### 2.  Напишите программу, которая выводит числа от 1 до 100, которые делятся на 3. (3, 6, 9, …).

```Go
package main

import "fmt"

func main() {
    for i := 1; i < 100; i++ {
        if i%3 == 0 {
            fmt.Print(i, " ")
        }
    }
}
```

> Ответ: 3 6 9 12 15 18 21 24 27 30 33 36 39 42 45 48 51 54 57 60 63 66 69 72 75 78 81 84 87 90 93 96 99

### 3.  Напишите программу, которая выводит числа от 1 до 100. Но для кратных трём нужно вывести «Fizz» вместо числа, для кратных пяти - «Buzz», а для кратных как трём, так и пяти — «FizzBuzz».

```Go
package main

import "fmt"

func main() {
    for i := 1; i < 100; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Println("FizzBuzz")
            continue
        }
        if i%3 == 0 {
            fmt.Println("Fizz")
            continue
        }
        if i%5 == 0 {
            fmt.Println("Buzz")
            continue
        }
        fmt.Println(i)
    }
}
```

## Массивы, срезы, карты

### 1.  Как обратиться к четвертому элементу массива или среза?

> Массивы нумеруются с нуля: `x[3]`

### 2.  Чему равна длина среза, созданного таким способом: *make([]int, 3, 9)*?

> 9 – длина массива, на который указывает срез  
3 – срез, который связан с массивом типа int, длиной 3  
Ответ: 3

### 3.  Дан массив:

```Go
x := [6]string{"a","b","c","d","e","f"}
```

что вернет вам *x[2:5]*?

> Ответ: [c d e]

### 4.  Напишите программу, которая находит самый наименьший элемент в этом списке:

```Go
x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
}
```

```Go
package main

import "fmt"

func main() {
    x := []int{
        48, 96, 86, 68,
        57, 82, 63, 70,
        37, 34, 83, 27,
        19, 97, 9, 17,
    }

    var total int = x[0]
    for _, value := range x {
        if total > value {
            total = value
        }
    }
    fmt.Println(total)
}
```

## Функции

### 1.  Функция *sum* принимает срез чисел и складывает их вместе. Как бы выглядела сигнатура этой функции?

```Go
func sum(xs []int) int
```

### 2.  Напишите функцию, которая принимает число, делит его пополам и возвращает *true* в случае, если исходное число чётное, и *false*, если нечетное. Например, *half(1)* должна вернуть *(0, false)*, в то время как *half(2)* вернет *(1, true)*.

```Go
package main

import "fmt"

func main() {
    i, flag := half(2)
    fmt.Println(i)
    fmt.Println(flag)
}

func half(number int) (int, bool) {
    if number%2 == 0 {
        return 0, true
    } else {
        return 1, false
    }
}
```

### 3.  Напишите функцию с переменным числом параметров, которая находит наибольшее число в списке.

```Go
package main

import "fmt"

func main() {
    i := max(2, 5, 2, 6)
    fmt.Println(i)
}

func max(args ...int) int {
    max := args[0]
    for _, v := range args {
        if max < v {
            max = v
        }
    }
    return max
}
```

### 4.  Используя в качестве примера функцию *makeEvenGenerator* напишите *makeOddGenerator*, генерирующую нечётные числа.

```Go
package main

import "fmt"

func main() {
    i := makeOddGenerator()
    fmt.Println(i())
    fmt.Println(i())
    fmt.Println(i())
    fmt.Println(i())
    fmt.Println(i())
}

func makeOddGenerator() func() uint {
    i := uint(1)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}
```

### 5.  Последовательность чисел Фибоначчи определяется как *fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2)*. Напишите рекурсивную функцию, находящую *fib(n)*.

```Go
package main

import "fmt"

func main() {
    i := fib(5)
    fmt.Println(i)
}

func fib(n uint) uint {
    if n == 1 || n == 2 {
        return 1
    }
    return fib(n-1) + fib(n-2)
}
```

### 6.  Что такое отложенный вызов, паника и восстановление? Как восстановить функцию после паники?

> В Go есть специальный оператор *defer*, который позволяет отложить вызов указанной функции до тех пор, пока не завершится текущая.  
Функция *panic* генерирует ошибку выполнения. Мы можем обрабатывать паники с помощью встроенной функции *recover*.

```Go
package main

import "fmt"

func main() {
    defer func() {    
        str := recover()
        fmt.Println(str)
    }()
    panic("PANIC")
}
```

## Указатели

### 1.  Как получить адрес переменной?

> Существует оператор `&`, который используется для получения адреса переменной.

### 2.  Как присвоить значение указателю?

> Используя указатель (`*int`), мы можем изменить значение оригинальной переменной.

### 3.  Как создать новый указатель?

> Функция `new` принимает аргументом тип, выделяет для него память и возвращает указатель на эту память.

### 4.  Какое будет значение у переменной *x* после выполнения программы:

```Go
func square(x *float64) {
    *x = *x * *x
}
func main() {
    x := 1.5
    square(&x)
}
```

> Ответ: 2.25

### 5.  Напишите программу, которая меняет местами два числа (*x := 1; y := 2; swap(&x, &y)* должно дать *x=2* и *y=1*).

```Go
package main

import "fmt"

func main() {
    x := 1
    y := 2
    fmt.Println(x)
    fmt.Println(y)
    swap(&x, &y)
    fmt.Println(x)
    fmt.Println(y)
}

func swap(x, y *int) {
    tmp := *x
    *x = *y
    *y = tmp
}
```

## Структуры и интерфейсы

### 1.  Какая разница между методом и функцией?

> Метод:

```Go
func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}
```

> Функция:

```Go
func circleArea(c *Circle) float64 {
    return math.Pi * c.r*c.r
}
```

> Разница в том, что при использовании метода не нужно использовать оператор  '&' (Go автоматически предоставляет доступ к указателю на *Circle* для этого метода).

### 2.  В каких случаях могут пригодиться встроенные (скрытые) поля?

> Поля структур представляют отношения принадлежности (включения).  
Например:

```Go
type Person struct {
    Name string
}
func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}
```

> Это будет работать, но мы можем захотеть создать другое отношение.

```Go
type Android struct {
    Person Person
    Model string
}
```

> Мы использовали тип (*Person*) и не написали его имя. Объявленная таким способом структура доступна через имя типа:

```Go
a := new(Android)
a.Person.Talk()
```

> Go поддерживает подобные отношения с помощью встраиваемых типов, также называемых анонимными полями.

```Go
type Android struct {
    Person
    Model string
}
```

> Но тут мы можем вызвать любой метод *Person* прямо из *Android*:

```Go
a := new(Android)
a.Talk()
```

### 3.  Добавьте новый метод *perimeter* в интерфейс *Shape*, который будет вычислять периметр фигуры. Имплементируйте этот метод для *Circle* и *Rectangle*.

```Go
package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	x, y, a, b float64
}

type Circle struct {
	x, y, r float64
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.a + r.b)
}

func (r *Rectangle) area() float64 {
	return r.a * r.b
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Shape interface {
	area() float64
	perimeter() float64
}

func Calculation(f func(Shape) float64, shapes ...Shape) []float64 {
	array := make([]float64, len(shapes))
	for i, s := range shapes {
		array[i] = f(s)
	}
	return array
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	r := Rectangle{x: 0, y: 0, a: 5, b: 3}
	value := Calculation(func(s Shape) float64 { return s.area() }, &c, &r)
	fmt.Println(value)
	value = Calculation(func(s Shape) float64 { return s.perimeter() }, &c, &r)
	fmt.Println(value)
}
```

## Многопоточность

### 1.  Как задать направление канала?

> Мы можем задать направление передачи сообщений в канале, сделав его только отправляющим или принимающим.  
Канал `c` будет только отправлять сообщение.  

```Go
func pinger(c chan<- string)
```

> Попытка получить сообщение из канала `c` вызовет ошибку компилирования.  
Канал `c` будет только принимать сообщение.  

```Go
func printer(c <-chan string)
```

### 2.  Напишите собственную функцию *Sleep*, используя *time.After*

```Go
package main

import (
    "fmt"
    "time"
)

func Sleep(sec int) {
    ch := time.After(time.Second * time.Duration(sec))
    for {
        select {
        case <-ch:
            fmt.Println("slept!")
            return
        }
    }
}

func main() {
    go Sleep(10)

    var input string
    fmt.Scanln(&input)
}
```

### 3.  Что такое буферизированный канал? Как создать такой канал с ёмкостью в 20 сообщений?

> Буферизированный канал – это канал с определенной ёмкостью. Работа осуществляется асинхронно — получение или отправка сообщения не заставляют стороны останавливаться.

```Go
c := make(chan int, 20)
```

## Пакеты и повтореное использование кода

### 1.  Зачем мы используем пакеты?

> Использование пакетов обусловлено тремя причинами:  
1. Снижение вероятности дублирование имён функций, что позволяет именам быть простыми и краткими  
2. Организация кода для упрощения поиска повторно используемых конструкций  
3. Ускорение компиляции, так как мы должны перекомпилировать только части программы. Несмотря на то, что мы используем пакет `fmt`, мы не должны перекомпилировать его при каждом использовании

### 2.  Чем отличаются программные сущности, названные с большой буквы? То есть, чем *Average* отличается от *average*?

> Любая сущность языка Go, которая начинается с заглавной буквы, означает, что другие пакеты и программы могут использовать эту сущность. Если бы мы назвали нашу функцию `average`, а не `Average`, то наша главная программа не смогла бы обратиться к ней.

### 3.  Что такое псевдоним пакета и как его сделать?

> Мы используем только краткое имя `math` когда мы обращаемся к функциям в нашем пакете. Если же мы хотим использовать оба пакета, то мы можем использовать псевдоним:

```Go
import m "golang-book/chapter11/math"

func main() {
    xs := []float64{1,2,3,4}
    avg := m.Average(xs)
    fmt.Println(avg)
}
```

> В этом коде `m` — псевдоним.

### 4.  Мы скопировали функцию *Average* из главы 7 в наш новый пакет. Создайте *Min* и *Max* функции для нахождения наименьших и наибольших значений в срезах дробных чисел типа *float64*.

```Go
package math

// Найти среднее в массиве чисел.
func Average(xs []float64) float64 {
    total := float64(0)
    for _, x := range xs {
        total += x
    }
    return total / float64(len(xs))
}

// Найти минимальное в массиве чисел.
func Min(xs []float64) (float64, bool) {
    if len(xs) == 0 {
        return 0.0, false
    }
    total := xs[0]
    for _, x := range xs {
        if total > x {
            total = x
        }
    }
    return total, true
}

// Найти максимальное в массиве чисел.
func Max(xs []float64) (float64, bool) {
    if len(xs) == 0 {
        return 0.0, false
    }
    total := xs[0]
    for _, x := range xs {
        if total < x {
            total = x
        }
    }
    return total, true
}
```

```Go
package main

import (
    "fmt"
    "golang-book/chapter11/math"
)

func main() {
    xs := []float64{6, 10, 2, 7}
    avg := math.Average(xs)
    fmt.Println(avg)
    avg, _ = math.Min(xs)
    fmt.Println(avg)
    avg, _ = math.Max(xs)
    fmt.Println(avg)
}
```

### 5.  Напишите документацию к функциям *Min* и *Max* из предыдущей задачи.

```
PS C:\Program Files\Go\src\golang-book\chapter11> go doc ./math
```

> func Average(xs []float64) float64  
func Max(xs []float64) (float64, bool)  
func Min(xs []float64) (float64, bool)  

```
PS C:\Program Files\Go\src\golang-book\chapter11> go doc ./math Min
```

> package math // import "golang-book/chapter11/math"  
func Min(xs []float64) (float64, bool)  
Найти минимальное в массиве чисел.

```
PS C:\Program Files\Go\src\golang-book\chapter11> go doc ./math Max
```

> package math // import "golang-book/chapter11/math"  
func Max(xs []float64) (float64, bool)  
Найти максимальное в массиве чисел.

## Тестирование

### 1.  Написать хороший набор тестов не всегда легко, но даже сам процесс их написания, зачастую, может выявить много проблем для первой реализации функции. Например, что произойдет с нашей функцией *Average*, если ей передать пустой список (*[]float64{}*)? Как нужно изменить функцию, чтобы она возвращала *0* в таких случаях?

```
PS C:\Program Files\Go\src\golang-book\chapter11\math> go test
```
> --- FAIL: TestAverage (0.00s)  
    math_test.go:9: Expected 1.5, got  NaN  
FAIL  
exit status 1  
FAIL    golang-book/chapter11/math      0.366s  

```Go
func Average(xs []float64) float64 {
    if len(xs) == 0 {
        return 0.0
    }
    total := float64(0)
    for _, x := range xs {
        total += x
    }
    return total / float64(len(xs))
}
```

### 2.  Напишите серию тестов для функций *Min* и *Max* из предыдущей главы.

```Go
package math

import "testing"

type testpair struct {
	values  []float64
	min     float64
	max     float64
	average float64
}

var tests = []testpair{
	{[]float64{6, 10, 2, 7}, 2.0, 10.0, 6.25},
	{[]float64{-8, 1, 8, -2, 4, -9}, -9.0, 8.0, -1.0},
	{[]float64{-1, 1}, -1.0, 1.0, 0.0},
	{[]float64{}, 0.0, 0.0, 0.0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range tests {
		v, ok := Min(pair.values)
		if !ok {
			t.Error("Len array 0")
		}
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range tests {
		v, ok := Max(pair.values)
		if !ok {
			t.Error("Len array 0")
		}
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}
```
