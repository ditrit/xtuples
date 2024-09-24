package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/supertokens/supertokens-golang/supertokens"
)

//go:embed web/dist
var efs embed.FS

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			fmt.Println("No .env file found")
		}
	}

	err = supertokens.Init(SuperTokensConfig)
	if err != nil {
		panic(err.Error())
	}

	router := gin.New()
	dist, err := fs.Sub(efs, "web/dist")
	if err != nil {
		log.Fatalf("dist file server error: %v", err)
	}

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: append([]string{"content-type"},
			supertokens.GetAllCORSHeaders()...),
		AllowCredentials: true,
	}))

	// Adding the SuperTokens middleware
	router.Use(func(c *gin.Context) {
		supertokens.Middleware(http.HandlerFunc(
			func(rw http.ResponseWriter, r *http.Request) {
				c.Next()
			})).ServeHTTP(c.Writer, c.Request)
		// we call Abort so that the next handler in the chain is not called, unless we call Next explicitly
		c.Abort()
	})

	router.StaticFS("/", http.FS(dist))
	router.Run(":3000")
}
