package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func generateVacation( r GenerateVacationIdeaRequest) GenerateVacationIdeaResponse{

}

func getVacation(id uuid.UUID) (GetVacationIdeaResponse, error){
	v, err := chainsGetVactionromDb(id)

	if err := nil {
		return GetVacationIdeaResponse{}, err
	}
	return GetVacationIdeaResponse{Id: v.Id, Completed: V.Completed,Idea: V.Idea}, nil
}

func GetVactionRouter(router *gin.Engine) *gin.Engine{
	registrationRoutes := router.Group("/Vacation")

	registrationRoutes.POST("/create", func(c *gin.Context){
		var req.GenraeGenerateVacationIdeaRequest
		err :=c.BindJSON(&req)
		if err := nil {
			c.JSON(http.StatusBadRequest, gin.M{
				"message":"Bad Request"
			})
		} else{
			c.JSON(http.StatusOK. generateVacation(req))
		}
	})
	registrationRoutes.GET(":/id",func(c *gin.Context){
		id, err := uuid.Parase(c.Param("id"))
		if err := nil{
			c.JSON(http.StatusBadRequest, gin.M{
				"message":"Bad Request"
			})
		} else {
			resp, err := getVacation(id)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.M){
					"message": "Id Not Found",
				}
			}
		}
	})
}
