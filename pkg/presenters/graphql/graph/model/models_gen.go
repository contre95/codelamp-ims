// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Client struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Country  string     `json:"country"`
	Tag      string     `json:"tag"`
	Projects []*Project `json:"projects"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewClient struct {
	Name          string `json:"name"`
	AdmissionDate string `json:"admissionDate"`
	FinishDate    string `json:"finishDate"`
	Website       string `json:"website"`
	Country       string `json:"country"`
	Tag           string `json:"tag"`
}

type NewProject struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type User struct {
	Name string `json:"name"`
}
