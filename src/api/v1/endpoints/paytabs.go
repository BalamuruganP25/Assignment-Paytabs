package endpoints

import (
	INC "Paytabs/src/messages"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Transaction(process *gin.Context) {

	var transaction INC.MoneyTransation
	var Common INC.CommonResponse

	INC.Logger.Println("Transaction Process")
	if err := process.BindJSON(&transaction); err != nil {
		INC.Logger.Println("Bind json error in transaction request -->", err)
		fmt.Println("BadRequest")
		process.JSON(http.StatusBadRequest, "BadRequest")
		return
	}

	if transaction.CurrentBalance < transaction.TransferAmount {
		Common.Status = "Insufficient Balance"
		transaction.Status = Common.Status
		INC.TransactionHistory = append(INC.TransactionHistory, transaction)
		process.JSON(http.StatusOK, Common)
		return

	}

	transaction.CurrentBalance = transaction.CurrentBalance - transaction.TransferAmount
	transaction.Status = "Success"

	INC.TransactionHistory = append(INC.TransactionHistory, transaction)
	//fmt.Println("TransactionHistory -->",INC.TransactionHistory)
	process.JSON(http.StatusOK, transaction)

}
