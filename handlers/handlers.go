package handlers

import (
	"encoding/json"
	"fmt"
	"go-grpc-example/connector"
	"go-grpc-example/proto"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSingleUserHandler - handles requests for a single user.
func GetSingleUserHandler(ctx *gin.Context) {

	// connection with server
	conn := connector.Connect()

	// client is created
	client := proto.NewUserManagementServiceClient(conn)

	// proto.RequestUser type variable is created
	var userData proto.RequestUser

	// request body is read as bytes
	requestBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Body",
		})
		return
	}

	// if request maped to proto struct
	if err := json.Unmarshal(requestBody, &userData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Body",
		})
	}

	// the corresponding server is called to get the desired response
	if response, err := client.GetSingleUser(ctx, &userData); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": response,
		})

	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		fmt.Println("ERROR : ", err.Error())

	}
}

// GetMultipleUserHandler - handles requests for a multiple user.
func GetMultipleUserHandler(ctx *gin.Context) {

	// connection with server
	conn := connector.Connect()

	// client is created
	client := proto.NewUserManagementServiceClient(conn)

	// proto.RequestUsers type variable is created
	var userData proto.RequestUsers

	// request body is read as bytes
	requestBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Body",
		})
		return
	}

	// if request maped to proto struct
	if err := json.Unmarshal(requestBody, &userData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Body",
		})
	}

	// the corresponding server is called to get the desired response
	if response, err := client.GetMultipleUsers(ctx, &userData); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": response,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})

	}

}
