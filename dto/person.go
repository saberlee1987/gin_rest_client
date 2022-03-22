package dto

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id           int    `json:"id,omitempty"`
	Firstname    string `json:"firstname,omitempty"`
	Lastname     string `json:"lastname,omitempty"`
	NationalCode string `json:"nationalCode,omitempty"`
	Age          int    `json:"age,omitempty"`
	Email        string `json:"email,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
}
type FindAllPersonResponse struct {
	Persons *[]Person      `json:"persons,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	Code            int              `json:"code,omitempty"`
	Message         string           `json:"message,omitempty"`
	OriginalMessage *interface{}     `json:"originalMessage,omitempty"`
	Validations     *[]ValidationDto `json:"validations,omitempty"`
}

type ValidationDto struct {
	FieldName     string `json:"fieldName,omitempty"`
	DetailMessage string `json:"detailMessage,omitempty"`
}

type DeletePersonResponse struct {
	Code int    `json:"code"`
	Text string `json:"text,omitempty"`
}

type AddPersonResponseDto struct {
	Response *Person        `json:"response,omitempty"`
	Error    *ErrorResponse `json:"error,omitempty"`
}

type FindPersonByNationalCodeResponseDto struct {
	Response *Person        `json:"response,omitempty"`
	Error    *ErrorResponse `json:"error,omitempty"` //null
}

type UpdatePersonResponseDto struct {
	Response *Person        `json:"response,omitempty"`
	Error    *ErrorResponse `json:"error,omitempty"`
}

type DeletePersonResponseDto struct {
	Response *DeletePersonResponse `json:"response,omitempty"`
	Error    *ErrorResponse        `json:"error,omitempty"`
}

func (response AddPersonResponseDto) String() string {
	marshal, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		return fmt.Sprintf("{\"response\":%s,\"error\":%s}", response.Response, response.Error)
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

////////////////////////////////////////////// ToString Methods ///////////////////////

func (error ErrorResponse) String() string {

	marshal, err := json.MarshalIndent(error, "", "\t")
	if err != nil {
		return fmt.Sprintf("{\"code\":\"%d\",\"message\":\"%s\",\"originalMessage\":\"%s\",\"validations\":%s}",
			error.Code, error.Message, error.OriginalMessage, error.Validations)
	}
	return string(marshal)
}

func (error DeletePersonResponse) String() string {

	marshal, err := json.MarshalIndent(error, "", "\t")
	if err != nil {
		return fmt.Sprintf("{\"code\":\"%d\",\"text\":\"%s\"}",
			error.Code, error.Text)
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
func (response FindPersonByNationalCodeResponseDto) String() string {
	marshal, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		return fmt.Sprintf("{\"response\":%s,\"error\":%s}", response.Response, response.Error)
	}
	return string(marshal)
}

func (response UpdatePersonResponseDto) String() string {
	marshal, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		return fmt.Sprintf("{\"response\":%s,\"error\":%s}", response.Response, response.Error)
	}
	return string(marshal)
}

func (response DeletePersonResponseDto) String() string {
	marshal, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		return fmt.Sprintf("{\"response\":%s,\"error\":%s}", response.Response, response.Error)
	}
	return string(marshal)
}

func (response FindAllPersonResponse) String() string {
	marshal, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		return fmt.Sprintf("{\"persons\":%s,\"error\":%s}", response.Persons, response.Error)
	}
	return string(marshal)
}
