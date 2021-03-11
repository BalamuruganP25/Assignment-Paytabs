package messages

type MoneyTransation struct {
	Fromname       string  `json:"fromname"`
	Usercode       string  `json:"usercode"`
	FromAccountid  uint    `json:"fromaccountid"`
	ToAccountid    uint    `json:"toaccountid"`
	Ifsccode       string  `json:"ifsccode"`
	CurrentBalance float64 `json:"currentbalance"`
	TransferAmount float64 `json:"transferamount"`
	Date           string  `json:"date"`
	Status         string  `json:"status"`
}

type CommonResponse struct {
	Status string `json:"status"`
}

var TransactionHistory []MoneyTransation
