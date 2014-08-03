package main

import (
	"sync"
	// "fmt"
	// "io/ioutil"
	// "net/http"
)

func main() {
	oReqChannel := make(chan *httpRequest, 100)
	oRepChannel := make(chan string, 100)

	// lauching workers
	for i := 0; i < 100; i++ {
		go httpWorker(oReqChannel, oRepChannel)
	}

	for i := 0; i < 100; i++ {
		go analyze(oRepChannel)
	}

	// fmt.Println("receiving " + oReq.sUrl)

	for i := 0; i < 5000; i++ {
		oTestRequest := &httpRequest{"http://www.twenga.fr/"}
		// oTestRequest2 := &httpRequest{"http://www.google.fr/"}

		// go httpWorker(oReqChannel)
		oReqChannel <- oTestRequest
		// oReqChannel <- oTestRequest2
	}

	var oWait sync.WaitGroup
	oWait.Add(1)
	oWait.Wait()

}
