package dve

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DeviceBase struct {
	ID              int    `json:"id"`
	Model           string `json:"model"`
	Sn              string `json:"sn"`
	WanCount        int    `json:"wanCount"`
	SoftwareVersion string `json:"softwareVersion"`
	Status          string `json:"status"`
	ServerAddr      string `json:"serverAddr"`
	ServerPort      int    `json:"serverPort"`
	Wan1IP          string `json:"wan1IP"`
	Wan1GW          string `json:"wan1GW"`
}

type Nde struct {
	DeviceBase
	Manufacturer    string   `json:"manufacturer"`
	HardwareVersion string   `json:"hardwareVersion"`
	Customer        Customer `json:"customer"`
	SupportMobile   bool     `json:"supportMobile"`
	CustomerName    string   `json:"customerName"`
	ActiveStatus    bool     `json:"activeStatus"`
	SupportRemote   bool     `json:"supportRemote"`
}

type Vde struct {
	DeviceBase
	Customer     Customer `json:"customer"`
	Support4G    bool     `json:"support4G"`
	CustomerName string   `json:"customerName"`
	ActiveStatus bool     `json:"activeStatus"`
	EnableRemote bool     `json:"enableRemote"`
}

type Wde struct {
	DeviceBase
	Manufacturer    string `json:"manufacturer"`
	HardwareVersion string `json:"hardwareVersion"`
	SupportMobile   bool   `json:"supportMobile"`
	SupportRemote   bool   `json:"supportRemote"`
	Level           string `json:"level"`
}

type DveGet interface {
	GetSn() string
	GetVersion() string
	IsOnline() bool
	GetEnterprise() string
}

func (n Nde) GetSn() string         { return n.Sn }
func (n Nde) GetVersion() string    { return n.SoftwareVersion }
func (n Nde) GetEnterprise() string { return n.CustomerName }
func (n Nde) IsOnline() bool        { return n.Status == "102" }

func (v Vde) GetSn() string         { return v.Sn }
func (v Vde) GetVersion() string    { return v.SoftwareVersion }
func (v Vde) GetEnterprise() string { return v.CustomerName }
func (v Vde) IsOnline() bool        { return v.Status == "102" }

func (w Wde) GetSn() string         { return w.Sn }
func (w Wde) GetVersion() string    { return w.SoftwareVersion }
func (w Wde) GetEnterprise() string { return "watsons" }
func (w Wde) IsOnline() bool        { return w.Status == "102" }
