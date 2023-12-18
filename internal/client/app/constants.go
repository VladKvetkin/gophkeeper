package client

import "errors"

const (
	passwordData = iota + 1
	cardData
	fileData
	textData
)

const (
	getUserDataList = iota + 1
	getUserData
	saveUserData
	editUserData
)

var (
	ErrNoData = errors.New(`user has no data`)
)
