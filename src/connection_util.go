package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func PingServer(ip string, port int, ctx *context.Context) bool {
	address := "http://" + ip + ":" + strconv.Itoa(port)

	fmt.Println("Pinging: " + address)
	req, err := http.NewRequestWithContext(*ctx, "GET", address, nil)
	if err != nil {
		fmt.Println("Create request failed")
		return false
	}
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Connection failed")
		return false
	}
	defer resp.Body.Close()

	fmt.Println("Connection success for " + ip)
	return true
}

func GetWorkingServer(c *Configurator) (string, error) {
	var wg sync.WaitGroup
	wg.Add(len(c.IPS))

	var mutex sync.Mutex

	var foundedIp string = ""
	var closeCh chan bool = make(chan bool, len(c.IPS))

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
			status := PingServer(ipAddr, c.Port, &ctx)
			if status {
				mutex.Lock()
				foundedIp = ipAddr
				mutex.Unlock()
			}
			closeCh <- status
			wg.Done()
		}(ip)
	}

	wg.Wait()
	if foundedIp == "" {
		return "", fmt.Errorf("No servers response")
	}
	return foundedIp, nil
}
