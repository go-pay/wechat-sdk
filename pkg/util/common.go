package util

const (
	TimeLayout = "2006-01-02 15:04:05"
	DateLayout = "2006-01-02"
	NULL       = ""
)

type File struct {
	Name    string `json:"name"`
	Content []byte `json:"content"`
}
