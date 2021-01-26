# Go Chuck Norris Fact 
This is a little test of GOLang and Fiber, this project gets a funny fact about Chuck Norris from 
this [API](https://api.chucknorris.io), map it to a struct and make the response available to the user.   

# More about GO
This link [here](https://tour.golang.org/) has the basic information about the GOLang, like variables, functions and so on.

# Setup

To setup the the fiber you need a ```go.mod``` file, to see what the commands bellow do, please read the [article](https://golangbyexample.com/go-mod-tidy/)

First run the command bellow: 

````go
go mod init bitwiseGo
go get github.com/gofiber/fiber/v2
go mod tidy
````