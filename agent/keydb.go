package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/reivaxt/xtuples/common"
)

func connectToKeyDB() *redis.Client {
	host := os.Getenv("KEYDB_HOST")
	if host == "" {
		host = "localhost"
	}
	//host := "keydb"
	port := os.Getenv("KEYDB_PORT")
	if port == "" {
		port = "6379"
	}
	password := os.Getenv("KEYDB_PASSWORD")

	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0, // Use the default database
	})

	// Ping the server to check if the connection is successful
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (t Task) Take(agent string) (Taken, error) {
	var taken Taken
	jsonFilter, err := t.Config.AgentByNames[agent].JsonFilter()
	fmt.Println(("exec take with jsonFilter " + jsonFilter))
	if err != nil {
		fmt.Println("Failed to convert filter to JSON:", err)
	} else {
		res, err := t.TakeScript.Run(context.Background(), t.Redis, t.Config.Keys, jsonFilter).Slice()
		if err != nil {
			fmt.Println("Failed to run script:", err)
		}
		fmt.Println("Take result: ")
		fmt.Println(res[0])
		fmt.Println(res[1])
		taken.Success = res[0].(int64) == 1
		taken.Key = res[1].(string)
	}
	return taken, err
}

func (t Task) Mutate(agent string, success bool, key string) error {
	jsonSuccessConfig, err := t.Config.AgentByNames[agent].JsonOnSuccess()
	fmt.Println("exec mutate with jsonSuccessConfig" + jsonSuccessConfig)
	fmt.Println("exec mutate with key" + key)
	fmt.Println("exec mutate with success" + fmt.Sprint(success))
	numSuccess := 0
	if success {
		numSuccess = 1
	}
	if err != nil {
		fmt.Println("Failed to convert filter to JSON:", err)
	} else {
		res, err := t.MutateScript.Run(context.Background(), t.Redis, t.Config.Keys, numSuccess, jsonSuccessConfig, key).Result()
		if err != nil {
			fmt.Println("Failed to run script:", err)
		}
		fmt.Println("Mutate result: ")
		fmt.Println(res)
	}
	return err
}

func setRedisScripts() {
	// Read the Lua script for 'take' command from file
	takeScriptPath := "./redis/scripts/take.lua"
	takeScript, err := common.ReadFromFile(takeScriptPath)
	Ctx.TakeScript = redis.NewScript(takeScript)
	if err != nil {
		fmt.Println("Failed to read script from file:", err)
	} else {
		fmt.Println("Read take script from file" + takeScriptPath)
	}

	// Read the Lua script for 'take' command from file
	mutateScriptPath := "./redis/scripts/mutate.lua"
	mutateScript, err := common.ReadFromFile(mutateScriptPath)
	Ctx.MutateScript = redis.NewScript(mutateScript)
	if err != nil {
		fmt.Println("Failed to read script from file:", err)
	}
}
