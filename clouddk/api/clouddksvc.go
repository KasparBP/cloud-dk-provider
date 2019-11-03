package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
)

type CloudServer struct {
	Identifier          string   `json:"identifier"`
	HostName            string   `json:"hostname"`
	Label               string   `json:"label"`
	InitialRootPassword *string  `json:"initialRootPassword"`
	Cpus                int      `json:"cpus"`
	Memory              int      `json:"memory"`
	Booted              bool     `json:"booted"`
	Disks               []Disk   `json:"disks"`
	NetworkInterfaces   []NetworkInterface `json:"networkInterfaces"`
	Template            Template `json:"template"`
	Location            Location `json:"location"`
	Package             Package  `json:"package"`
}

type Disk struct {
	Identifier string `json:"identifier"`
	Label      string `json:"label"`
	Size       int    `json:"size"`
	Primary    int    `json:"primary"`
}

type NetworkInterface struct {
	Identifier           string `json:"identifier"`
	Label                string `json:"label"`
	RateLimit            int    `json:"rate_limit"`
	DefaultFirewallRule  string `json:"default_firewall_rule"`
	Primary              int    `json:"primary"`
	IpAddresses          []IpAddress    `json:"ipAddresses"`
	FirewallRules        []FirewallRule `json:"firewallRules"`
}

type IpAddress struct {
	Address                    string `json:"address"`
	Network                    string `json:"network"`
	Netmask                    string `json:"netmask"`
	Gateway                    string `json:"gateway"`
	NetworkInterfaceIdentifier string `json:"network_interface_identifier"`
}

type FirewallRule struct {
	Identifier string `json:"identifier"`
	// TODO
}

type Template struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

type Location struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

type Package struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

type ClouddkService struct {
	c *Client
}

// List existing cloudservers
func (svc *ClouddkService) ListCloudServers(ctx context.Context) ([]CloudServer, error) {
	req, err := svc.c.newRequest("GET", "/v1/cloudservers", nil)
	if err != nil {
		return nil, err
	}
	var cloudServers []CloudServer
	_, err = svc.c.do(ctx, req, &cloudServers)
	return cloudServers, err
}

// Create a new cloud server
func (svc *ClouddkService) CreateCloudServer(ctx context.Context, cs *CloudServer) (*CloudServer, error) {
	body, err := json.Marshal(cs)
	if err != nil {
		return nil, err
	}
	req, err := svc.c.newRequest("POST", "/v1/cloudservers", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	var createdCloudServer CloudServer
	_, err = svc.c.do(ctx, req, &createdCloudServer)
	return &createdCloudServer, err
}

// Get cloud server with specified identifier
func (svc *ClouddkService)  GetCloudServer(ctx context.Context, identifier string) (*CloudServer, error) {
	req, err := svc.c.newRequest("GET", "/v1/cloudservers/" + identifier, nil)
	if err != nil {
		return nil, err
	}
	var cs CloudServer
	_, err = svc.c.do(ctx, req, &cs)
	return &cs, err
}

// Delete cloud server with specified identifier
func (svc *ClouddkService) DeleteCloudServer(ctx context.Context, identifier string) error {
	// TODO
	return errors.New("Delete cloud server not implemented yet")
}

// Update cloud server with specified identifier
func (svc *ClouddkService) UpdateCloudServer(ctx context.Context, identifier string, cs *CloudServer) (*CloudServer, error) {
	// TODO
	return nil, errors.New("Update cloud server not implemented yet")
}
