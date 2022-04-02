package main

import (
	"fmt"
	"gin_rest_client/config"
	"gin_rest_client/consul"
	"gin_rest_client/docs"
	"gin_rest_client/services"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
)

// @title gin rest client
// @version 1.0.0-1401/01/13
// @description gin rest client
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email saberazizi66@yahoo.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9099
// @BasePath /service/gin-rest-client
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	fmt.Println("Hello World @@@@@")

	applicationConfig := config.ReadConfigFromYamlFile()

	registerConsul(applicationConfig)

	router := gin.Default()

	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(router)

	router.Use(CORSMiddleware())
	server := applicationConfig.Server
	api := applicationConfig.Service.Api

	/// config swagger docs
	docs.SwaggerInfo_swagger.BasePath = api.BasePath
	docs.SwaggerInfo_swagger.Title = api.SwaggerTitle
	docs.SwaggerInfo_swagger.Description = api.SwaggerTitle
	docs.SwaggerInfo_swagger.Version = api.SwaggerVersion
	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d%s", server.Port, api.SwaggerPath)) // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	service := services.GetPersonService()

	personRoute := router.Group(api.BasePath + "/person")
	{
		personRoute.GET("/findAll", findAllPerson(service))
		personRoute.GET("/find/:nationalCode", findPersonByNationalCode(service))
		//personRoute.PUT("/update/:nationalCode", updatePersonByNationalCode)
		//personRoute.DELETE("/delete/:nationalCode", deletePersonByNationalCode)
		//personRoute.POST("/add", addPerson)
	}

	router.Run(fmt.Sprintf(":%d", server.Port))

}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// HealthCheck godoc
// @Summary findAllPerson
// @Description get the status of server.
// @Tags person api
// @Accept */*
// @Produce json
// @Success 200 {object}  dto.FindAllPersonResponse
// @Router /person/findAll [get]
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func findAllPerson(service services.PersonService) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		response, statusCode := service.FindAllPerson()
		if statusCode != 200 {
			errorResponse := response.Error
			context.JSON(statusCode, errorResponse)
		} else {
			context.JSON(statusCode, response)
		}
	}
	return fn
}

// HealthCheck godoc
// @Summary findPersonByNationalCode
// @Description get the status of server.
// @Tags person api
// @Accept */*
// @Param nationalCode path string true "nationalCode param"
// @Produce json
// @Success 200 {object}  dto.FindPersonByNationalCodeResponseDto
// @Failure 400,404,406,500,504 {object} dto.ErrorResponse
// @Router /person/find/{nationalCode} [get]
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func findPersonByNationalCode(service services.PersonService) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		nationalCode := context.Param("nationalCode")
		response, statusCode := service.FindPersonByNationalCode(nationalCode)
		if statusCode != 200 {
			errorResponse := response.Error
			context.JSON(statusCode, errorResponse)
		} else {
			context.JSON(statusCode, response)
		}
	}
	return fn
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
