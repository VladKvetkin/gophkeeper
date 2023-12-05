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
	DataType string
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
	DataType string
	// Data – бинарные данные пользователя.
	Data []byte
	// ID – идентификатор.
	ID int64
	// Version – версия данных.
	Version int64
}
