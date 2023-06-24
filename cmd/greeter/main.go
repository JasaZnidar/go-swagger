package main

import (
	"Greeter/gen/restapi"
	"Greeter/gen/restapi/operations"
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"strings"
)

//"go.mongodb.org/mongo-driver/mongo/readpref"

var portFlag = flag.Int("port", 3000, "Port to run this server on.")

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewGreeterAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	flag.Parse()
	server.Port = *portFlag

	// GET hello?name
	api.GetHelloHandler = operations.GetHelloHandlerFunc(func(params operations.GetHelloParams) middleware.Responder {
		name := swag.StringValue(params.Name)

		if name == "" {
			name = "World"
		}

		greeting := fmt.Sprintf("Hello, %s!", name)
		return operations.NewGetHelloOK().WithPayload(greeting)
	})

	// GET time
	api.GetTimeHandler = operations.GetTimeHandlerFunc(func(params operations.GetTimeParams) middleware.Responder {
		return operations.NewGetTimeOK().WithPayload((time.Now()).String())
	})

	// GET user?name&&pass
	api.GetUserHandler = operations.GetUserHandlerFunc(func(params operations.GetUserParams) middleware.Responder {
		// extract parameters
		name := swag.StringValue(params.Name)
		pass := swag.StringValue(params.Pass)
		fmt.Printf("name = %s\npass = %s\n", name, pass)

		// connect to DB
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
		if err != nil {
			errStr := err.Error()
			payload := operations.GetUserInternalServerErrorBody{Error: &errStr}
			return operations.NewGetUserInternalServerError().WithPayload(&payload)
		}

		// find user
		var result bson.Raw
		collection := client.Database("users").Collection("users")
		result, err = collection.FindOne(ctx, bson.D{{"user", name}}).DecodeBytes()
		if err != nil {
			errStr := err.Error()
			payload := operations.GetUserInternalServerErrorBody{Error: &errStr}
			return operations.NewGetUserInternalServerError().WithPayload(&payload)
		}

		// check password
		hash := fmt.Sprint(result.Lookup("pass"))
		hash = hash[1 : len(hash)-1]

		// return OK
		payload := operations.GetUserOKBody{Name: &name, Pass: hash == pass}
		return operations.NewGetUserOK().WithPayload(&payload)
	})

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
