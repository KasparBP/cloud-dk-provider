package api

import "context"

type CloudServer struct {
	Identifier string   `json:"identifier"`
	HostName   string   `json:"hostname"`
	Label      string   `json:"label"`
	Cpus       int      `json:"cpus"`
	Memory     int      `json:"memory"`
	Booted     bool     `json:"booted"`
	Disks      []Disk   `json:"disks"`
	NetworkInterfaces []NetworkInterface `json:"networkInterfaces"`
	Template   Template `json:"template"`
	Location   Location `json:"location"`
	Package    Package  `json:"package"`
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
	req, err := svc.c.newRequest("GET", "/v1/cloudservers")
	if err != nil {
		return nil, err
	}
	var cloudServers []CloudServer
	_, err = svc.c.do(ctx, req, &cloudServers)
	return cloudServers, err
}
