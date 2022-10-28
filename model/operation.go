package model

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

type Operation Opt

func (o Operation) IsSn(sn string) bool {
	var is = false
	for i := 0; i < len(o.Data); i++ {
		if sn == o.Data[i].Sn {
			is = true
			break
		}
		continue
	}
	return is
}

func(o Operation) SnInMode(sn string) string {
	var mode string

	for i := 0; i < len(o.Data); i++ {

		if sn == o.Data[i].Sn {
			mode = o.Data[i].Platform
			break
		}
		continue
	}

	if mode == "nuxes" {
		mode = "nexus"
	}else if mode =="nuxes-watsons" {
		mode = "watsons"
	} else if mode =="watsons-ha" {
		mode = "watsonsha"
	}

	return mode
}
