package control

// Manager 管理
type Manager interface {
	Shower
	// Conner
}

// Shower 展示
type Shower interface {
	Show()
}

// Conner 连接
type Conner interface {
	MasterIsEmpty() bool
	BackupIsEmpty() bool
	ConnMaster()
	ConnBackup()
}