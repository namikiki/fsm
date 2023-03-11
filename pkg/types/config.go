package types

type Config struct {
	DataBase DataBase
	Redis    Redis
	Minio    MinIO
	Develop  Develop
}

type Develop struct {
	DevMod bool
}

type DataBase struct {
	DSN string
}

type Redis struct {
	Address  string
	Password string
}

type MinIO struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}
