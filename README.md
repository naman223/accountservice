After check out of the Code<br>
1. export GOPATH=`pwd`<br>
2. cd src/github.com/callistaenterprise/goblog/accountservice<br>
3. go get github.com/boltdb/bolt<br>
4. go get github.com/gorilla/mux<br>
4. go run main.go<br>
5. To test: curl http://localhost:6767/accounts/10000<Br>
On startup bolt db connection is open. Some dummy records created and you can verify the same using curl command.

