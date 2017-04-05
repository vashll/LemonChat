package msg

type Message struct {
	MsgId       string
	DstId       string // destination id
	SrcId       string //source id
	MsgType     int
	Content     string
	CreateTime  int64
	IsRead      bool
	SendSuccess bool
}
