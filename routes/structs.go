package routes


import uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"

type GenerateVacationIdeaRequest struct{
	Favourite Season string `json: "favourite_season"`
	Hobbies []string `json:"hobbies"`
	Budget int `json:"budget"`
}

type GenerateVacationIdeaResponse struct{
	Id  		uuid.UUID 		`json:"ID"`
	Completed	bool			`json:"Completed"`
}


type GetVacationIdeaResponse struct{
	Id				uuid.UUID		`json:"id"`
	Completed		bool 			`json: "completed"`
	Idea			string      	`json:"idea"`
}