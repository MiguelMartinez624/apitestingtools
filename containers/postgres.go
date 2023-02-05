package containers

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
)

type PostgresContainer struct {
	container *tc.Container
	URI       string
}

var postgresContainer *PostgresContainer

func NewPostgresContainer() *PostgresContainer {

	if postgresContainer != nil {
		return postgresContainer
	}

	log.Println("Starting postgres container...")
	postgresPort := nat.Port("5432/tcp")

	postgres, err := tc.GenericContainer(context.Background(),
		tc.GenericContainerRequest{
			ContainerRequest: tc.ContainerRequest{
				Image:        "postgres",
				ExposedPorts: []string{postgresPort.Port()},
				Env: map[string]string{
					"POSTGRES_PASSWORD": "pass",
					"POSTGRES_USER":     "user",
				},
				WaitingFor: wait.ForAll(
					wait.ForLog("database system is ready to accept connections"),
					wait.ForListeningPort(postgresPort),
				),
			},
			Started: true, // auto-start the container
		})
	if err != nil {
		log.Fatal("start:", err)
	}
	hostPort, err := postgres.MappedPort(context.Background(), postgresPort)
	if err != nil {
		log.Fatal("map:", err)
	}

	postgresURLTemplate := "postgres://user:pass@localhost:%s?sslmode=disable"
	postgresURL := fmt.Sprintf(postgresURLTemplate, hostPort.Port())
	postgresContainer = &PostgresContainer{
		container: &postgres,
		URI:       postgresURL,
	}
	return postgresContainer

}
