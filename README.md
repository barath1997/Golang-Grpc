# Golang-Grpc

This is a goalng-grpc example app is used to get user information.
The app is built on client,server grpc model.Instead of a client a gateway can also be used. 

To learn more about GRPC : https://grpc.io/docs/languages/go/quickstart/

The user data is mock data , generated when the service is started.

The maximum number of user data present is 5 here.

Steps to build and run :

   0. cd Golang-Grpc
   1. go mod init "module name"
   2. go build ./...
   3. go run server.go   (to start server)
   4. go run client/client.go  (to start client)

APIs avaialble : 

    1. GetSingleUser : 

       EndPoint : http://localhost:8080/user-management/get-user

       RequestBody : `{
                     "user_id": 1
                      }`

       ResponseBody : `{
                   "result": {
                   "id": 1,
                   "fname": "saravanan",
                   "city": "chennai",
                   "phone": 986537287,
                   "height": 5.1
                        }
                      }`

    2. GetMultipleUsers : 
    
        EndPoint : http://localhost:8080/user-management/get-users

        RequestBody : `{
                     "user_id": [
                          1
                         ]
                       }`

        ResponseBody : `{
                     "result": {
                      "Users": [
                        {
                         "id": 1,
                         "fname": "saravanan",
                         "city": "chennai",
                         "phone": 986537287,
                         "height": 5.1
                        }
                          ]
                        }
                      }`


Unit testing is done for both the APIs and the test cases in the test file are for valid/true test cases.

Failure test cases :

    1. GetSingleUser : 

         userIds := []int64{0, 9, 8, -3} 

    2. GetMultipleUsers :
         
         userIds := [][]int64{{0, -1, 97565}, {7, 0, -466565}, {35, -44, 89}}

Dockerization :
   
   The application server is dockerized.

   Steps to create image and run image : 
     
     1. docker build -t "tag-name" .
     2. docker run -it -p 4040:4040 "image-name"  (starts the server)
    
     Once the server is started use : go run client/client.go (starts client)








