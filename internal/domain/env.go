package domain

type Env struct {
	PORT                   int    `envconfig:"PORT"`
	APP_ENV                string `envconfig:"APP_ENV"`
	DB_HOST                string `envconfig:"DB_HOST"`
	DB_PORT                string `envconfig:"DB_PORT"`
	DB_DATABASE            string `envconfig:"DB_DATABASE"`
	DB_USERNAME            string `envconfig:"DB_USERNAME"`
	DB_PASSWORD            string `envconfig:"DB_PASSWORD"`
	AccessTokenExpiryHour  int    `envconfig:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `envconfig:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `envconfig:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `envconfig:"REFRESH_TOKEN_SECRET"`
}
