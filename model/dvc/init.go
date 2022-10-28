package dvc

// Customer customer
type Customer struct {
	ID 		int `json:"id"`
	Name 	string `json:"name"`
}

type Ndvc struct {
	ID int `json:"id"`
	Manufacturer string `json:"manufacturer"`
	Model string `json:"model"`
	Sn string `json:"sn"`
	WanCount int `json:"wanCount"`
	HardwareVersion string `json:"hardwareVersion"`
	SoftwareVersion string `json:"softwareVersion"`
	Status string `json:"status"`
	ServerAddr string `json:"serverAddr"`
	ServerPort int `json:"serverPort"`
	Wan1IP string `json:"wan1IP"`
	Wan1GW string `json:"wan1GW"`
	Customer Customer `json:"customer"`
	SupportMobile bool `json:"supportMobile"`
	CustomerName string `json:"customerName"`
	ActiveStatus bool `json:"activeStatus"`
	SupportRemote bool `json:"supportRemote"`
}

type Vdvc struct {
	ID int `json:"id"`
	Model string `json:"model"`
	Sn string `json:"sn"`
	WanCount int `json:"wanCount"`
	SoftwareVersion string `json:"softwareVersion"`
	Status string `json:"status"`
	ServerAddr string `json:"serverAddr"`
	ServerPort int `json:"serverPort"`
	Wan1IP string `json:"wan1IP"`
	Wan1GW string `json:"wan1GW"`
	Customer Customer `json:"customer"`
	Support4G bool `json:"support4G"`
	CustomerName string `json:"customerName"`
	ActiveStatus bool `json:"activeStatus"`
	EnableRemote bool `json:"enableRemote"`
}

type Wdvc struct {
	ID int `json:"id"`
	Manufacturer string `json:"manufacturer"`
	Model string `json:"model"`
	Sn string `json:"sn"`
	WanCount int `json:"wanCount"`
	HardwareVersion string `json:"hardwareVersion"`
	SoftwareVersion string `json:"softwareVersion"`
	Status string `json:"status"`
	ServerAddr string `json:"serverAddr"`
	ServerPort int `json:"serverPort"`
	Wan1IP string `json:"wan1IP"`
	Wan1GW string `json:"wan1GW"`
	SupportMobile bool `json:"supportMobile"`
	SupportRemote bool `json:"supportRemote"`
	Level string `json:"level"`
}
