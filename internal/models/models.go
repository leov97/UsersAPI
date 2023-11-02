package models

type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"password,min=8"`
}

type LoginUsers struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type DatabaseConfig struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Network  string `yaml:"network"`
		Address  string `yaml:"address"`
		DBName   string `yaml:"dbName"`
	} `yaml:"database"`

	SecreKey struct {
		Key string `yaml:"key"`
	} `yaml:"keytoken"`
}

type Authlogin struct {
	Token string `json:"token" validate:"required"`
}
