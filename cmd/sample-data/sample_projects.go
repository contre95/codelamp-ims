package main

import (
	"codelamp-ims/pkg/domain/projects"
	"time"
)

var projectDurations, err = time.ParseDuration("76hs")

var sampleProjecs = []projects.Project{
	{
		Name:       "Rugiero Transportes",
		StartDate:  time.Now(),
		FinishDate: time.Now().Add(projectDurations),
		Website:    "https://portal.transporterugiero/admin",
		Tag:        "sys-transporterugiero",
		Type:       "sys",
	},
	{
		Name:       "Antiplaga Norte",
		StartDate:  time.Now(),
		FinishDate: time.Now().Add(projectDurations),
		Website:    "https://portal.antiplaganorte.com/admin",
		Tag:        "sys-antiplaganorte",
		Type:       "sys",
	},
	{
		Name:       "Master TV",
		StartDate:  time.Now(),
		FinishDate: time.Now().Add(projectDurations),
		Website:    "https://mastertv.com.ar",
		Tag:        "imp-mastertv",
		Type:       "imp",
	},
}
