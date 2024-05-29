package interfaces

// avoid import cycle
type RpcUtil interface {
	Get(name string, path string, querys map[string]string, result any) error
	GetWithPathVariable(name string, path string, pathVar string, result any) error
	Post(name string, path string, body any, result any) error
}
