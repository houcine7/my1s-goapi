package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/houcine7/my1s-goapi/api"
	"github.com/houcine7/my1s-goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

// get the coin balance of a user
func GetCoinBalance(writer http.ResponseWriter, req *http.Request){
	var params =api.CoinsBalanceParams{}
	var decoder *schema.Decoder= schema.NewDecoder()
	var err error 

	err = decoder.Decode(&params, req.URL.Query())
	if err != nil{
		log.Error(err)
		api.RequestErrorHandler(writer,err)
		return
	}

	var db *tools.DatabaseI
	db,err = tools.NewDatabase()
	if err != nil{
		api.InternalErrorHandler(writer,err)
		return
	}	
	
	var coinsDetails *tools.CoinDetails
	coinsDetails = (*db).GetUserCoins(params.Username)
	if coinsDetails == nil{
		log.Error(err)
		api.InternalErrorHandler(writer,err)
		return
	}

	response := api.CoinsBalanceResponse{
		StatusCode: http.StatusOK,
		Balance: (*coinsDetails).Coins,
	}
	writer.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(writer).Encode(response)
	if(err !=nil){
		log.Error(err)
		api.InternalErrorHandler(writer,err)
		return
	}

}