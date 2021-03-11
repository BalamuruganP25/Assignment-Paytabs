package server

import (
	logservice "Paytabs/src/Logservice"
	endpoints "Paytabs/src/api/v1/endpoints"
	config "Paytabs/src/config"
	INC "Paytabs/src/messages"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTransaction(t *testing.T) {

	config.ReadEnvironmentVariable()
	logservice.CreateLogMainFolder()
	logservice.CreateLogSubFolder()
	logservice.Applicationdetails()

	router := gin.New()
	router.POST("/v1/transaction", endpoints.Transaction)
	Input := INC.MoneyTransation{
		Fromname:       "bala",
		Usercode:       "F4459ok",
		FromAccountid:  3444444,
		ToAccountid:    3444484,
		Ifsccode:       "XYZBUDI",
		CurrentBalance: 3453.56,
		TransferAmount: 53.56,
		Date:           "2021-03-08T11:38:16.462+00:00",
		Status:         "MoneyTransation",
	}

	RequestJson, err := json.Marshal(Input)
	if err != nil {
		fmt.Println(err)
	}

	req, _ := http.NewRequest("POST", "/v1/transaction", bytes.NewBuffer(RequestJson))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	fmt.Println(resp.Body.String())

}
