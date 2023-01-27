package main

import (
	"fmt"
	"sync"
)

type employee struct {
	name string
	age  int
}

func print(s employee, wg *sync.WaitGroup) {
	fmt.Println(s.name, " ", s.age)
	wg.Done()
}

func process(s employee, ch chan employee, wg *sync.WaitGroup) {
	s.age = s.age + 1
	ch <- s
	wg.Done()
}
func main() {
	name := "Senjuti"
	ch := make(chan employee, 5)
	var wg sync.WaitGroup
	for i := 23; i <= 32; i = i + 1 {
		s := employee{
			name: name,
			age:  i,
		}
		wg.Add(1)
		go process(s, ch, &wg)

		wg.Add(1)
		go print(<-ch, &wg)
	}

	wg.Wait()

	fmt.Printf("All Done")
}
