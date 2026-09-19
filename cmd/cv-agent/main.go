package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/domain"
	"awesomeProject/internal/fileutil"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "config error:", err)
		os.Exit(1)
	}
	if cfg.OpenRouterAPIKey == "" {
		fmt.Fprintln(os.Stderr, "error: OPENROUTER_API_KEY is required")
		os.Exit(1)
	}

	if cfg.OpenRouterModel == "" {
		cfg.OpenRouterModel = "openrouter/free"
	}

	cvPath := flag.String("cv", "", "path to CV")
	vacancyPath := flag.String("vacancy", "", "path to Vacancy")
	country := flag.String("country", "", "country")
	language := flag.String("language", "", "language")

	flag.Parse()

	if *cvPath == "" || *vacancyPath == "" || *language == "" || *country == "" {
		fmt.Println("error provide all required params")
		os.Exit(1)
	}

	cvData, err := fileutil.ReadFile(*cvPath)

	if err != nil {
		fmt.Println("error reading CV file:", err)
		os.Exit(1)
	}

	vacancyData, err := fileutil.ReadFile(*vacancyPath)

	if err != nil {
		fmt.Println("error reading Vacancy file:", err)
		os.Exit(1)
	}

	request := domain.TailorRequest{
		CVText:      cvData,
		VacancyText: vacancyData,
		Country:     *country,
		Language:    *language,
	}

	mockResponse := []byte(`{
		"summary": "Backend developer with experience in Laravel and TypeScript.",
		"matched_skills": ["REST APIs", "PostgreSQL", "TypeScript"],
		"missing_skills": ["Go", "Docker"],
		"cover_letter": "I am applying for the backend developer position..."
	}`)

	var result domain.TailorResult

	if err := json.Unmarshal(mockResponse, &result); err != nil {
		fmt.Println("error parsing mock response:", err)
		os.Exit(1)
	}

	if err := request.Validate(); err != nil {
		fmt.Fprintln(os.Stderr, "error validating request:", err)
		os.Exit(1)
	}

	payload, err := json.MarshalIndent(request, "", " ")

	if err != nil {
		fmt.Fprintln(os.Stderr, "error encoding request:", err)
		os.Exit(1)
	}

	fmt.Println("Request JSON size:", len(payload))

	fmt.Println("Summary:", result.Summary)
	fmt.Println("Cover letter:", result.CoverLetter)

	fmt.Println("Matched skills:")
	for _, skill := range result.MatchedSkills {
		fmt.Println("-", skill)
	}

	fmt.Println("Missing skills:")
	for _, skill := range result.MissingSkills {
		fmt.Println("-", skill)
	}
}
