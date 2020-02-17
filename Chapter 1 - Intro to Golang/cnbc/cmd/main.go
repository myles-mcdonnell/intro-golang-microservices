package main

import (
	"fmt"
	"github.com/myles-mcdonnell/cnbc"
	"io/ioutil"
	"net/http"
)

func main() {
	memo := cnbc.New(GetBody)

	for i := 0;i<10000;i++ {
		_, err := memo.Get("https://www.bbc.com/news")
		if err!=nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println("Memo.Get() called 1000 times.")
}

func GetBody(url string) (interface{}, error) {
	fmt.Printf("HTTP GET %v\r\n", url)
	resp, err := http.Get(url)
	if err!=nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		return nil, err
	}

	return body, nil
}


