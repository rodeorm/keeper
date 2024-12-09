/*
		Package core отражает предметную область.
		Для минимизации размера padding bytes, все fields определены от highest allocation to lowest allocation.
	    Это положит любые обязательные padding bytes на "дно" структур и уменьшит общий размер обязательных padding bytes
*/
package core

//ServerConfig - конфигурация сервера
type ServerConfig struct {
	RunAddress              string `yaml:"RUN_ADDRESS"`                       //Адрес для запуска сервера
	DBProd                  string `yaml:"DB_PRODUCTIVE"`                     //Адрес продуктового сервера БД
	SSLPath                 string `yaml:"SSL_SERTIFICATE_RELATIVE_PATH"`     //Путь к сертификату SSL
	SSLKey                  string `yaml:"SSL_SERTIFICATE_KEY_RELATIVE_PATH"` //Путь к ключу SSL
	DBTest                  string `yaml:"DB_TEST"`                           //Адрес тестового сервера БД
	SMTPServer              string `yaml:"SMTP_SERVER"`                       //Адрес сервера электронной почты
	SMTPLogin               string `yaml:"SMTP_LOGIN"`                        //Логин сервера электронной почты
	SMTPPass                string `yaml:"SMTP_PASSWORD"`                     //Пароль сервера электронной почты
	MaleTemplate            string `yaml:"MALE_TEMPLATE"`                     //Шаблон для сообщений электронной почты
	PasswordKey             string `yaml:"PASSWORD_KEY"`                      //Ключ для шифрования JWT токена
	OneTimePasswordLiveTime int    `yaml:"OTP_LIVE_TIME"`                     //Время жизни пароля для двухфакторной авторизации
	TokenLiveTime           int    `yaml:"TOKEN_LIVE_TIME"`                   //Время жизни авторизационного токена
	MessageBatchSize        int    `yaml:"MESSAGE_BATCH_SIZE"`                //Время жизни авторизационного токена
}

//ClientConfig - конфигурация клиента
type ClientConfig struct {
	RunAddress    string `yaml:"RUN_ADDRESS"`    //Адрес для запуска клиента
	ServerAddress string `yaml:"SERVER_ADDRESS"` //Адрес сервера, с которым должно быть установлено соединение
}
