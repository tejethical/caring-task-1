package work

import (
	"caring-task-1/pkgs/tcpNet/client"
	"log"
	"sync"
)

type WorkPayload struct {
	Data []string
}

func AssignWork(clients []client.Client, chunks [][]string) []WorkPayload {
	result := make([]WorkPayload, 3)
	wg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(j int) {
			defer func() {
				wg.Done()
			}()
			err := client.SendPayload(clients[j], WorkPayload{Data: chunks[j]})
			if err != nil {
				log.Fatal("Error sending payload to ", clients[j].Addr)
			}
			client.ReadResponse(clients[j], &result[j])
		}(i)

	}
	wg.Wait()

	return result
}
