package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func GetResponse(ip string, port int, ctx *context.Context, request string) (*http.Response, bool) {
	address := "http://" + ip + ":" + strconv.Itoa(port) + "?name=" + request
	fmt.Println("Pinging: " + address)
	req, err := http.NewRequestWithContext(*ctx, "GET", address, nil)
	if err != nil {
		fmt.Println("Create request failed")
		return nil, false
	}
	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Connection failed")
		return nil, false
	}

	fmt.Println("Connection success for " + ip)
	return resp, true
}

func TryAllServers(c *Configurator, request string) (IwanResponse, error) {
	var wg sync.WaitGroup
	wg.Add(len(c.IPS))

	var mutex sync.Mutex
	var closeCh chan bool = make(chan bool, len(c.IPS))

	var res IwanResponse

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for status := range closeCh {
			if status {
				cancel()
				return
			}
		}
	}()

	for _, ip := range c.IPS {
		go func(ipAddr string) {
			resp, status := GetResponse(ipAddr, c.Port, &ctx, request)
			if status {
				mutex.Lock()

				content, err := io.ReadAll(resp.Body)
				if err != nil {
					panic(err.Error())
				}

				var iwanResponse IwanResponse
				unmarshalError := json.Unmarshal(content, &iwanResponse)
				if unmarshalError != nil {
					panic(unmarshalError.Error())
				}

				if iwanResponse.Status == "ERR" {
					fmt.Println("Server returned an error: " + iwanResponse.Content)
				} else {
					res = iwanResponse
					closeCh <- true
				}

				mutex.Unlock()
			}
			wg.Done()
		}(ip)
	}

	wg.Wait()
	if res == (IwanResponse{}) {
		return res, fmt.Errorf("No servers response")
	}
	return res, nil
}
