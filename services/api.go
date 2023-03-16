package services

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func MulRequst(urls []string) {
	var baseUrl string = "http://127.0.0.1:8086/sanhua/api/dashboard"
	for {
		go func() {
			for _, v := range urls {
				data, err := Get(fmt.Sprintf("%s/%s", baseUrl, v))
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(string(data))
			}
		}()
		time.Sleep(time.Second * 20)
	}
}

func Get(apiURL string) (res []byte, err error) {
	req, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}

	defer req.Body.Close()

	resp, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Read body failed: ", err)
		return nil, err
	}
	return resp, nil
}
