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

type DveInfo interface {
	GetSn() string
	GetVersion() string
	GetEnterprise() string
	IsOnline() bool
	IsRemoteEnabled() bool
}

func (b DeviceBase) GetSn() string      { return b.Sn }
func (b DeviceBase) GetVersion() string { return b.SoftwareVersion }
func (b DeviceBase) IsOnline() bool     { return b.Status == "102" }

func (n Nde) GetEnterprise() string { return n.CustomerName }
func (n Nde) IsRemoteEnabled() bool { return n.SupportRemote == true }

func (v Vde) GetEnterprise() string { return v.CustomerName }
func (v Vde) IsRemoteEnabled() bool { return v.EnableRemote == true }

func (w Wde) GetEnterprise() string { return "watsons" }
func (w Wde) IsRemoteEnabled() bool { return w.SupportRemote == true }
