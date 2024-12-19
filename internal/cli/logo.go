package cli

type LogoScreen struct {
	Choice        int  // Текущий выбор
	Chosen        bool // Cделан выбор или нет
	Quitting      bool // Выходим из приложения или нет
	Authenticated bool // Авторизован пользователь или нет
	Verified      bool // Подтвержден OTP или нет
}

func initLogoScreen() LogoScreen {
	return LogoScreen{}
}
