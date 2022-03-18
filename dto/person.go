package dto

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id           int    `json:"id,omitempty"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	NationalCode string `json:"nationalCode"`
	Age          int    `json:"age"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
}
type FindAllPersonResponse struct {
	Persons []Person `json:"persons"`
}

type ErrorResponse struct {
	Code            int             `json:"code"`
	Message         string          `json:"message"`
	OriginalMessage interface{}     `json:"originalMessage,omitempty"`
	Validations     []ValidationDto `json:"validations,omitempty"`
}

type ValidationDto struct {
	FieldName     string `json:"fieldName"`
	DetailMessage string `json:"detailMessage"`
}

func (error ErrorResponse) String() string {

	marshal, err := json.MarshalIndent(error, "", "\t")
	if err != nil {
		return fmt.Sprintf("{\"code\":\"%d\",\"message\":\"%s\",\"originalMessage\":\"%s\",\"validations\":%s}",
			error.Code, error.Message, error.OriginalMessage, error.Validations)
	}
	return string(marshal)
}
func (validationDto ValidationDto) String() string {
	marshal, err := json.MarshalIndent(validationDto, "", "\t")

	if err != nil {
		return fmt.Sprintf("{\"fieldName\":\"%s\",\"detailMessage\":\"%s\"}", validationDto.FieldName, validationDto.DetailMessage)
	}

	return string(marshal)
}

func (person Person) String() string {

	marshal, err := json.MarshalIndent(person, "", "\t")
	if err != nil {
		return fmt.Sprintf("{\"id\":%d,\"firstName\":\"%s\",\"lastName\":\"%s\",\"nationalCode\":\"%s\",\"age\":%d,\"email\":\"%s\",\"mobile\":\"%s\"}", person.Id,
			person.Firstname, person.Lastname, person.NationalCode, person.Age, person.Email, person.Mobile)
	}
	return string(marshal)
}

func (response FindAllPersonResponse) String() string {
	marshal, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		return fmt.Sprintf("{\"persons\":%s}", response.Persons)
	}
	return string(marshal)
}
