package routes

import (
	"net/http"
	"langchaingo/chains"
	"github.com/gin-gonic/gin"
	"github.com/google.uuid"
)

func generateVacation( r GenerateVacationIdeaRequest) GenerateVacationIdeaResponse{
		id := uuid.New()
		go chains.GenerateVacationIdeaChange(id, r.Budget, r.FavouriteSeason, r.Hobbies)
		return GenerateVacationIdeaResponse{Id : id, Completed: false}
}

func getVacation(id uuid.UUID) (GetVacationIdeaResponse, error){
	v, err := chains.GetVacationFromDb(id)

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
				c.JSON(http.StatusNotFound, gin.M{
					"message": "Id Not Found",
				})
			} else {
				c.JSON(http.StatusOk, resp)
			}
		}
	})
	return router
}
