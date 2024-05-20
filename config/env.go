package config

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Database struct {
		User        string `toml:"user"`
		Passwd      string `toml:"passwd"`
		Cluster     string `toml:"cluster"`
		ClusterName string `toml:"clustername"`
		DB          string `toml:"db"`
	}
}

var Envs = initConfig()

func initConfig() Config {
	file, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	if err := toml.Unmarshal(file, &config); err != nil {
		log.Fatal(err)
	}
	return config
}

func (c Config) FormatURI() string {
	return fmt.Sprintf("mongodb+srv://%s:%s@%s.qmdogpy.mongodb.net/?retryWrites=true&w=majority&appName=%s",
		c.Database.User,
		c.Database.Passwd,
		c.Database.Cluster,
		c.Database.ClusterName)
}
