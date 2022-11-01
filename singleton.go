/*
Singleton allows us to make sure that after a variable has been properly initialized,
all new invocations won't re instantiate it, instead they'll receive the first value.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (Database) CreateSingleConn() {
	fmt.Println("creating...")
	time.Sleep(2 * time.Second)
	fmt.Println("created!")
}

var db *Database
var lock sync.Mutex

func getDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("creating new database connection")
		db = &Database{}
		db.CreateSingleConn()
	} else {
		fmt.Println("database already instantiated")
	}
	return db
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}
	wg.Wait()
}
