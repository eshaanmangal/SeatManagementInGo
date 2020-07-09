package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	gorillaMux "github.com/gorilla/mux"
	"net/http"
)

import (
	zoho "github.com/goseatmanagement/Zoho"
	)

var (
	router = gorillaMux.NewRouter()
	)

func main(){
	zohoClient := zoho.ZohoClient{}
	zohoClient.Setup()
	fmt.Println(zohoClient.GetEmployeeDatasetEmailAndEmployeeID(),
							len(zohoClient.GetEmployeeDatasetEmailAndEmployeeID()))

	router.HandleFunc("/zoho/getXebiaLocations",getXebiaLocations).Methods("GET")

}

func getXebiaLocations(writer http.ResponseWriter, request *http.Request) {

}

func getDepartment(context *gin.Context) {

}

func getEmailAndEmployeeIDDataByName(context *gin.Context) {

}

func getCompleteEmployeeDetailsByEmployeeID(context *gin.Context) {

}