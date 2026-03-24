package iwanCore

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func GetResponse(url string, ctx *context.Context, request string) (*http.Response, bool) {
	address := url + "?name=" + request

	Log("Pinging: " + address)

	req, err := http.NewRequestWithContext(*ctx, "GET", address, nil)
	if err != nil {
		Log("Create request failed")
		return nil, false
	}
	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		Log("Connection failed")
		return nil, false
	}

	Log("Connection success for " + url)
	return resp, true
}

func TryAllServers(c *Configurator, request string) (IwanResponse, error) {
	var wg sync.WaitGroup
	wg.Add(len(c.URLS))

	var mutex sync.Mutex
	var closeCh chan bool = make(chan bool, len(c.URLS))

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

	for _, url := range c.URLS {
		go func(urlAdr string) {
			resp, status := GetResponse(urlAdr, &ctx, request)
			if status {
				mutex.Lock()

				content, err := io.ReadAll(resp.Body)
				if err != nil {
					panic(err.Error())
				}

				var iwanResponse IwanResponse
				unmarshalError := json.Unmarshal(content, &iwanResponse)
				if unmarshalError != nil {
					Log(unmarshalError.Error())
					os.Exit(1)
				}

				if iwanResponse.Status == "ERR" {
					Log("Server returned an error: " + iwanResponse.Content)
				} else {
					res = iwanResponse
					closeCh <- true
				}

				mutex.Unlock()
			}
			wg.Done()
		}(url)
	}

	wg.Wait()
	if res == (IwanResponse{}) {
		return res, fmt.Errorf("No servers response")
	}
	return res, nil
}
