package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

// we're creating a wait group here to WAIT for each go routine to finish. the for loop quickly
// iterates over every index, and adds a counter for each i. then when each db call is done it decrements this counter
// when the counter is back to 0, we can continue
// without go routines, each database call would have to wait for the previous one before it
func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	// you might think, what if theres a race condition and we done before the next add? essentially make the counter 0 before we're actually done
	// since we call wait ONLY after all adds have been added to the group, this is impossible
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
	// simulate getting data from a real database
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("The result from the database is: %v\n", dbData[i])
	// since each call is simultaneous sometimes we can get race conditions where each call tries to write to the same
	// location. to get around this, we check a mutex. if the mutex is already locked, that means another call is working
	// and you must wait your turn
	mutex.Lock()
	results = append(results, dbData[i])
	mutex.Unlock()
	wg.Done()
}
