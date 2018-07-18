After check out of the Code<br>
1. export GOPATH=`pwd`<br>
2. cd src/github.com/callistaenterprise/goblog/accountservice<br>
3. go get github.com/boltdb/bolt<br>
4. go get github.com/gorilla/mux<br>
5. go get github.com/Sirupsen/logrus<br>
6. go run main.go<br>
7. To test:<Br>
    curl -X "POST" http://localhost:6767/accounts/123/naman<br>
    curl -X "POST" -H "Content-Type: application/json" -d '{"id":"223","name":"Naman"}' http://localhost:6767/accounts<br>
    curl http://localhost:6767/accounts/123<Br><Br>
On startup bolt db connection is open. Some dummy records created and you can verify the same using curl command.

