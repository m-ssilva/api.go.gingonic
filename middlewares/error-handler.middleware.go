package middlewares

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// ErrorModel struct
type ErrorModel struct {
	Name    string `json:"name"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func getErrorSchemas() []ErrorModel {
	schemas := make([]ErrorModel, 30)
	raw, _ := ioutil.ReadFile("./helpers/errors-schemas.json")
	json.Unmarshal(raw, &schemas)
	return schemas
}

func getErrorMessage(errorString string, errorModels ...ErrorModel) ErrorModel {
	for _, schema := range errorModels {
		if errorString == schema.Name {
			fmt.Println("ENTROU NO IF")
			errorModel := ErrorModel{schema.Name, schema.Status, schema.Message}
			return errorModel
		}
	}
	fmt.Println("NAO ENTROU NO IF")
	errorModel := ErrorModel{errorModels[0].Name, errorModels[0].Status, errorModels[0].Message}
	return errorModel
}

// ErrorHandle handle errors and returns custom responses
func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errorModels := getErrorSchemas()
		err := c.Errors.Last()
		if err == nil {
			return
		}
		responseMessage := getErrorMessage(err.Error(), errorModels...)
		c.JSON(responseMessage.Status, gin.H{"message": responseMessage.Message})
	}
}
