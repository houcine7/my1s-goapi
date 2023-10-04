package tools

import (
	log "github.com/sirupsen/logrus"
)


type LoginDetails struct {
	Username string
	AuthToken string
}

type CoinDetails struct {
	Username string
	Coins int64
}

type DatabaseI interface{
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseI, error){
	var db DatabaseI = &mockDB{}
	err := db.SetupDatabase()
	if err !=nil{
		log.Error(err)
		return nil,err
	}
	return &db,nil
}