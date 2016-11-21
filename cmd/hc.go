package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/samalba/dockerclient"
	"log"
	"net/http"
)

var docker *dockerclient.DockerClient

func containerHealth(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("ok"))
}

func allHealth(rw http.ResponseWriter, req *http.Request) {
	containers, err := docker.ListContainers(false, false, "")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, c := range containers {
		status := fmt.Sprintf("%s %s: %s\n", c.Names[0], c.Image, c.Status)
		rw.Write([]byte(status))
	}

}

func main() {

	log.Println("Create docker client")
	var err error

	docker, err = dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	r := mux.NewRouter()
	r.HandleFunc("/health/{container-name}", containerHealth)
	r.HandleFunc("/health", allHealth)
	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}
