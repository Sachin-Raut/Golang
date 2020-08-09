How to create/add modules
(dependency management)


1. create a repo (for eg. "bookstore-users-api") on github
2. copy its https path (https://github.com/Sachin-Raut/bookstore-users-api.git)
3. In terminal, navigate to Desktop. Clone the repo.
4. $ git clone https://github.com/Sachin-Raut/bookstore-users-api.git
5. In terminal, navigate to "bookstore-users-api"
6. $ go mod init github.com/Sachin-Raut/bookstore-users-api 
7. $ go test ./...