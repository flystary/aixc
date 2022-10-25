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
	Dvc string `yaml:"dvc"`
}

type valor Domain

type nexus Domain

type watsons Domain

type tassadar Domain

type watsonsha Domain