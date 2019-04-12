package main

import (
	"fmt"
	"os"
)

func plus(a int, b int) int {
	return a + b
}

/*
Person struct
*/
type Person struct {
	age  int
	Name string
}

/*
* Go has no such thing as classes, but this is how you declare methods for structs
*
 */
func (p *Person) setAge(age int) {
	p.age = age
}

func (p *Person) getAge() int {
	return p.age
}

func main() {
	fmt.Println("Hello World")
	fmt.Printf("%d", plus(2, 3))
	x := 5
	satish := Person{age: 19, Name: "Satish"}
	fmt.Println(x, satish.getAge()) //notice how the satish has a .getAge()
	os.Create("./data.txt")
	file, err := os.OpenFile("./data.txt", os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("Thre was an error opening the file!")
		os.Exit(-1)
	}

	fmt.Fprintln(file, "Putting datas into file")
	file.Close()
	file, err = os.OpenFile("./data.txt", os.O_RDWR, 0777)
	BUF := make([]byte, 16)

	bytesread, err := file.Read(BUF)
	fmt.Print(bytesread, " bytes were read.")
	fmt.Println(BUF)
	fmt.Println(string(BUF))
	file.Close()

	names := make([]string, 16)
	for x := 0; x < 10; x++ {
		names[x] = "e"
	}

	i := 0
	for i < 10 {
		fmt.Println(names[i])
		i++
	}

	cus := Customer{}
	cus.age = 20
	cus.yell()
}
