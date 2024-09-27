package config

type Log struct {
	Level      string `json:"level" yaml:"level" mapstructure:"level"`
	Dir        string `json:"dir" yaml:"root_dir" mapstructure:"dir"`
	Filename   string `json:"filename" yaml:"filename" mapstructure:"filename"`
	Format     string `json:"format" yaml:"format" mapstructure:"format"`
	ShowLine   bool   `son:"show_line" yaml:"show_line" mapstructure:"show_line"`
	MaxBackups int    `json:"max_backups" yaml:"max_backups" mapstructure:"max_backups"`
	MaxSize    int    `json:"max_size" yaml:"max_size" mapstructure:"max_size"` // MB
	MaxAge     int    `json:"max_age" yaml:"max_age" mapstructure:"max_age"`    // day
	Compress   bool   `json:"compress" yaml:"compress" mapstructure:"compress"`
}
