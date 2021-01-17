package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Config struct {
	DB DB `json:"db"`
}

type DB struct {
	DBName string `json:"dbname"`
	Users  []User
}

type User struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	ps()

	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	//m := new(Dispatch)
	//var m interface{}
	var c Config
	json.Unmarshal(file, &c)
	fmt.Printf("Results: %v\n", c)
}

func ps() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	cli.ContainerExecAttach(ctx, execID, config)
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", strings.Join(container.Names, ","), container.Image)
	}
}
