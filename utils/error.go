package utils

type Error struct {
	Level string
	Desc  string
}

type ErrorCollector interface {
	AppendError(desc string)
	AppendWarning(desc string)
	AppendInfo(desc string)
}
