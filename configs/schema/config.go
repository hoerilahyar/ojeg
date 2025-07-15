package schema

type Config struct {
	AppName   string   `yaml:"app_name"`
	DB        DBConfig `yaml:"db"`
	JWTSecret string   `yaml:"jwt_secret"`
}

type DBConfig struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SSLMode  string `yaml:"sslmode"`
}
