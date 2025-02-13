package configs

import "github.com/spf13/viper"

type Config struct {
	AuthHost     string `mapstructure:"AUTHHOST"`
	AuthPort     string `mapstructure:"AUTHPORT"`
	AuthUser     string `mapstructure:"AUTHUSER"`
	AuthPassword string `mapstructure:"AUTHPASSWORD"`
	AuthName     string `mapstructure:"AUTHDATABASE"`

	EduHost     string `mapstructure:"EDUHOST"`
	EduPort     string `mapstructure:"EDUPORT"`
	EduUser     string `mapstructure:"EDUUSER"`
	EduPassword string `mapstructure:"EDUPASSWORD"`
	EduName     string `mapstructure:"EDUDATABASE"`

	NewsHost     string `mapstructure:"NEWSHOST"`
	NewsPort     string `mapstructure:"NEWSPORT"`
	NewsUser     string `mapstructure:"NEWSUSER"`
	NewsPassword string `mapstructure:"NEWSPASSWORD"`
	NewsName     string `mapstructure:"NEWSNAME"`

	RabbitMQURL string `mapstructure:"TRANSPORTER"`

	CACHER string `mapstructure:"CACHER"`

	KeycloakURL      string `mapstructure:"KEYCLOAK_URL"`
	KeycloakRealm    string `mapstructure:"KEYCLOAK_REALM"`
	KeycloakClientID string `mapstructure:"KEYCLOAK_CLIENT_ID"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
