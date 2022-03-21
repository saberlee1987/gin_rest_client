package services

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gin_rest_client/dto"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

func getTlsConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}
func getClient() fasthttp.Client {
	return fasthttp.Client{
		ReadTimeout:         30 * time.Second,
		TLSConfig:           getTlsConfig(),
		MaxConnDuration:     30 * time.Second,
		MaxConnWaitTimeout:  30 * time.Second,
		MaxIdleConnDuration: 30 * time.Second,
		MaxConnsPerHost:     3000,
	}
}

func getAuthorization(username string, password string) string {
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password))))
}

func FindAllPerson(url string, username string, password string) dto.FindAllPersonResponse {

	authorization := getAuthorization(username, password)
	log.Printf("Request for findAllPerson by  to url %s\n", url)

	request := fasthttp.AcquireRequest()
	request.SetRequestURI(url)
	request.Header.Set("Authorization", authorization)
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	client := getClient()
	err := client.Do(request, response)
	if err != nil {
		log.Fatal(err)
	}
	statusCode := response.StatusCode()
	bodyBytes := response.Body
	responseBody := string(bodyBytes())

	fmt.Println(fmt.Sprintf("response with statusCode %d with body %s", statusCode, responseBody))

	findAllPersonResponse := dto.FindAllPersonResponse{}

	if statusCode != 200 {
		err := json.Unmarshal(bodyBytes(), &findAllPersonResponse.Error)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = json.Unmarshal(bodyBytes(), &findAllPersonResponse)
		if err != nil {
			log.Fatal(err)
		}
	}

	return findAllPersonResponse
}

func FindPersonByNationalCode(url string, nationalCode string, username string, password string) dto.FindPersonByNationalCodeResponseDto {

	authorization := getAuthorization(username, password)
	log.Printf("Request for findPerson by natioanlCode %s to url %s\n", nationalCode, url)
	request := fasthttp.AcquireRequest()
	request.SetRequestURI(url)
	request.Header.Set("Authorization", authorization)
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	client := getClient()
	err := client.Do(request, response)
	if err != nil {
		log.Fatal(err)
	}
	statusCode := response.StatusCode()
	bodyBytes := response.Body
	responseBody := string(bodyBytes())

	fmt.Println(fmt.Sprintf("response for findPerson by natioanlCode %s with statusCode %d with body %s", nationalCode, statusCode, responseBody))

	findPersonByNationalCodeResponseDto := dto.FindPersonByNationalCodeResponseDto{}

	if statusCode != 200 {
		err := json.Unmarshal(bodyBytes(), &findPersonByNationalCodeResponseDto.Error)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = json.Unmarshal(bodyBytes(), &findPersonByNationalCodeResponseDto.Response)
		if err != nil {
			log.Fatal(err)
		}
	}

	return findPersonByNationalCodeResponseDto
}
