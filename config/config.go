package config

type Configuration struct {
	App      App      `json:"app" yaml:"app"`
	Log      Log      `json:"log" yaml:"log"`
	Database Database `json:"database" yaml:"database"`
}

type App struct {
	Env     string `json:"env" yaml:"env"`
	Port    string `json:"port" yaml:"port"`
	AppName string ` json:"app_name" yaml:"app_name"`
	AppUrl  string `json:"app_url" yaml:"app_url"`
}
