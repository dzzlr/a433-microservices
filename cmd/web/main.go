package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	vault "github.com/hashicorp/vault/api"
	"github.com/nothinux/karsajobs/pkg/models/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	jobs    *mongodb.JobModel
	counter *int
}

func main() {
	client, err := openDB()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)

	db := client.Database("karsajobs")
	count := 10

	app := &application{
		jobs: &mongodb.JobModel{
			DB: db,
		},
		counter: &count,
	}

	log.Printf("application running on port %s", os.Getenv("APP_PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), app.routes()); err != nil {
		log.Fatal(err)
	}
}

func getMongoCredential() (string, string, error) {
	config := vault.DefaultConfig()
	config.Address = os.Getenv("VAULT_ADDR")

	client, err := vault.NewClient(config)
	if err != nil {
		return "", "", err
	}

	// Kubernetes auth
	jwt, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		return "", "", err
	}

	loginData := map[string]interface{}{
		"role": "karsajobs-role",
		"jwt":  string(jwt),
	}

	resp, err := client.Logical().Write("auth/kubernetes/login", loginData)
	if err != nil {
		return "", "", err
	}

	client.SetToken(resp.Auth.ClientToken)

	// Request dynamic DB credential
	secret, err := client.Logical().Read("database/creds/karsajobs-role")
	if err != nil {
		return "", "", err
	}

	username := secret.Data["username"].(string)
	password := secret.Data["password"].(string)

	return username, password, nil
}

func openDB() (*mongo.Client, error) {
	user, pass, err := getMongoCredential()
	if err != nil {
		return nil, err
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:27017/?authsource=admin", user, pass, os.Getenv("MONGO_HOST"))))
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	return client, nil
}
