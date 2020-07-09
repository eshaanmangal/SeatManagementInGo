package zohoclient

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type EmployeeData struct {
	ModifiedTime    string `json:"modifiedTime"`
	BaseLocation    string `json:"Base Location"`
	Designation     string `json:"Designation"`
	EmailID         string `json:"Xebia Email ID"`
	CoE             string `json:"COE Type"`
	FullName        string `json:"Full Name"`
	FirstName       string `json:"First Name"`
	OwnerID         string `json:"ownerID"`
	ApprovalStatus  string `json:"ApprovalStatus"`
	RecordID        string `json:"recordId"`
	Department      string `json:"Department"`
	SittingLocation string `json:"Sitting Location (System)"`
	OwnerName       string `json:"ownerName"`
	CreatedTime     string `json:"createdTime"`
	EmployeeID      string `json:"EmployeeID"`
	LastName        string `json:"Last Name"`
	MobileNumber    string `json:"Mobile Phone"`
	ReportingTo     string `json:"Reporting To"`
	ClientLocation  string `json:"Client Location"`
}

type SelectedEmployeeDetails struct {
	EmailID         string `json:"Xebia Email ID"`
	EmployeeID      string `json:"EmployeeID"`
}

type ZohoClient struct {
	allEmployeeData []EmployeeData
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func (c *ZohoClient) Setup() {

	var gettingData bool = true
	var sIndex int = 0

	for gettingData {
		req, err := http.NewRequest("GET", "https://people.zoho.com/people/api/forms/P_EmployeeView/records", nil)
		if err != nil {
			gettingData = false
			fmt.Println("Request was unsuccessful", err.Error())
			log.Print(err)
			os.Exit(1)
		}

		sIndexString := strconv.Itoa(sIndex)
		q := req.URL.Query()
		q.Add("authtoken", goDotEnvVariable("TOKEN"))
		q.Add("sIndex", sIndexString)
		req.URL.RawQuery = q.Encode()

		response, err := http.Get(req.URL.String())
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			fmt.Println("Getting Data From Zoho ....")
			var dataCollection []EmployeeData
			body, _ := ioutil.ReadAll(response.Body)
			_ = json.Unmarshal(body, &dataCollection)
			if len(dataCollection) <= 1 {
				fmt.Println("Data Collection Process Completed !!!! ")
				gettingData = false
				break
			}
			c.allEmployeeData = append(c.allEmployeeData, dataCollection...)
			sIndex += 200
		}
	}
}

func (c *ZohoClient)  GetEmployeeDetails() map[string]EmployeeData{
	employeeData := make(map[string]EmployeeData)
	for _,employeeDetails := range c.allEmployeeData{
		employeeData[employeeDetails.EmployeeID] = employeeDetails
	}
	return employeeData
}

func (c *ZohoClient) GetEmployeeDatasetEmailAndEmployeeID() map[string]SelectedEmployeeDetails{
	employeeData := make(map[string]SelectedEmployeeDetails)
	for _,employeeDetails := range c.allEmployeeData{
		inputDetails := SelectedEmployeeDetails{EmailID: employeeDetails.EmailID,
												EmployeeID: employeeDetails.EmployeeID}
		employeeData[employeeDetails.FullName] = inputDetails
	}
	return employeeData
}

func (c *ZohoClient) GetLocations() []string{
	locationsMap := make(map[string]int)
	for _,values := range c.allEmployeeData{
		locationsMap[values.BaseLocation]++
	}
	locationsList := make([]string,0)
	for index,_:= range locationsMap{
		locationsList = append(locationsList,index)
	}
	return locationsList
}

func (c *ZohoClient) GetUsers() map[string][]EmployeeData {

	users := make(map[string][]EmployeeData)

	for _, employee := range c.allEmployeeData {
		var tempData EmployeeData
		if employee.Department == "Human Resources" || employee.Department == "Admin" || employee.Department == "Delivery" {
			tempData.EmailID = employee.EmployeeID
			tempData.EmployeeID = employee.EmployeeID
			tempData.FullName = employee.FullName
			users[employee.Department] = append(users[employee.Department], tempData)
		}
	}
	return users
}

func (c *ZohoClient) GetDepartments() []string{
	departmentsMap := make(map[string]int)
	for _, employee := range c.allEmployeeData {
		departmentsMap[employee.Department]++
	}

	departmentList := make([]string,0)
	for index,_:=range departmentsMap{
		departmentList = append(departmentList,index)
	}
	return departmentList
}