package route

// Route
type Route struct {
	InitURL   string   `yaml:"initurl"`
	Tokenurl  string   `yaml:"tokenurl"`
	Operation string   `yaml:"operation"`
	Modes     []string `yaml:"modes"`

	Valor     valor     `yaml:"valor"`
	Nexus     nexus     `yaml:"nexus"`
	Watsons   watsons   `yaml:"watsons"`
	Tassadar  tassadar  `yaml:"tassadar"`
	WatsonsHa watsonsha `yaml:"watsonsha"`
}

// Domain
type Domain struct {
	Cpe string `yaml:"cpe"`
	Pop string `yaml:"pop"`
	Dve string `yaml:"dve"`
}

type valor struct{
	Cpe string `yaml:"cpe"`
	Pop string `yaml:"pop"`
	Dve string `yaml:"dve"`
	Pse int    `yaml:"pse"`
}

type nexus Domain

type watsons Domain

type tassadar Domain

type watsonsha Domain