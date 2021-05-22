package http

import (
	"codelamp-ims/pkg/domain/clients"
	"time"
)

type Client struct {
	//ID   string
	Name          string `json:"name"`
	AdmissionDate string `json:"admission_date"`
	FinishDate    string `json:"finish_date"`
	Country       string `json:"country"`
	Website       string `json:"website"`
	Tag           string `json:"tag"`
	Contacts      []uint `json:"contacts"`
	//Projects      []Project
}

type Project struct {
	//ID            string `json:"id"`
	Name          string `json:"name"`
	StartDate     string `json:"start_date"`
	FinishDate    string `json:"finish_date"`
	Website       string `json:"website"`
	GitRepository string `json:"git_repository"`
	Type          string `json:"type"`
	State         string `json:"state"`
	Tag           string `json:"tag"`
	//Contacts []contacts.ContactID `json:"contacts"`
}

func parseJSONClient(jc Client) (*clients.Client, error) {
	var err error
	var admissionDate, finishDate time.Time
	admissionDate, err = time.Parse(time.RFC3339, jc.AdmissionDate)
	if jc.FinishDate != "" { // Mmm ...
		finishDate, err = time.Parse(time.RFC3339, jc.FinishDate)
	}
	if err != nil {
		return nil, err
	}
	client := &clients.Client{
		Name:          jc.Name,
		AdmissionDate: admissionDate,
		FinishDate:    finishDate,
		Website:       jc.Website,
		Country:       jc.Country,
		Tag:           jc.Tag,
	}

	err = client.Validate()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func parseDomainClient(dc clients.Client) *Client {
	panic("Implement me")
}

func parseDomainProject(dp Project) *clients.Project {
	panic("Implement me")
}

func parseJSONProject(jp clients.Project) *Project {
	panic("Implement me")
}
