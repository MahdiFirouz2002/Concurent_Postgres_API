package concurentmodel

import (
	"fmt"
	"sync"
	"time"
	"users/common"
	"users/database"
	"users/jsonReader"
)

const workerCount = 10
const channelBuffer = 10

func InsertUsersConcurently() {
	users, err := jsonReader.UnmarshalData_from_json()
	if err != nil {
		panic(err)
	}

	connectToDb_err := database.ConnectToDB()
	if connectToDb_err != nil {
		panic("something went wrong with connection!")
	}

	defer database.Close_database()

	var wg sync.WaitGroup
	userData_in_chan := make(chan common.User, channelBuffer)
	startTime := time.Now()

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go userInsertWorker(i, userData_in_chan, &wg)
	}

	for _, user := range users {
		userData_in_chan <- user
	}

	close(userData_in_chan)
	wg.Wait()

	fmt.Printf("inserting users took: %fm \n", time.Since(startTime).Minutes())
}

func userInsertWorker(id int, userChan chan common.User, wg *sync.WaitGroup) {
	defer wg.Done()

	for userData := range userChan {
		insertErr := database.InsertUser(userData)
		if insertErr != nil {
			fmt.Printf("go routin: %d failed to insert userId: %s, error: %s \n", id, userData.Id, insertErr.Error())
		} else {
			fmt.Printf("go routin: %d successfully inserted userId: %s \n", id, userData.Id)
		}
	}
}
