package rest

type Client struct {
	//ID   string
	Name          string `json:"name"`
	AdmissionDate string `json:"admission_date"`
	FinishDate    string `json:"finish_date"`
	Country       string `json:"country"`
	Website       string `json:"website"`
	Tag           string `json:"tag"`
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
