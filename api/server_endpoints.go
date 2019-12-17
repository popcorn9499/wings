package api

import (
	"encoding/json"
	"fmt"
	"github.com/pterodactyl/wings/parser"
)

const (
	ProcessStopCommand    = "command"
	ProcessStopSignal     = "signal"
	ProcessStopNativeStop = "stop"
)

// Defines the process configuration for a given server instance. This sets what the
// daemon is looking for to mark a server as done starting, what to do when stopping,
// and what changes to make to the configuration file for a server.
type ServerConfiguration struct {
	Startup struct {
		Done            string   `json:"done"`
		UserInteraction []string `json:"userInteraction"`
	} `json:"startup"`
	Stop struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"stop"`
	ConfigurationFiles []parser.ConfigurationFile `json:"configs"`
}

// Fetches the server configuration and returns the struct for it.
func (r *PanelRequest) GetServerConfiguration(uuid string) (*ServerConfiguration, *RequestError, error) {
	resp, err := r.Get(fmt.Sprintf("/servers/%s/configuration", uuid))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	r.Response = resp

	if r.HasError() {
		return nil, r.Error(), nil
	}

	res := &ServerConfiguration{}
	b, _ := r.ReadBody()

	if err := json.Unmarshal(b, res); err != nil {
		return nil, nil, err
	}

	return res, nil, nil
}
