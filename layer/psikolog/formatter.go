package psikolog

import (
	"project-individu-go-react/entity"
	"time"
)

type PsikologFormat struct {
	ID              int    `json:"id"`
	FirstName       string `json:"firstname"`
	LastName        string `json:"lastname"`
	Title           string `json:"title"`
	Price           int    `json:"price"`
	JenisKonsultasi string `json:"jenis_konsultasi"`
	Description     string `json:"description"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatPsikolog(psikolog entity.Psikologi) PsikologFormat {
	var formatPsikolog = PsikologFormat{
		ID:              psikolog.ID,
		FirstName:       psikolog.Firstname,
		LastName:        psikolog.Lastname,
		Title:           psikolog.Title,
		Price:           psikolog.Price,
		JenisKonsultasi: psikolog.JenisKonsultasi,
		Description:     psikolog.Description,
	}

	return formatPsikolog
}

func FormatDeletePsikolog(msg string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
