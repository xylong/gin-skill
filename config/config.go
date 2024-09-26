package config

type Configuration struct {
	App      App      `json:"app" json:"app" yaml:"app"`
	Database Database `json:"database" json:"database" yaml:"database"`
}

type App struct {
	Env     string `json:"env" json:"env" yaml:"env"`
	Port    string `json:"port" json:"port" yaml:"port"`
	AppName string `json:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl  string `json:"app_url" json:"app_url" yaml:"app_url"`
}

// Database db配置
type Database struct {
	Driver              string `json:"driver" json:"driver" yaml:"driver"`
	Host                string `json:"host" json:"host" yaml:"host"`
	Port                int    `json:"port" json:"port" yaml:"port"`
	Database            string `json:"database" json:"database" yaml:"database"`
	UserName            string `json:"username" json:"username" yaml:"username"`
	Password            string `json:"password" json:"password" yaml:"password"`
	Charset             string `json:"charset" json:"charset" yaml:"charset"`
	MaxIdleConns        int    `json:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns        int    `json:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	LogMode             string `json:"log_mode" json:"log_mode" yaml:"log_mode"`
	EnableFileLogWriter bool   `json:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"`
	LogFilename         string `json:"log_filename" json:"log_filename" yaml:"log_filename"`
}
