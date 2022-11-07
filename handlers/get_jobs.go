package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"ridhodinar/dans-tes/models"
	"ridhodinar/dans-tes/restapi/operations/job"
	"ridhodinar/dans-tes/utils"

	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
	"gorm.io/gorm"
)

type getAllJobsHandlerImpl struct {
	dbClient *gorm.DB
}

func NewGetAllJobsHandler(db *gorm.DB) job.GetAllJobsHandler {
	return &getAllJobsHandlerImpl{
		dbClient: db,
	}
}

func (impl *getAllJobsHandlerImpl) Handle(params job.GetAllJobsParams, principal interface{}) middleware.Responder {
	_, err := utils.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return job.NewGetAllJobsInternalServerError().WithPayload("error in parsing token")
	}
	jobs, err := getJobsData(params)
	if err != nil {
		log.Errorf(err.Error())
		return job.NewGetAllJobsInternalServerError().WithPayload("Error getting info")
	}
	return job.NewGetAllJobsOK().WithPayload(*jobs)
}

func getJobsData(params job.GetAllJobsParams) (*models.JobList, error) {
	var err error
	var client = &http.Client{}
	var data = &models.JobList{}

	url := "http://dev3.dansmultipro.co.id/api/recruitment/positions.json"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	addQuery(q, "description", params.Description)
	addQuery(q, "location", params.Location)
	addQuery(q, "full_time", params.FullTime)
	addQuery(q, "page", params.Page)
	request.URL.RawQuery = q.Encode()

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

func addQuery(q url.Values, key string, value *string) {
	if value != nil {
		q.Add(key, *value)
	}
}
