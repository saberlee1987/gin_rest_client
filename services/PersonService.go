package services

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gin_rest_client/config"
	"gin_rest_client/dto"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

type PersonService struct {
	PersonClient config.PersonClient
}

func GetPersonService() PersonService {
	service := PersonService{}
	serverConfig := config.ReadConfigFromYamlFile()
	service.PersonClient = serverConfig.Service.PersonClient
	return service
}

func (personService PersonService) getTlsConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}
func (personService PersonService) getClient() fasthttp.Client {
	connectionConfig := personService.PersonClient.ConnectionConfig

	return fasthttp.Client{
		ReadTimeout:         time.Duration(connectionConfig.ReadTimeout) * time.Millisecond,
		TLSConfig:           personService.getTlsConfig(),
		MaxConnDuration:     time.Duration(connectionConfig.ConnectionDuration) * time.Millisecond,
		MaxConnWaitTimeout:  time.Duration(connectionConfig.ConnectTimeout) * time.Millisecond,
		MaxIdleConnDuration: time.Duration(connectionConfig.ConnectionDuration) * time.Millisecond,
		MaxConnsPerHost:     connectionConfig.ConnectionPerHost,
	}
}

func (personService PersonService) getAuthorization() string {
	authorization := personService.PersonClient.Authorization
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", authorization.Username, authorization.Password))))
}

func (personService PersonService) FindAllPerson() (dto.FindAllPersonResponse, int) {

	authorization := personService.getAuthorization()
	personClient := personService.PersonClient
	url := fmt.Sprintf("%s:%d%s%s", personClient.URL, personClient.Port, personClient.Baseurl, personClient.FindAll)

	log.Printf("Request for findAllPerson by  to url %s\n", url)

	request := fasthttp.AcquireRequest()
	request.SetRequestURI(url)
	request.Header.Set("Authorization", authorization)
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	client := personService.getClient()
	err := client.Do(request, response)
	if err != nil {
		log.Fatal(err)
	}
	statusCode := response.StatusCode()
	bodyBytes := response.Body
	responseBody := string(bodyBytes())

	log.Printf("Response with statusCode %d with body %s", statusCode, responseBody)

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

	return findAllPersonResponse, statusCode
}

func (personService PersonService) FindPersonByNationalCode(nationalCode string) (dto.FindPersonByNationalCodeResponseDto, int) {

	authorization := personService.getAuthorization()
	personClient := personService.PersonClient
	url := fmt.Sprintf("%s:%d%s%s/%s", personClient.URL, personClient.Port, personClient.Baseurl, personClient.FindByNationalCode, nationalCode)
	log.Printf("Request for findPerson by natioanlCode %s to url %s\n", nationalCode, url)
	request := fasthttp.AcquireRequest()
	request.SetRequestURI(url)
	request.Header.Set("Authorization", authorization)
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	client := personService.getClient()
	err := client.Do(request, response)
	if err != nil {
		log.Fatal(err)
	}
	statusCode := response.StatusCode()
	bodyBytes := response.Body
	responseBody := string(bodyBytes())

	log.Printf("Response for findPerson by natioanlCode %s with statusCode %d with body %s", nationalCode, statusCode, responseBody)

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

	return findPersonByNationalCodeResponseDto, statusCode
}

func (personService PersonService) AddPerson(personDto dto.Person) dto.AddPersonResponseDto {

	authorization := personService.getAuthorization()
	personClient := personService.PersonClient
	url := fmt.Sprintf("%s:%d%s%s", personClient.URL, personClient.Port, personClient.Baseurl, personClient.AddPerson)
	log.Printf("Request for addPerson by body %s  to url %s\n", personDto, url)

	request := fasthttp.AcquireRequest()
	request.SetRequestURI(url)
	request.SetBodyString(personDto.String())
	request.Header.Set("Authorization", authorization)
	request.Header.SetContentType("application/json")
	request.Header.SetMethod("POST")
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	client := personService.getClient()
	err := client.Do(request, response)
	if err != nil {
		log.Fatal(err)
	}
	statusCode := response.StatusCode()
	bodyBytes := response.Body
	responseBody := string(bodyBytes())

	log.Printf("Response for addPerson with statusCode %d with body %s\n", statusCode, responseBody)

	addPersonResponse := dto.AddPersonResponseDto{}

	if statusCode != 200 {
		err := json.Unmarshal(bodyBytes(), &addPersonResponse.Error)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := json.Unmarshal(bodyBytes(), &addPersonResponse.Response)
		if err != nil {
			log.Fatal(err)
		}
	}
	return addPersonResponse
}

func (personService PersonService) DeletePerson(nationalCode string) dto.DeletePersonResponseDto {

	authorization := personService.getAuthorization()
	personClient := personService.PersonClient
	url := fmt.Sprintf("%s:%d%s%s/%s", personClient.URL, personClient.Port, personClient.Baseurl, personClient.DeletePerson, nationalCode)
	log.Printf("Request for deletePerson by natioanlCode %s to url %s\n", nationalCode, url)
	request := fasthttp.AcquireRequest()
	request.SetRequestURI(url)
	request.Header.Set("Authorization", authorization)
	request.Header.SetMethod("DELETE")
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	client := personService.getClient()
	err := client.Do(request, response)
	if err != nil {
		log.Fatal(err)
	}
	statusCode := response.StatusCode()
	bodyBytes := response.Body
	responseBody := string(bodyBytes())

	log.Printf("Response for deletePerson by natioanlCode %s with statusCode %d with body %s", nationalCode, statusCode, responseBody)

	deletePersonResponseDto := dto.DeletePersonResponseDto{}

	if statusCode != 200 {
		err := json.Unmarshal(bodyBytes(), &deletePersonResponseDto.Error)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := json.Unmarshal(bodyBytes(), &deletePersonResponseDto.Response)
		if err != nil {
			log.Fatal(err)
		}
	}
	return deletePersonResponseDto

}

func (personService PersonService) UpdatePerson(nationalCode string, personDto dto.Person) dto.UpdatePersonResponseDto {
	authorization := personService.getAuthorization()
	personClient := personService.PersonClient
	url := fmt.Sprintf("%s:%d%s%s/%s", personClient.URL, personClient.Port, personClient.Baseurl, personClient.UpdatePerson, nationalCode)
	log.Printf("Request for updatePerson by natioanlCode %s  with body %sto url %s\n", nationalCode, personDto, url)
	request := fasthttp.AcquireRequest()
	request.SetRequestURI(url)
	request.Header.Set("Authorization", authorization)
	request.Header.SetContentType("application/json")
	request.Header.SetMethod("PUT")
	request.SetBodyString(personDto.String())
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	client := personService.getClient()
	err := client.Do(request, response)
	if err != nil {
		log.Fatal(err)
	}
	statusCode := response.StatusCode()
	bodyBytes := response.Body
	responseBody := string(bodyBytes())

	log.Printf("Response for updatePerson by natioanlCode %s with statusCode %d with body %s", nationalCode, statusCode, responseBody)

	updatePersonResponse := dto.UpdatePersonResponseDto{}

	if statusCode != 200 {
		err := json.Unmarshal(bodyBytes(), &updatePersonResponse.Error)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := json.Unmarshal(bodyBytes(), &updatePersonResponse.Response)
		if err != nil {
			log.Fatal(err)
		}
	}
	return updatePersonResponse
}
