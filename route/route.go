package route

// Route
type Route struct {
	InitURL   string   `yaml:"initurl"`
	Tokenurl  string   `yaml:"tokenurl"`
	Operation string   `yaml:"operation"`
	Modes     []string `yaml:"modes"`

	Valor     valor     `yaml:"valor"`
	Yifeng    valor     `yaml:"yifeng"`
	Watsons   watsons   `yaml:"watsons"`
	Tassadar  tassadar  `yaml:"tassadar"`
	WatsonsHa watsonsha `yaml:"watsonsha"`
}

// Classify
type Classify struct {
	Cpe string `yaml:"cpe"`
	Pop string `yaml:"pop"`
	Dve Device `yaml:"dve"`
}

type tassadar Classify

type valor struct{
	Cpe string `yaml:"cpe"`
	Pop string `yaml:"pop"`
	Dve Device `yaml:"dve"`
	Pse int    `yaml:"pse"`
}

type watsons struct{
	Cpe string `yaml:"cpe"`
	Pop string `yaml:"pop"`
	Dve Device `yaml:"dve"`
	Pse int    `yaml:"pse"`
}

type watsonsha struct{
	Cpe string `yaml:"cpe"`
	Pop string `yaml:"pop"`
	Dve Device `yaml:"dve"`
	Pse int    `yaml:"pse"`
}

type Device struct {
	Pool   string `yaml:"pool"`
	Remote string `yaml:"remote"`
}
