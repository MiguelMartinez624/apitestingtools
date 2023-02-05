package containers

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"strconv"
	"time"
)

type MongoDBContainer struct {
	container *tc.Container
	URI       string
}

const defaultMongoDBPort = 27017

func NewMongoDBContainer() *MongoDBContainer {
	log.Println("Starting mongoDB container...")
	mongoPort, _ := nat.NewPort("", strconv.Itoa(defaultMongoDBPort))
	timeout := 5 * time.Minute

	postgres, err := tc.GenericContainer(context.Background(),
		tc.GenericContainerRequest{
			ContainerRequest: tc.ContainerRequest{
				Image:        "mongo:4.4.3",
				ExposedPorts: []string{mongoPort.Port()},
				Env: map[string]string{
					"MONGO_INITDB_ROOT_PASSWORD": "pass",
					"MONGO_INITDB_ROOT_USERNAME": "user",
				},
				WaitingFor: wait.ForLog("Waiting for connections").WithStartupTimeout(timeout),
			},

			Started: true, // auto-start the container
		})
	if err != nil {
		log.Fatal("start:", err)
	}
	hostPort, err := postgres.MappedPort(context.Background(), mongoPort)
	if err != nil {
		log.Fatal("map:", err)
	}
	host, err := postgres.Host(context.Background())
	if err != nil {
		log.Fatal("map:", err)
	}

	databaseAuth := fmt.Sprintf("%s:%s@", "user", "pass")
	databaseHost := fmt.Sprintf("%s:%d", host, uint(hostPort.Int()))
	postgresURL := fmt.Sprintf("mongodb://%s%s/?connect=direct", databaseAuth, databaseHost)
	return &MongoDBContainer{
		container: &postgres,
		URI:       postgresURL,
	}

}
