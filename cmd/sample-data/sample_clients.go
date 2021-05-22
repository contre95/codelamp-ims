package main

import (
	"codelamp-ims/pkg/domain/clients"
	"time"
)

var projectDurations, err = time.ParseDuration("76hs")

var SampleClients = []clients.Client{
	{
		Name:          "Rugiero Transporte",
		AdmissionDate: time.Now(),
		Website:       "https://transporterugiero.com/",
		Country:       "Argentina",
		Tag:           "transporte-rugiero",
		Projects: []clients.Project{
			{
				Name:       "Rugiero Sistema de Gestión",
				StartDate:  time.Now(),
				FinishDate: time.Now().Add(projectDurations),
				Website:    "https://portal.transporterugiero.com/admin",
				Tag:        "sys-transporterugiero",
				Type:       "sys",
			},
			{
				Name:       "Rugiero Sistema Web",
				StartDate:  time.Now(),
				FinishDate: time.Now().Add(projectDurations),
				Website:    "https://transporterugiero.com",
				Tag:        "web-transporterugiero",
				Type:       "web",
			},
		},
	}, {
		Name:          "Antiplaga Norte",
		AdmissionDate: time.Now(),
		Website:       "https://antiplaganorte.com/",
		Country:       "Argentina",
		Tag:           "antiplaga-norte",
		Projects: []clients.Project{
			{
				Name:       "Antiplaga Sistema de Gestión",
				StartDate:  time.Now(),
				FinishDate: time.Now().Add(projectDurations),
				Website:    "https://portal.antiplaganorte.com/admin",
				Tag:        "sys-antiplaganorte",
				Type:       "sys",
			},
		},
	},
}
