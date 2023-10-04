package tools

import (
	"time"
)

type mockDB struct{
}

var mockLoginDetails = map[string]LoginDetails{
	"houcine": LoginDetails{
		AuthToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
		Username: "houcine",
	},
	"mohamed": LoginDetails{
		AuthToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
		Username: "mohamed",
	},
}
var mockCoinDetails = map[string]CoinDetails{
	"houcine": CoinDetails{
		Coins: 100,
		Username: "houcine",
	},
	"mohamed": CoinDetails{
		Coins: 200,
		Username: "mohamed",
	},
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails{
	//simulate a db call
	time.Sleep(1* time.Second)

	var loginDetails = LoginDetails{}

	 loginDetails,ok  :=mockLoginDetails[username]
	if !ok{
		return nil
	}
	return &loginDetails
} 

func (db *mockDB) GetUserCoins(username string) *CoinDetails{
	time.Sleep(1* time.Second)

	var coinDetails = CoinDetails{}

	coinDetails,ok  :=mockCoinDetails[username]
	if !ok{
		return nil
	}
	return &coinDetails
}

func (db *mockDB) SetupDatabase() error{
	return nil
}