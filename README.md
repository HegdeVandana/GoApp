# GoApp


$ cd $GOPATH/go                                                                 

$ cd /go/src/github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete  


$ swagger generate server -A todo-list -f ./swagger.yml                                    

$ go install ./cmd/todo-list-server/                                                        


$ todo-list-server 


$ curl -i url

REPONSE CAN BE OBSERVED

OPTIONAL(adding content to databases)


In the same directory

$ sqlite3

$ .open data.db

$ INSERT INTO profiles (email,id,name)

  VALUES( , , );                        
  
$ .quit
