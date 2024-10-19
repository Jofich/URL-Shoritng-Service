package config

type Config struct {
	HTTPServer `yaml:"http_server"`
	StorageCfg `yaml:"storage"`
}

type HTTPServer struct {
	Address string `yaml:"address"`
}
type StorageCfg struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Login    string `yaml:"username"`
	Password string `yaml:"password"`
	DB_name  string `yaml:"db_name"`
}
