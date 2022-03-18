package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gin_rest_client/dto"
	"github.com/valyala/fasthttp"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello World @@@@@")

	//sendRequestWithHttpNet()
	//readConfigFromYamlFile()
	//sendRequestWithFastHttp()
	sendRequestWithFastHttp2("0079028748")

	person := dto.Person{
		Firstname:    "saber",
		Lastname:     "Azizi",
		NationalCode: "0079028748",
		Age:          34,
		Email:        "saberazizi66@yahoo.com",
		Mobile:       "09365627895",
	}
	sendRequestWithFastHttp3(person)
}
func sendRequestWithFastHttp() {

	username := "saber66"
	password := "saber@123"
	authorization := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password))))

	fmt.Printf("send Request with fastHttp with authorization %s \n", authorization)
	request := fasthttp.AcquireRequest()
	request.SetRequestURI("https://localhost:9090/services/person/findAll")
	//request.Header.Set("Authorization", "Basic c2FiZXI2NjpzYWJlckAxMjM=")
	request.Header.Set("Authorization", authorization)
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)
	tls := &tls.Config{
		InsecureSkipVerify: true,
	}
	client := fasthttp.Client{
		ReadTimeout:         30 * time.Second,
		TLSConfig:           tls,
		MaxConnDuration:     30 * time.Second,
		MaxConnWaitTimeout:  30 * time.Second,
		MaxIdleConnDuration: 30 * time.Second,
		MaxConnsPerHost:     3000,
	}
	err := client.Do(request, response)
	if err != nil {
		log.Fatal(err)
	}
	statusCode := response.StatusCode()
	bodyBytes := response.Body
	responseBody := string(bodyBytes())
	fmt.Println(fmt.Sprintf("response with statusCode %d with body %s", statusCode, responseBody))

	persons := dto.FindAllPersonResponse{}

	err = json.Unmarshal(bodyBytes(), &persons)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(persons)
}

func sendRequestWithFastHttp2(nationalCode string) {

	username := "saber66"
	password := "saber@123"
	authorization := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password))))

	url := fmt.Sprintf("https://localhost:9090/services/person/find/%s", nationalCode)

	fmt.Printf("send Request with fastHttp with authorization %s \n", authorization)
	request := fasthttp.AcquireRequest()
	request.SetRequestURI(url)
	//request.Header.Set("Authorization", "Basic c2FiZXI2NjpzYWJlckAxMjM=")
	request.Header.Set("Authorization", authorization)
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)
	tls := &tls.Config{
		InsecureSkipVerify: true,
	}
	client := fasthttp.Client{
		ReadTimeout:         30 * time.Second,
		TLSConfig:           tls,
		MaxConnDuration:     30 * time.Second,
		MaxConnWaitTimeout:  30 * time.Second,
		MaxIdleConnDuration: 30 * time.Second,
		MaxConnsPerHost:     3000,
	}
	err := client.Do(request, response)
	if err != nil {
		log.Fatal(err)
	}
	statusCode := response.StatusCode()
	bodyBytes := response.Body
	//responseBody := string(bodyBytes())
	//fmt.Println(fmt.Sprintf("response with statusCode %d with body %s", statusCode, responseBody))

	if statusCode != 200 {
		errorResponseDto := dto.ErrorResponse{}
		err := json.Unmarshal(bodyBytes(), &errorResponseDto)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(errorResponseDto)
	} else {
		person := dto.Person{}
		err := json.Unmarshal(bodyBytes(), &person)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(person)

	}

}

func sendRequestWithFastHttp3(person dto.Person) {

	username := "saber66"
	password := "saber@123"
	authorization := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password))))

	url := "https://localhost:9090/services/person/add"

	fmt.Printf("send Request with fastHttp with authorization %s \n", authorization)
	request := fasthttp.AcquireRequest()
	request.SetRequestURI(url)
	request.SetBodyString(person.String())
	//request.Header.Set("Authorization", "Basic c2FiZXI2NjpzYWJlckAxMjM=")
	request.Header.Set("Authorization", authorization)
	request.Header.SetContentType("application/json")
	request.Header.SetMethod("POST")
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)
	tls := &tls.Config{
		InsecureSkipVerify: true,
	}
	client := fasthttp.Client{
		ReadTimeout:         30 * time.Second,
		TLSConfig:           tls,
		MaxConnDuration:     30 * time.Second,
		MaxConnWaitTimeout:  30 * time.Second,
		MaxIdleConnDuration: 30 * time.Second,
		MaxConnsPerHost:     3000,
	}
	err := client.Do(request, response)
	if err != nil {
		log.Fatal(err)
	}
	statusCode := response.StatusCode()
	bodyBytes := response.Body
	//responseBody := string(bodyBytes())
	//fmt.Println(fmt.Sprintf("response with statusCode %d with body %s", statusCode, responseBody))

	if statusCode != 200 {
		errorResponseDto := dto.ErrorResponse{}
		err := json.Unmarshal(bodyBytes(), &errorResponseDto)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(errorResponseDto)
	} else {
		person := dto.Person{}
		err := json.Unmarshal(bodyBytes(), &person)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(person)

	}

}

func sendRequestWithHttpNet() {
	tls := &tls.Config{
		InsecureSkipVerify: true,
	}
	tr := &http.Transport{
		MaxIdleConns:        1024,
		MaxIdleConnsPerHost: 1024,
		IdleConnTimeout:     10 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     tls,
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   30 * time.Second,
	}

	request, err := http.NewRequest("GET", "https://localhost:9090/services/person/findAll", nil)
	request.SetBasicAuth("saber66", "saber@123")

	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	statusCode := response.StatusCode
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("response with statusCode %d with body %s", statusCode, string(bodyBytes)))
}

func readConfigFromYamlFile() {
	file, err := ioutil.ReadFile("application.yml")
	if err != nil {
		log.Fatal(err)
	}
	config := dto.Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}

	serverPort := config.Server.Port
	applicationName := config.Gin.Application.Name
	fmt.Printf("Server port %d with application name is %s \n", serverPort, applicationName)

}
