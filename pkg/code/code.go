package code

const (
	QuerySuccess uint32 = iota
	QueryFailed
)

var messageForCode = map[uint32]string{
	QuerySuccess: "查询成功",
	QueryFailed:  "查询失败",
}

func GetMessage(code uint32) string {
	return messageForCode[code]
}