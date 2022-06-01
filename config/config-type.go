package config

type (
	ENV struct {
		Port     string
		Database Database
		Jwt      Jwt
	}

	Database struct {
		URI  string
		Name string
	}

	Jwt struct {
		SecretKey string
		TokenLife string
	}
)
