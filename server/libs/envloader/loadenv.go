package envloader

import "os"

type EnvType struct {
	USER, PASSWORD, HOST, PORT, DATABASE string
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
