package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type Bool bool

type CloudServer struct {
	Identifier          string   `json:"identifier,omitempty"`
	HostName            string   `json:"hostname,omitempty"`
	Label               string   `json:"label,omitempty"`
	InitialRootPassword *string  `json:"initialRootPassword,omitempty"`
	Cpus                int      `json:"cpus,omitempty"`
	Memory              int      `json:"memory,omitempty"`
	Booted              Bool     `json:"booted,omitempty"`
	Disks               []Disk   `json:"disks,omitempty"`
	NetworkInterfaces   []NetworkInterface `json:"networkInterfaces,omitempty"`
	Package             Package  `json:"package,omitempty"`
	Template            Template `json:"template,omitempty"`
	Location            Location `json:"location,omitempty"`
}

type Disk struct {
	Identifier string `json:"identifier,omitempty"`
	Label      string `json:"label,omitempty"`
	Size       int    `json:"size,omitempty"`
	Primary    Bool   `json:"primary,omitempty"`
}

type NetworkInterface struct {
	Identifier           string `json:"identifier,omitempty"`
	Label                string `json:"label,omitempty"`
	RateLimit            int    `json:"rate_limit,omitempty"`
	DefaultFirewallRule  string `json:"default_firewall_rule,omitempty"`
	Primary              Bool   `json:"primary,omitempty"`
	IpAddresses          []IpAddress    `json:"ipAddresses,omitempty"`
	FirewallRules        []FirewallRule `json:"firewallRules,omitempty"`
}

type IpAddress struct {
	Address                    string `json:"address,omitempty"`
	Network                    string `json:"network,omitempty"`
	Netmask                    string `json:"netmask,omitempty"`
	Gateway                    string `json:"gateway,omitempty"`
	NetworkInterfaceIdentifier string `json:"network_interface_identifier,omitempty"`
}

type FirewallRule struct {
	Identifier string `json:"identifier"`
	// TODO
}

type Template struct {
	Identifier string `json:"identifier,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Location struct {
	Identifier string `json:"identifier,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Package struct {
	Identifier string `json:"identifier,omitempty"`
	Name       string `json:"name,omitempty"`
}

type ClouddkService struct {
	c *Client
}

func (bit *Bool) UnmarshalJSON(b []byte) error {
	txt := string(b)
	*bit = txt == "1" || txt == "true"
	return nil
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
	req, err := svc.c.newRequest("DELETE", "/v1/cloudservers/"+identifier, nil)
	if err != nil {
		return err
	}
	var output Bool
	_, err = svc.c.do(ctx, req, &output)
	return err
}

// Update cloud server with specified identifier
func (svc *ClouddkService) UpdateCloudServer(ctx context.Context, identifier string, cs *CloudServer) (*CloudServer, error) {
	// TODO
	return nil, errors.New("update cloud server not implemented yet")
}

func (svc *ClouddkService) CreateDisk(ctx context.Context, csid string, disk Disk) (*Disk, error) {
	body, err := json.Marshal(disk)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("/v1/cloudservers/%s/disks", csid)
	req, err := svc.c.newRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	var createdDisk Disk
	_, err = svc.c.do(ctx, req, &createdDisk)
	return &createdDisk, err
}

func (svc *ClouddkService) GetDisk(ctx context.Context, csid string, id string) (*Disk, error) {
	url := fmt.Sprintf("/v1/cloudservers/%s/disks/%s", csid, id)
	req, err := svc.c.newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var disk Disk
	_, err = svc.c.do(ctx, req, &disk)
	return &disk, err
}

