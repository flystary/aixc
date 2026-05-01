package pop

var (
	entry Entry
	pop   Pop
)

type PopBase struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Operators string `json:"operators"`
	NeedSync  bool   `json:"needSync"`
	DeviceNum int    `json:"deviceNum"`
}

type Pop struct {
	PopBase
	PopIP string `json:"popIp"`
}

type Entry struct {
	PopBase
	EntryIP   string `json:"entryIp"`
	EntryType string `json:"entryType"`
}

type SPop struct {
	PopBase
	PopIP   string `json:"popIp"`
	PopType string `json:"popType"`
	BgpAs   int    `json:"bgpAs"`
}

type PopInfo interface {
	GetID() int
	GetName() string
	GetIP() string
}

func (b PopBase) GetID() int      { return b.ID }
func (b PopBase) GetName() string { return b.Name }

func (p Pop) GetIP() string { return p.PopIP }

func (e Entry) GetIP() string { return e.EntryIP }

func (s SPop) GetIP() string { return s.PopIP }
