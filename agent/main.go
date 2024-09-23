package main

import (
	"fmt"
	"os"

	"github.com/reivaxt/xtuples/common"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type Task struct {
	Config       *common.TaskConfig
	Redis        *redis.Client
	JobConn      *grpc.ClientConn
	JobService   *JobServiceClient
	TakeScript   *redis.Script
	MutateScript *redis.Script
}

type Taken struct {
	Key     string
	Success bool
}

var Ctx Task

func AgentExec(job_name string) {

	// Take a key from the xuple
	taken, err := Ctx.Take(job_name)
	if err != nil {
		fmt.Println("Failed to take key:", err)
		return
	}

	// call something with the key
	if taken.Success {
		fmt.Printf("JobCall( %s )\n", taken.Key)
		result := JobCall(taken.Key)

		// Mutate the token
		Ctx.Mutate(job_name, result, taken.Key)
	} else {
		fmt.Println("Condition not met")
	}
}

func AgentLoop() {
	job_name := os.Getenv("JOB_NAME")
	if job_name != "" {
		for {
			AgentExec(job_name)
		}
	} else {
		fmt.Println("No JOB_NAME specified")
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			fmt.Println("No .env file found")
		}
	}

	// Connect to xtuples memory
	redisConn := connectToKeyDB()
	defer redisConn.Close()
	Ctx.Redis = redisConn

	// Set the redis scripts
	setRedisScripts()

	// connect to grpc server
	grpcConn := connectToGrpc()
	defer grpcConn.Close()
	grpc := NewJobServiceClient(grpcConn)
	Ctx.JobService = &grpc

	// Read the config
	configfile := os.Getenv("SERVICE_CONFIG")
	if configfile == "" {
		configfile = "./config.yaml"
	}

	Ctx.Config = common.ParseConfig(configfile)

	AgentLoop()
}
