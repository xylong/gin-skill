package config

type Log struct {
	Level      string `json:"level" yaml:"level"`
	Dir        string `json:"dir" yaml:"root_dir"`
	Filename   string `json:"filename" yaml:"filename"`
	Format     string `json:"format" yaml:"format"`
	ShowLine   bool   `son:"show_line" yaml:"show_line"`
	MaxBackups int    `json:"max_backups" yaml:"max_backups"`
	MaxSize    int    `json:"max_size" yaml:"max_size"` // MB
	MaxAge     int    `json:"max_age" yaml:"max_age"`   // day
	Compress   bool   `json:"compress" yaml:"compress"`
}
