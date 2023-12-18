package models

// AuthModel – модель данных для запроса регистрации или логина.
type AuthModel struct {
	// Login – логин.
	Login string
	// Password – пароль.
	Password string
}

// AuthToken – токен авторизации.
type AuthToken string

// UserDataList – модель текстовых данных пользователя.
type UserDataList struct {
	// Name – название данных.
	Name string
	// DataType – тип данных.
	DataType int64
	// ID – идентификатор.
	ID int64
	// Version – версия данных.
	Version int64
}

// UserDataModel – модель для получения данных.
type UserDataModel struct {
	// ID – идентификатор.
	ID int64
}

// UserData – модель бинарных данных пользователя.
type UserData struct {
	// Name – название данных.
	Name string
	// DataType – тип данных.
	DataType int64
	// Data – бинарные данные пользователя.
	Data []byte
	// ID – идентификатор.
	ID int64
	// Version – версия данных.
	Version int64
}

// PasswordData – структура для типа данных Пароль.
type PasswordData struct {
	// Site – сайт, пароль от которого пользователь хочет сохранить.
	Site string `json:"site"`
	// Login – логин пользователь.
	Login string `json:"login"`
	// Password – пароль пользователя.
	Password string `json:"password"`
}

// CardData – структура для типа данных Карта.
type CardData struct {
	// Number – номер карты.
	Number string `json:"number"`
	// ExpDate – дата, до которой валидна карта.
	ExpDate string `json:"exp_date"`
	// CardHolder – держатель карты.
	CardHolder string `json:"card_holder"`
}

// FileData – структура для типа данных Файл.
type FileData struct {
	// Path – путь до файла.
	Path string `json:"path"`
	// Data – файл в бинарном представлении.
	Data []byte `json:"data"`
}

// TextData – структура для типа данных Текст.
type TextData struct {
	// Text – текст.
	Text string `json:"text"`
}
