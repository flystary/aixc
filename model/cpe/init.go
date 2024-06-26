package cpe

type Cpe struct {
	ID                      int         `json:"id"`
	Model                   string      `json:"model"`
	Sn                      string      `json:"sn"`
	Alias                   string      `json:"alias"`
	WanCount                string      `json:"wanCount"`
	SoftwareVersion         string      `json:"softwareVersion"`
	Status                  string      `json:"status"`
	ServerAddr              string      `json:"serverAddr"`
	ServerPort              int         `json:"serverPort"`
	Wan1IP                  string      `json:"wan1IP"`
	Wan1GW                  string      `json:"wan1GW"`
	Wan1Netmask             string      `json:"wan1Netmask"`
	Wan2IP                  string      `json:"wan2IP"`
	Wan2GW                  string      `json:"wan2GW"`
	Wan2Netmask             string      `json:"wan2Netmask"`
	PubIP                   string      `json:"pubIP"`
	Customer                Customer    `json:"customer"`
	MasterPopID             int         `json:"masterPopId"`
	BackupPopID             int         `json:"backupPopId"`
	Mobile4GInPopID         int         `json:"mobile4gInPopId"`
	Mobile5GInPopID         int         `json:"mobile5gInPopId"`
	MasterPopFecMode        string      `json:"masterPopFecMode"`
	MasterPopEnableKcp      bool        `json:"masterPopEnableKcp"`
	BackupPopFecMode        string      `json:"backupPopFecMode"`
	BackupPopEnableKcp      bool        `json:"backupPopEnableKcp"`
	Mobile4GInPopFecMode    string      `json:"mobile4gInPopFecMode"`
	Mobile4GInPopEnableKcp  bool        `json:"mobile4gInPopEnableKcp"`
	Mobile5GInPopFecMode    string      `json:"mobile5gInPopFecMode"`
	Mobile5GInPopEnableKcp  bool        `json:"mobile5gInPopEnableKcp"`
	Mobile4GOutPopID        int         `json:"mobile4gOutPopId"`
	Mobile4GOutPopType      string      `json:"mobile4gOutPopType"`
	Mobile4GOutPopFecMode   string      `json:"mobile4gOutPopFecMode"`
	Mobile4GOutPopEnableKcp bool        `json:"mobile4gOutPopEnableKcp"`
	Enable                  bool        `json:"enable"`
	Support4G               bool        `json:"support4G"`
	Support5G               bool        `json:"support5G"`
	CustomerName            string      `json:"customerName"`
	RedteaContainerID       int         `json:"redteaContainerId"`
	ShanjinContainerID      int         `json:"shanjinContainerId"`
	MobileProvider          string      `json:"mobileProvider"`
	AgencyID                int         `json:"agencyId"`
	VpnServerIP1            string      `json:"vpnServerIp1"`
	VpnServerIP2            string      `json:"vpnServerIp2"`
	MasterPopType           string      `json:"masterPopType"`
	BackupPopType           string      `json:"backupPopType"`
	Mobile4GInPopType       string      `json:"mobile4gInPopType"`
	Mobile5GInPopType       string      `json:"mobile5gInPopType"`
	MasterPopIP             string      `json:"masterPopIp"`
	BackupPopIP             string      `json:"backupPopIp"`
	Mobile4GInPopIP         string      `json:"mobile4gInPopIp"`
	Mobile5GInPopIP         string      `json:"mobile5gInPopIp"`
	EntryUpdateTime         string      `json:"entryUpdateTime"`
	MobileRedteaCID         int         `json:"mobileRedteaCID"`
	Mode                    interface{} `json:"mode"`
}

type Spe struct {
	ID                      int         `json:"id"`
	Model                   string      `json:"model"`
	Sn                      string      `json:"sn"`
	Alias                   string      `json:"alias"`
	WanCount                string      `json:"wanCount"`
	SoftwareVersion         string      `json:"softwareVersion"`
	Status                  string      `json:"status"`
	ServerAddr              string      `json:"serverAddr"`
	ServerPort              int         `json:"serverPort"`
	Wan1IP                  string      `json:"wan1IP"`
	Wan1GW                  string      `json:"wan1GW"`
	Wan1Netmask             string      `json:"wan1Netmask"`
	Wan2IP                  string      `json:"wan2IP"`
	Wan2GW                  string      `json:"wan2GW"`
	Wan2Netmask             string      `json:"wan2Netmask"`
	PubIP                   string      `json:"pubIP"`
	Customer                Customer    `json:"customer"`
	MasterPopID             int         `json:"masterPopId"`
	BackupPopID             int         `json:"backupPopId"`
	Mobile4GInPopID         int         `json:"mobile4gInPopId"`
	Mobile5GInPopID         int         `json:"mobile5gInPopId"`
	MasterPopFecMode        string      `json:"masterPopFecMode"`
	MasterPopEnableKcp      bool        `json:"masterPopEnableKcp"`
	BackupPopFecMode        string      `json:"backupPopFecMode"`
	BackupPopEnableKcp      bool        `json:"backupPopEnableKcp"`
	Mobile4GInPopFecMode    string      `json:"mobile4gInPopFecMode"`
	Mobile4GInPopEnableKcp  bool        `json:"mobile4gInPopEnableKcp"`
	Mobile5GInPopFecMode    string      `json:"mobile5gInPopFecMode"`
	Mobile5GInPopEnableKcp  bool        `json:"mobile5gInPopEnableKcp"`
	Mobile4GOutPopID        int         `json:"mobile4gOutPopId"`
	Mobile4GOutPopType      string      `json:"mobile4gOutPopType"`
	Mobile4GOutPopFecMode   string      `json:"mobile4gOutPopFecMode"`
	Mobile4GOutPopEnableKcp bool        `json:"mobile4gOutPopEnableKcp"`
	Enable                  bool        `json:"enable"`
	Support4G               bool        `json:"support4G"`
	Support5G               bool        `json:"support5G"`
	CustomerName            string      `json:"customerName"`
	RedteaContainerID       int         `json:"redteaContainerId"`
	ShanjinContainerID      int         `json:"shanjinContainerId"`
	MobileProvider          string      `json:"mobileProvider"`
	AgencyID                int         `json:"agencyId"`
	VpnServerIP1            string      `json:"vpnServerIp1"`
	VpnServerIP2            string      `json:"vpnServerIp2"`
	MasterPopType           string      `json:"masterPopType"`
	BackupPopType           string      `json:"backupPopType"`
	Mobile4GInPopType       string      `json:"mobile4gInPopType"`
	Mobile5GInPopType       string      `json:"mobile5gInPopType"`
	MasterPopIP             string      `json:"masterPopIp"`
	BackupPopIP             string      `json:"backupPopIp"`
	Mobile4GInPopIP         string      `json:"mobile4gInPopIp"`
	Mobile5GInPopIP         string      `json:"mobile5gInPopIp"`
	PopUpdateTime           string      `json:"popUpdateTime"`
	MobileRedteaCID         int         `json:"mobileRedteaCID"`
	Mode                    interface{} `json:"mode"`
}

type Box struct {
	ID                   int         `json:"id"`
	Model                string      `json:"model"`
	Sn                   string      `json:"sn"`
	Alias                string      `json:"alias"`
	HardwareVersion      float32     `json:"hardwareVersion"`
	SoftwareVersion      string      `json:"softwareVersion"`
	Status               interface{} `json:"status"`
	ServerAddr           string      `json:"serverAddr"`
	ServerPort           int         `json:"serverPort"`
	Wan1IP               string      `json:"wan1IP"`
	Wan1GW               string      `json:"wan1GW"`
	Wan2IP               string      `json:"wan2IP"`
	Wan2GW               string      `json:"wan2GW"`
	PubIP                string      `json:"pubIP"`
	Customer             Customer    `json:"customer"`
	MasterEntryID        int         `json:"masterEntryId"`
	BackupEntryID        int         `json:"backupEntryId"`
	MobileEntryID        int         `json:"mobileEntryId"`
	MasterEntryFecMode   string      `json:"masterEntryFecMode"`
	MasterEntryEnableKcp bool        `json:"masterEntryEnableKcp"`
	BackupEntryFecMode   string      `json:"backupEntryFecMode"`
	BackupEntryEnableKcp bool        `json:"backupEntryEnableKcp"`
	MobileEntryFecMode   string      `json:"mobileEntryFecMode"`
	MobileEntryEnableKcp bool        `json:"mobileEntryEnableKcp"`
	MobileProxy          bool        `json:"mobileProxy"`
	MobilePopID          int         `json:"mobilePopId"`
	MobilePopType        string      `json:"mobilePopType"`
	MobilePopFecMode     string      `json:"mobilePopFecMode"`
	MobilePopEnableKcp   bool        `json:"mobilePopEnableKcp"`
	Enable               bool        `json:"enable"`
	SupportMobile        bool        `json:"supportMobile"`
	CustomerName         string      `json:"customerName"`
	RedteaContainerID    int         `json:"redteaContainerId"`
	ShanjinContainerID   int         `json:"shanjinContainerId"`
	MobileProvider       interface{} `json:"mobileProvider"`
	AgencyID             int         `json:"agencyId"`
	VpnServerEnable      bool        `json:"vpnServerEnable"`
	VpnServerIP          string      `json:"vpnServerIp"`
	Authority            int         `json:"authority"`
	MasterEntryType      string      `json:"masterEntryType"`
	BackupEntryType      string      `json:"backupEntryType"`
	MobileEntryType      string      `json:"mobileEntryType"`
	MasterEntryIP        string      `json:"masterEntryIp"`
	BackupEntryIP        string      `json:"backupEntryIp"`
	MobileEntryIP        string      `json:"mobileEntryIp"`
	EntryUpdateTime      string      `json:"entryUpdateTime"`
	SupportRemote        interface{} `json:"supportRemote"`
	Mode                 interface{} `json:"mode"`
}

type Customer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Info         string `json:"info"`
	Province     string `json:"province"`
	City         string `json:"city"`
	Contact      string `json:"contact"`
	ContactPhone string `json:"contactPhone"`
	Email        string `json:"email"`
	SextantName  string `json:"sextantName"`
}

type Vox struct {
	ID                   int         `json:"id"`
	Model                string      `json:"model"`
	Sn                   string      `json:"sn"`
	Alias                string      `json:"alias"`
	WanCount             float64     `json:"wanCount"`
	HardwareVersion      float32     `json:"hardwareVersion"`
	SoftwareVersion      string      `json:"softwareVersion"`
	Status               interface{} `json:"status"`
	ServerAddr           string      `json:"serverAddr"`
	ServerPort           int         `json:"serverPort"`
	Wan1IP               string      `json:"wan1IP"`
	Wan1GW               string      `json:"wan1GW"`
	Wan2IP               string      `json:"wan2IP"`
	Wan2GW               string      `json:"wan2GW"`
	PubIP                string      `json:"pubIP"`
	SpeedUp              int         `json:"speedUp"`
	SpeedDown            int         `json:"speedDown"`
	MasterEntryID        int         `json:"masterEntryId"`
	BackupEntryID        int         `json:"backupEntryId"`
	MobileEntryID        int         `json:"mobileEntryId"`
	Enable               bool        `json:"enable"`
	SupportMobile        bool        `json:"supportMobile"`
	CustomerName         string      `json:"customerName"`
	RedteaContainerID    int         `json:"redteaContainerId"`
	ShanjinContainerID   int         `json:"shanjinContainerId"`
	AgencyID             int         `json:"agencyId"`
	Ssl1TCPIP            string      `json:"ssl1TcpIp"`
	Ssl1TCPNetmask       string      `json:"ssl1TcpNetmask"`
	Ssl1UDPIP            string      `json:"ssl1UdpIp"`
	Ssl1UDPNetmask       string      `json:"ssl1UdpNetmask"`
	Ssl2IP               string      `json:"ssl2Ip"`
	Ssl2Netmask          string      `json:"ssl2Netmask"`
	Type3IP              string      `json:"type3Ip"`
	Type3Netmask         string      `json:"type3Netmask"`
	ServerIP             string      `json:"serverIp"`
	IpsecIP              string      `json:"ipsecIp"`
	IpsecNetmask         string      `json:"ipsecNetmask"`
	Authority            int         `json:"authority"`
	MobileProxy          bool        `json:"mobileProxy"`
	MobilePopID          int         `json:"mobilePopId"`
	MobilePopType        string      `json:"mobilePopType"`
	MobilePopFecMode     string      `json:"mobilePopFecMode"`
	MobilePopEnableKcp   bool        `json:"mobilePopEnableKcp"`
	MasterEntryType      string      `json:"masterEntryType"`
	BackupEntryType      string      `json:"backupEntryType"`
	MobileEntryType      string      `json:"mobileEntryType"`
	MasterEntryIP        string      `json:"masterEntryIp"`
	BackupEntryIP        string      `json:"backupEntryIp"`
	MobileEntryIP        string      `json:"mobileEntryIp"`
	MasterEntryFecMode   string      `json:"masterEntryFecMode"`
	MasterEntryEnableKcp bool        `json:"masterEntryEnableKcp"`
	BackupEntryFecMode   string      `json:"backupEntryFecMode"`
	BackupEntryEnableKcp bool        `json:"backupEntryEnableKcp"`
	MobileEntryFecMode   string      `json:"mobileEntryFecMode"`
	MobileEntryEnableKcp bool        `json:"mobileEntryEnableKcp"`
	EntryUpdateTime      string      `json:"entryUpdateTime"`
	Level                interface{} `json:"level"`
	SupportRemote        interface{} `json:"supportRemote"`
	MobileRedteaCID      int         `json:"mobileRedteaCID"`
	Mode                 interface{} `json:"mode"`
}

var (
	cpe Cpe
	spe Spe
	box Box
	vox Vox
)

var sns = make([]string, 0)
