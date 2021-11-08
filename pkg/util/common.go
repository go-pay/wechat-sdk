package util

const (
	NULL = ""
)

type File struct {
	Name    string `json:"name"`
	Content []byte `json:"content"`
}
