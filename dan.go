package dan

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

type dan struct {
	name string
}

func (d dan) String() string {
	return "Name is: " + d.name
}

func Dosomething(a chan int) {
	log.Printf("Starting Dosomething")

	for i := range a {
		fmt.Println("hello: " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}

	log.Printf("Exiting Dosomething")
}

func do(somechan chan int) {
	time.Sleep(3 * time.Second)

	go func() {
		for x := 0; x < 10; x++ {
			somechan <- x
		}
	}()

	time.Sleep(3 * time.Second)

	close(somechan)
}

func dothing(a int) (i int, e error) {
	e = nil
	i = a
	if a > 10 {
		return 0, errors.New("my error")
	}
	return
}

func main2() {
	if a, err := dothing(5); err != nil {
		log.Fatal("got error: ", err, a)
	}

	test := dan{name: "hello"}
	fmt.Println(test)

	somechan := make(chan int, 100)
	defer close(somechan)

	go Dosomething(somechan)

	time.Sleep(3 * time.Second)
	go do(somechan)

	log.Printf("Leaving main2")
}

func main() {
	main2()
	time.Sleep(3 * time.Second)
	log.Printf("Leaving main")
}
