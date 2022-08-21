package mythic

// VPS -
type VPS struct {
	// Basic properties - always present
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	Product    string `json:"product"`
	Dormant    bool   `json:"dormant"`

	// Detailed properties
	Status     string `json:"status,omitempty"`
	HostServer string `json:"host_server,omitempty"`

	Zone struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"zone,omitempty"`

	CPUMode    string   `json:"cpu_mode,omitempty"`
	NetDevice  string   `json:"net_device,omitempty"`
	DiskBus    string   `json:"disk_bus,omitempty"`
	Price      float32  `json:"price,omitempty"`
	ISOImage   string   `json:"iso_image,omitempty"`
	BootDevice string   `json:"boot_device,omitempty"`
	IPv4       []string `json:"ipv4,omitempty"`
	IPv6       []string `json:"ipv6,omitempty"`

	Specs struct {
		DiskType string `json:"disk_type"`
		DiskSize int    `json:"disk_size"`
		Cores    int    `json:"cores"`
		RAM      int    `json:"ram"`
	} `json:"specs,omitempty"`

	MACs []string `json:"macs,omitempty"`

	AdminConsole struct {
		Username string `json:"username"`
		Hostname string `json:"hostname"`
	} `json:"admin_console,omitempty"`

	SSHProxy struct {
		Hostname string `json:"hostname"`
		Port     int    `json:"port"`
	} `json:"ssh_proxy,omitempty"`

	VNC struct {
		Mode     string `json:"mode"`
		Password string `json:"password"`
		IPv4     string `json:"ipv4"`
		IPv6     string `json:"ipv6"`
		Port     int    `json:"port"`
		Display  int    `json:"display"`
	} `json:"vnc,omitempty"`
}