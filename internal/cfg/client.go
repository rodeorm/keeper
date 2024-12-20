package cfg

import (
	"fmt"
	"os"

	"github.com/rodeorm/keeper/internal/logger"
	"gopkg.in/yaml.v2"
)

// GetClientConfigFromFile возвращает конфигурацию из конфигурационного файла клиента
func GetClientConfigFromFile() (*ClientConfig, error) {
	name := "config.yml"
	//fmt.Println(os.Getwd())
	file, err := os.Open(name)

	if err != nil {
		logger.Error("GetClientConfigFromFile", fmt.Sprintf("открытие файла %s", name), err.Error())
		return nil, err
	}
	defer file.Close()

	cfg := &ClientConfig{}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(cfg)
	if err != nil {
		logger.Error("GetClientConfigFromFile", fmt.Sprintf("открытие файла %s", name), err.Error())
		return nil, err
	}

	return cfg, nil
}
