package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type S struct {
	str string
	sync.Mutex
}

func main() {
	/* // example 1
	go infinite() */

	/* // example 2
	go infinite()
	for i := 0; i < 10; i++ {
		fmt.Println("Hi from main!")
		time.Sleep(1 * time.Second)
	} */

	/* // example 3 - waitgroup
	wg := &sync.WaitGroup{}
	wg.Add(5)

	go infiniteWithWg("t-1", wg)
	go infiniteWithWg("t-2", wg)
	go infiniteWithWg("t-3", wg)
	go infiniteWithWg("t-4", wg)
	go infiniteWithWg("t-5", wg)
	wg.Wait() */

	// example 4 - deadlock
	/* wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second) // time.Sleep does not cause a thread to be "blocked"
		wg.Wait()                   // wg.Wait, does though.
	}()
	wg.Wait() // as all threads are blocked, we're in "deadlock" and Go will fatally error */

	/* // example 5 - unsafe accessing of structs
	s := &S{I: 0, str: "A"}
	for {
		go addOne("t-1", s)
		go addOne("t-2", s)
		go addOne("t-3", s)
		go addOne("t-4", s)
		go addOne("t-5", s)
		time.Sleep(5 * time.Second)
	} */

	// example 6 - mutex accessing of structs
	/* s := &S{str: "A"}
	for {
		go addOne("t-1", s)
		go addOne("t-2", s)
		go addOne("t-3", s)
		go addOne("t-4", s)
		go addOne("t-5", s)
		time.Sleep(5 * time.Second)
	} */
}

func infinite() {
	for {
		fmt.Println("Hi from infinite loop!")
		time.Sleep(1 * time.Second)
	}
}

func infiniteWithWg(name string, wg *sync.WaitGroup) {
	for {
		i := rand.Intn(10)
		if i == 0 {
			fmt.Printf("%s: i = %d -- I'm done!\n", name, i)
			wg.Done()
			return
		}

		fmt.Printf("%s: i = %d\n", name, i)
		time.Sleep(1 * time.Second)
	}
}

func addOne(name string, s *S) {
	s.Lock()
	defer s.Unlock()

	before := s.str
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

	lastLetter := before[len(before)-1]
	nextLetter := string(lastLetter + 1)

	s.str = s.str + nextLetter

	fmt.Printf("%s adding next letter %s: new_value=%s\n", name, nextLetter, s.str)
}
