Prerequisites 
1.Golang version go1.15 required

2.Before Run program please follow the steps
2.1.install golang  (if your already install golang please skip the step)
    Note  - User - root
   
    Download  Golang package 
		https://golang.org/doc/install
	Unzip the download folder 
   	   tar -xzf golangpackage name 
	Copy go folder 
		cp -rf  go /user/local
	Set golang path 
 		vim ~/.bashrc
	Add the below line in the bashrc 
		export GOROOT=/usr/local/go
		export GOPATH=$HOME/Goprojects
		export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
	Save the file
		source  ~/.bashrc
	Check go version in comment line
		
		go version 
	 Out put 
		go version go1.15 linux/amd64

 2.2 Add environment variable in bashrc file 
 	  vim ~/.bashrc

	 export Paytaps_API_IP=192.168.43.100
 	 export Paytaps_API_PORT=8081
 	 export Paytaps_LOG_FOLDER_NAME=Paytaps_log
 	 export Paytaps_LOG_FILEPATH=/home/

	
	source  ~/.bashrc


2.3 Unit Test 

	2.2 steps should Be Done 

	1.Go to server folder 
	  cd /Paytabs/src/server 
	2.enter below comments
	  go test 

	  OutPut 
	   NOTE - sholud be run with out error 

	   {"fromname":"bala","usercode":"F4459ok","fromaccountid":3444444,"toaccountid":3444484,"ifsccode":"XYZBUDI","currentbalance":3400,"transferamount":53.56,"date":"2021-03-08T11:38:16.462+00:00","status":"Success"}
	   PASS
	   ok  	Paytabs/src/server	0.004s


2.4 Run the program  
    
    Go to project src folder 
    go run main.go

2.5 Test Api  
   	Tool - POSTMAN 

	url - http://ip:port/v1/transaction"
	Method - POST
	Content-type - application/json
    Request Message - 

    {
    "fromname":"bala",
    "usercode":"F4459ok",
    "fromaccountid":3444444,
    "toaccountid":3444484,
    "ifsccode":"XYZBUDI",
    "currentbalance":3321.56,
    "transferamount":10.89,
    "date":"2021-03-08T11:38:16.462+00:00",
    "status":"MoneyTransation"
 
}
