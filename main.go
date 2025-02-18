package main

import (
	"fmt"
	"net"
	"sync"
)

func poortSniffer(ip string, poorten []int, wg *sync.WaitGroup) {
	for _, poort := range poorten {
		wg.Add(1)
		go func(poort int) {
			defer wg.Done()
			addr := fmt.Sprintf("%s:%d", ip, poort)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				return
			}
			defer conn.Close()
			fmt.Printf("Poort %d op %s is open!\n", poort, ip)
		}(poort)
	}
}

func main() {
	ip := "127.0.0.1"
	poorten := []int{22, 80, 443, 8080}
	var wg sync.WaitGroup

	poortSniffer(ip, poorten, &wg)
	wg.Wait()
}
