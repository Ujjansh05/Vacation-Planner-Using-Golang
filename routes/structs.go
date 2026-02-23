package routes

import "github.com/google/uuid"


type GenerateVacationIdeaRequest struct{
FavouriteSeason string `json:"favourite_season"`
	Hobbies []string `json:"hobbies"`
	Budget int `json:"budget"`
}

type GenerateVacationIdeaResponse struct{
Id  		uuid.UUID 		`json:"id"`
Completed	bool			`json:"completed"`
}


type GetVacationIdeaResponse struct{
Id				uuid.UUID		`json:"id"`
Completed		bool 			`json:"completed"`
Idea			string      	`json:"idea"`
}


