package cfg

//ServerConfig - конфигурация сервера
type ServerConfig struct {
	RunAddress              string `yaml:"RUN_ADDRESS"`                       //Адрес для запуска сервера
	DBProd                  string `yaml:"DB_PRODUCTIVE"`                     //Адрес продуктового сервера БД
	SSLPath                 string `yaml:"SSL_SERTIFICATE_RELATIVE_PATH"`     //Путь к сертификату SSL
	SSLKey                  string `yaml:"SSL_SERTIFICATE_KEY_RELATIVE_PATH"` //Путь к ключу SSL
	DBTest                  string `yaml:"DB_TEST"`                           //Адрес тестового сервера БД
	SMTPServer              string `yaml:"SMTP_SERVER"`                       //Адрес сервера электронной почты
	SMTPPort                int    `yaml:"SMTP_PORT"`                         //Порт сервера электронной почты
	SMTPLogin               string `yaml:"SMTP_LOGIN"`                        //Логин сервера электронной почты
	SMTPPass                string `yaml:"SMTP_PASSWORD"`                     //Пароль сервера электронной почты
	MaleTemplate            string `yaml:"MALE_TEMPLATE"`                     //Шаблон для сообщений электронной почты
	PasswordKey             string `yaml:"PASSWORD_KEY"`                      //Ключ для шифрования JWT токена
	From                    string `yaml:"FROM"`                              //От кого отправлять сообщения
	FileName                string `yaml:"FILE_NAME"`                         //Имя файла вложения
	CryptKey                string `yaml:"CRYPT_KEY"`                         //Ключ для шифрования данных
	OneTimePasswordLiveTime int    `yaml:"OTP_LIVE_TIME"`                     //Время жизни пароля для двухфакторной авторизации
	TokenLiveTime           int    `yaml:"TOKEN_LIVE_TIME"`                   //Время жизни авторизационного токена (в часах)
	MessageSendPeriod       int    `yaml:"MESSAGE_SEND_PERIOD"`               //Периодичность отправки сообщений (В секундах)
	QueueFillPeriod         int    `yaml:"QUEUE_FILL_PERIOD"`                 //Периодичность наполнения очереди на отправку (В секундах)
	SenderQuantity          int    `yaml:"SENDER_QUANTITY"`                   //Количество отправителей

}

//ClientConfig - конфигурация клиента
type ClientConfig struct {
	RunAddress    string `yaml:"RUN_ADDRESS"`    //Адрес для запуска клиента
	ServerAddress string `yaml:"SERVER_ADDRESS"` //Адрес сервера, с которым должно быть установлено соединение
}
