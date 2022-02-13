package builder

import "github.com/cristian0193/golang-service-template/internal/utils"

type Configuration struct {
	ServerPort int
	LogLevel   string
	Country    string
}

func LoadConfiguration() (*Configuration, error) {
	serverPort, err := utils.GetInt("SERVER_PORT")
	if err != nil {
		return nil, err
	}
	logLevel, err := utils.GetString("LOG_LEVEL")
	if err != nil {
		return nil, err
	}
	country, err := utils.GetString("COUNTRY")
	if err != nil {
		return nil, err
	}

	return &Configuration{
		ServerPort: serverPort,
		LogLevel:   logLevel,
		Country:    country,
	}, nil
}
