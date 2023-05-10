package configs

import "os"

type EnvType struct {
	MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_TCP_PORT, MYSQL_DATABASE string
}

func GetEnv() EnvType {
	env := EnvType{
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_TCP_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	}

	return env
}
