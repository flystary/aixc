package opt


type Opt struct {
	Total 			int `json:"total"`
	Data 			[]Data `json:"data"`
}

type Data struct {
	Sn 				string `json:"sn"`
	ServerPort  	string `json:"serverPort"`
	Model 			string `json:"model"`
	Cid 			string `json:"cid"`
	Imei 			string `json:"imei"`
	Alias 			string `json:"alias"`
	CustomerName 	string `json:"customerName"`
	Platform 		string `json:"platform"`
	ProtalAccount 	string `json:"protalAccount"`
	PortalPassword 	string `json:"portalPassword"`
	AgentVersion 	string `json:"agentVersion"`
}