package control

// Ucpe 根据SN获取Cpe的数据
type Ucpe struct {
	sn          string
	model       string
	version     string
	updatetime  string
	masterpopip string
	mastercpeip string
	backuppopip string
	backupcpeip string
	remoteport  string
}

// UCpes 多
type UCpes []Ucpe

// Manager 控制器
type Manager interface {
	Shower
	Conner
}

// Shower 控制器
type Shower interface {
	Show()
}

// Conner 控制器
type Conner interface {
	MasterIsEmpty() bool
	BackupIsEmpty() bool
	ConnMaster()
	ConnBackup()
}
