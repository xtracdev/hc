package main

import (
	"github.com/gorilla/mux"
	"github.com/samalba/dockerclient"
	"log"
	"net/http"
	"encoding/json"
)

var docker *dockerclient.DockerClient

type ContainerStatus struct {
	Name string
	Image string
	Status string
}

func containerHealth(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	containerName := vars["container-name"]

	containers, err := docker.ListContainers(false, false, "")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	for _,c := range containers {
		for _, name := range c.Names {
			if len(name) >= 1 && name[0] == '/' {
				name = name[1:]
			}

			if name == containerName {
				status := ContainerStatus {
					Name: c.Names[0],
					Image: c.Image,
					Status: c.Status,
				}

				bytes, err := json.Marshal(&status)
				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
					return
				}

				rw.Write(bytes)
				return
			}
		}
	}

	http.Error(rw, "", http.StatusNotFound)

}

func allHealth(rw http.ResponseWriter, req *http.Request) {
	containers, err := docker.ListContainers(false, false, "")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var healths []ContainerStatus
	for _, c := range containers {
		//Don't include the health check container as it won't serve
		//traffic if it is unhealthy.
		if c.Image == "xtracdev/hc" {
			continue
		}

		status := ContainerStatus {
			Name: c.Names[0],
			Image: c.Image,
			Status: c.Status,
		}

		healths = append(healths, status)

	}

	bytes, err := json.Marshal(healths)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Write(bytes)

}

func main() {

	log.Println("Create docker client")
	var err error

	docker, err = dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mount routes")
	r := mux.NewRouter()
	r.HandleFunc("/health/{container-name}", containerHealth)
	r.HandleFunc("/health", allHealth)
	http.Handle("/", r)

	log.Println("Serve HTTP")
	http.ListenAndServe(":5000", nil)
}
