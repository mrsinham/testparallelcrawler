package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type httpRequest struct {
	sUrl string
	// retry int
}

func httpWorker(oRequestChan chan *httpRequest, oResponseChanne chan string) {
	for {
		select {
		case oReq := <-oRequestChan:
			fmt.Println("receiving " + oReq.sUrl)
			client := &http.Client{}

			oResponse, err := client.Get(oReq.sUrl)
			fmt.Println("tt")
			fmt.Println(oResponse.Status)
			if err != nil {
				fmt.Println(err.Error())
				panic(err.Error())
			}
			fmt.Println("receiving " + oReq.sUrl)

			body, errorRead := ioutil.ReadAll(oResponse.Body)
			if errorRead != nil {
				panic(errorRead.Error())
			}
			oResponseChanne <- string(body)
		}
	}

	// }

}
