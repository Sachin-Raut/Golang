1. In terminal, navigate to "server" and execute

$ go run main.go

2. Open another terminal window (don't close the first terminal) & navigate to "client"

$ go run main.go --type 1 --playerId 2
(check the client terminal)


$ go run main.go --type 2 --playerId 3
(check the client terminal)


$ go run main.go --type 3 
(check the client terminal & server terminal)

/* If you don't provide playerId, then by default 2 will be provided as playerId */