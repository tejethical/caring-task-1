package main

import (
	"caring-task-1/lib"
	cli_args "caring-task-1/pkgs/cli-args"
	"caring-task-1/pkgs/sorting"
	tcp_client "caring-task-1/pkgs/tcp-net/client"
	"caring-task-1/pkgs/tcp-net/server"
	"caring-task-1/pkgs/work"
	"fmt"
	"log"
)

//TODO:Fault Tolerance
var workerPorts = []string{"91", "92", "93"}

func main() {
	args := cli_args.Get()
	switch args.Role {
	case "master":

		clients, err := tcp_client.GetClients(workerPorts)
		if err != nil {
			log.Fatal("Error connecting to workers: ", err)
		}

		unsortedData, err := lib.GetDataFromFile(args.FilePath)
		if err != nil {
			log.Fatal("Error reading names from file", err)
		}

		chunks, err := sorting.GetSliceChunks(unsortedData, 3)
		if err != nil {
			log.Fatal(err)
		}

		result := work.AssignWork(clients, chunks)

		sortedSlice := sorting.MergeSlices(result[0].Data, result[1].Data, result[2].Data)

		fmt.Println(sortedSlice)
		break
	case "worker":
		err := server.NewWorker(args.Port)
		if err != nil {
			log.Fatal("Failed to configure node", err)
		}
		break

	default:
		log.Fatal("Invalid role")
	}
}
