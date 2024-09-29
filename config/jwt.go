package config

type JWT struct {
	Secret string `json:"secret" mapstructure:"secret" yaml:"secret"`
	Expire int64  `json:"expire" mapstructure:"expire" yaml:"expire"`
}
