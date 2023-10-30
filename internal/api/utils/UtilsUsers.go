package utils

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUsers struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DatabaseConfig struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Network  string `yaml:"network"`
		Address  string `yaml:"address"`
		DBName   string `yaml:"dbName"`
	} `yaml:"database"`
}
