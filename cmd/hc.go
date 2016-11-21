package main

import (
	"log"
	"github.com/samalba/dockerclient"
	"net/http"
)

var docker *dockerclient.DockerClient

func healthHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("ok"))
}

func main() {

	log.Println("Create docker client")
	var err error

	docker, err = dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":5000", mux)
}
