package vcap

import (
	"encoding/json"
	"os"
	"time"
)

type VCAP struct {
	Application struct {
		ID            string   `json:"application_id"`
		Name          string   `json:"application_name"`
		Version       string   `json:"application_version"`
		InstanceID    string   `json:"instance_id"`
		InstanceIndex int      `json:"instance_index"`
		Host          string   `json:"host"`
		Port          int      `json:"port"`
		Users         string   `json:"users"`
		URIs          []string `json:"application_uris"`
		Limits        struct {
			Memory int `json:"mem"`
			Disk   int `json:"disk"`
			Files  int `json:"fds"`
		} `json:"limits"`
		Started *time.Time `json:"started_at_timestamp"`
		State   *time.Time `json:"state_timestamp"`
	}
}

func New() *VCAP {
	vcap := &VCAP{}

	applicationJson := os.Getenv("VCAP_APPLICATION")
	json.Unmarshal([]byte(applicationJson), vcap.Application)

	// set defaults in case of local development / missing VCAP_APPLICATION
	if vcap.Application.ID == "" {
		vcap.Application.ID = "123-456-789"
	}
	if vcap.Application.Name == "" {
		vcap.Application.Name = "devapp"
	}
	if vcap.Application.InstanceID == "" {
		vcap.Application.InstanceID = "987-654-321"
	}
	if vcap.Application.InstanceIndex == 0 {
		vcap.Application.InstanceIndex = 1
	}
	if vcap.Application.Host == "" {
		vcap.Application.Host = "localhost"
	}
	if vcap.Application.Port == 0 {
		vcap.Application.Port = 4000
	}

	return vcap
}
