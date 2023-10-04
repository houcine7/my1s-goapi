package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/houcine7/my1s-goapi/api"
	"github.com/houcine7/my1s-goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)


var UnAuthorizedUser = errors.New("Invalid username or token")


func Authorization(next http.Handler)  http.Handler{

	return http.HandlerFunc(func(resWriter http.ResponseWriter,req *http.Request){

		var username = req.URL.Query().Get("username")
		var token = req.Header.Get("Authorization")
		var err error
		fmt.Println("token:",token)
		if username =="" || token ==""{
			 fmt.Println("username or token is empty")
			 log.Error(UnAuthorizedUser)
			 api.RequestErrorHandler(resWriter,UnAuthorizedUser)
			 return
		}

		var database *tools.DatabaseI
		database,err = tools.NewDatabase()
		if err!=nil {
			api.InternalErrorHandler(resWriter,err)
			return
		}

		var loginDetails *tools.LoginDetails 
		loginDetails = (*database).GetUserLoginDetails(username)
		fmt.Println("loginDetails:",loginDetails)
		if loginDetails == nil || ((*loginDetails).AuthToken != token){
			fmt.Println("heeere")
			log.Error(UnAuthorizedUser)
			api.RequestErrorHandler(resWriter,UnAuthorizedUser)
			return
		}
		next.ServeHTTP(resWriter,req)
})
}