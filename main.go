package main

import (
	"fmt"
	"gin_rest_client/config"
	"gin_rest_client/consul"
	"gin_rest_client/dto"
	"gin_rest_client/services"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("Hello World @@@@@")

	applicationConfig := readConfigFromYamlFile()

	registerConsul(applicationConfig)

	//sendRequestWithHttpNet()
	//personUrl := personClient.URL
	//findAllPerson()
	//sendRequestWithFastHttp2("0079028748")
	//
	//person := dto.Person{
	//	Firstname:    "saber",
	//	Lastname:     "Azizi",
	//	NationalCode: "0079028748",
	//	Age:          34,
	//	Email:        "saberazizi66@yahoo.com",
	//	Mobile:       "09365627895",
	//}
	//sendRequestWithFastHttp3(person)

	personClient := applicationConfig.Service.PersonClient
	nationalCode := "0079028748"

	person := dto.Person{
		Firstname:    "saber",
		Lastname:     "Azizi",
		NationalCode: "0079028748",
		Age:          34,
		Email:        "saberazizi66@yahoo.com",
		Mobile:       "09365627895",
	}

	service := services.PersonService{}
	service.PersonClient = personClient

	//findAllPersonResponse := service.FindAllPerson()
	//addPersonResponseDto := service.AddPerson(person)
	//service.AddPerson(person)
	person.Firstname = "saber66"
	person.Lastname = "Azizi22"
	updatePerson := service.UpdatePerson(nationalCode, person)
	fmt.Println(updatePerson)
	findPersonByNationalCodeResponse := service.FindPersonByNationalCode(nationalCode)
	fmt.Println(findPersonByNationalCodeResponse)
	//deletePerson := service.DeletePerson(nationalCode)
	//fmt.Println(findAllPersonResponse)
	//fmt.Println(findPersonByNationalCodeResponse)
	//fmt.Println(addPersonResponseDto)
	//fmt.Println(deletePerson.Response.Code)
	//fmt.Println(deletePerson.Response.Code)
	//fmt.Println(deletePerson)
	//fmt.Println(addPersonResponseDto)

}

func readConfigFromYamlFile() config.Config {
	file, err := ioutil.ReadFile("application.yml")
	if err != nil {
		log.Fatal(err)
	}
	config := config.Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}

	serverPort := config.Server.Port
	applicationName := config.Gin.Application.Name
	fmt.Printf("Server port %d with application name is %s \n", serverPort, applicationName)

	return config
}

func registerConsul(config config.Config) {
	application := config.Gin.Application
	c := config.Gin.Consul
	port := config.Server.Port

	client, err := consul.NewConsulClient(c.Host, c.Port)
	if err != nil {
		log.Println("Error for get client consul with error ====> " + err.Error())
	}
	err = client.Register(application.Name, port)
	if err != nil {
		log.Println("Error for register consul with error ====> " + err.Error())
	} else {
		log.Printf("%s  register successfully in consul by address http://%s:%d", application.Name, c.Host, c.Port)
	}

}
