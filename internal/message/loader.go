package message

type Type int

const (
	ExchaneKeys Type = iota
	UserLogin
	DownloadCheat
	Heartbeat
)

type Msg struct {
	Type Type        `json:"type"`
	Hash string      `json:"Hash"`
	Data interface{} `json:"data"`
}
