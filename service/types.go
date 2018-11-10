package service

type RegisterIdentificationValueInfo struct {
	DataHash string
	Proof    string
	Info     string
}

type RegisterIdentificationContentInfo struct {
	Path   string
	Values []RegisterIdentificationValueInfo
}

type RegisterIdentificationInfo struct {
	Id       string
	Sign     string
	Contents []RegisterIdentificationContentInfo
}