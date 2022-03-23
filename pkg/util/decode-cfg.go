package util

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
	log "github.com/inconshreveable/log15"
)

/*
	---workdir/
		| -- bin/
		|     |-- chat(I am here)
		|
		| -- etc/
			  |-- config.toml
			  |-- config.json
*/
func findFile() (string, error) {
	var configPath = ""
	log.Info("runtime:", "os", runtime.GOOS)
	if runtime.GOOS == `windows` {
		configPath = "etc/config.toml"
	} else {
		pwd, err := filepath.Abs(filepath.Dir(filepath.Dir(os.Args[0])))
		if err != nil {
			log.Info("get project pwd err", "err", err)
			return "", err
		}
		err = os.Chdir(pwd)
		if err != nil {
			log.Info("get project is empty", "err", err)
			return configPath, err
		}
		d, _ := os.Getwd()
		log.Info("project info:", "dir", d)
		configPath = d + "/etc/config.toml"
	}
	return configPath, nil
}

func Decode(c interface{}) error {
	configPath, err := findFile()
	if err != nil {
		return err
	}
	log.Info("path:", configPath)
	if _, err := toml.DecodeFile(configPath, c); err != nil {
		return err
	}
	log.Info("config info", "cfg", c)
	return nil
}
