# bookstore
An bookstore management system dashboard for monitoring the book lending activity in Zine and Zin

## Get started

You can clone the repository into a working folder in your local machine and then initialize the go package modules and get started

1. cd /folder
2. git clone https://github.com/Ezinnem/bookstore.git
3. cd bookstore 
4. go mod init <yourmodulename.com>
5. go mod tidy
Remember to change the mysql details in pkg/config/app.go file 
6. cd cmd/main
7. go build
8. go run main.go
9. In Postman, go to the server route http:localhost:2023/

