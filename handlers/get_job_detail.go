package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ridhodinar/dans-tes/models"
	"ridhodinar/dans-tes/restapi/operations/job"
	"ridhodinar/dans-tes/utils"

	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
	"gorm.io/gorm"
)

type getJobDetailHandlerImpl struct {
	dbClient *gorm.DB
}

func NewGetJobDetailHandler(db *gorm.DB) job.GetJobDetailHandler {
	return &getJobDetailHandlerImpl{
		dbClient: db,
	}
}

func (impl *getJobDetailHandlerImpl) Handle(params job.GetJobDetailParams, principal interface{}) middleware.Responder {
	_, err := utils.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return job.NewGetJobDetailInternalServerError().WithPayload("error in parsing token")
	}
	jobs, err := getJobDetail(params)
	if err != nil {
		log.Errorf(err.Error())
		return job.NewGetJobDetailInternalServerError().WithPayload("Error getting info")
	}
	return job.NewGetJobDetailOK().WithPayload(jobs)
}

func getJobDetail(params job.GetJobDetailParams) (*models.JobDetail, error) {
	var err error
	var client = &http.Client{}
	var data = &models.JobDetail{}

	url := fmt.Sprintf("http://dev3.dansmultipro.co.id/api/recruitment/positions/%s", params.ID)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
