package domain

import (
	"errors"
	"strings"
)

type TailorRequest struct {
	CVText      string `json:"cv_text"`
	VacancyText string `json:"vacancy_text"`
	Country     string `json:"country"`
	Language    string `json:"language"`
}

type TailorResult struct {
	Summary       string   `json:"summary"`
	MatchedSkills []string `json:"matched_skills"`
	MissingSkills []string `json:"missing_skills"`
	CoverLetter   string   `json:"cover_letter"`
}

func (r TailorRequest) Validate() error {
	if strings.TrimSpace(r.CVText) == "" {
		return errors.New("CV text is required")
	}

	if strings.TrimSpace(r.VacancyText) == "" {
		return errors.New("vacancy text is required")
	}

	if strings.TrimSpace(r.Language) == "" {
		return errors.New("language is required")
	}

	if strings.TrimSpace(r.Country) == "" {
		return errors.New("country is required")
	}

	return nil
}
