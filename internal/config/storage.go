package config

type Storage struct {
	Driver    string `json:"driver" yaml:"driver"`
	Directory string `json:"directory" yaml:"directory"`
	Bucket    string `json:"bucket" yaml:"bucket"`
	Host      string `json:"host" yaml:"host"`
	AutoHost bool `json:"auto_host" yaml:"auto_host"`
}

const (
	StorageLocal = "local"
)
