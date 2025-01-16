package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/proto"
)

// Model— это состояние клиентского приложения. Она содержит данные, которые отображаются пользователю.
type Model struct {
	User             core.User // Данные текущего пользователя
	CurrentScreen    string    // Текущий экран для отображения
	Token            string    // Токен
	FilePath         string
	OTPMessageSended bool // Было послано сообщение с OTP
	Authenticated    bool // Авторизован пользователь или нет
	Verified         bool // Подтвержден OTP или нет
	Quitting         bool

	cards []core.Card

	Logo
	Reg
	Auth
	Verify
	Main
	CardCreate
	CardList
	BinaryPick
	BinaryAdd
	BinaryList
	CoupleList
	CoupleCreate
	TextCreate
	TextList

	sc proto.KeeperServiceClient
}

// InitialModel инициализирует модель со значениями по умолчанию
func InitialModel(sc proto.KeeperServiceClient, filePath string) Model {
	var m Model
	m.CurrentScreen = "logo"
	m.User = core.User{}
	m.Token = ""
	m.OTPMessageSended = false
	m.Authenticated = false
	m.Verified = false
	m.Quitting = false
	m.cards = make([]core.Card, 0)
	m.Logo = initLogo()
	m.Auth = initAuth()
	m.Reg = initReg()
	m.Verify = initVerify()
	m.Main = initMain()
	m.FilePath = filePath
	// Данные кредитных карт
	m.CardCreate = initCardCreate()
	m.CardList = initCardList()
	// Данные бинарных файлов
	m.BinaryList = initBinaryList() // Общий список
	m.BinaryPick = initBinaryPick() // Создание (выбор файла)
	m.BinaryAdd = initBinaryAdd()   // Создание (сохранение файла с новым именем и метой)
	// Данные пар логин-пароль
	m.CoupleCreate = initCoupleCreate()
	m.CoupleList = initCoupleList()
	// Данные текстовых сообщений
	m.TextCreate = initTextCreate()
	m.TextList = initTextList()
	m.sc = sc

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}
