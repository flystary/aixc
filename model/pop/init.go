package pop

var (
     entry  Entry
     pop    Pop 
)

type Pop struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	PopIP      string        `json:"popIp"`
	Operators  string        `json:"operators"`
	NeedSync   bool          `json:"needSync"`
	DeviceNum  int           `json:"deviceNum"`
}

type Entry struct {
	ID 			int `json:"id"`
	Name 		string `json:"name"`
	EntryIP 	string `json:"entryIp"`
	Operators 	string `json:"operators"`
	EntryType 	string `json:"entryType"`
	NeedSync 	bool `json:"needSync"`
	DeviceNum 	int `json:"deviceNum"`
}

type SPop struct {
	ID 			int `json:"id"`
	Name 		string `json:"name"`
	EntryIP 	string `json:"entryIp"`
	Operators 	string `json:"operators"`
	PopType 	string `json:"popType"`
	NeedSync 	bool `json:"needSync"`
	DeviceNum 	int `json:"deviceNum"`
	BgpAs 		int `json:"bgpAs"`
}
