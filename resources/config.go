package resources

type LoadConfig struct {
	Config Config `json:"config"`
}

type Config struct {
	Postgres PostgresConfig `json:"postgres"`
}

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}
