package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/domain"
	"awesomeProject/internal/fileutil"
	"awesomeProject/internal/openrouter"
	"context"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
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

	if err := request.Validate(); err != nil {
		fmt.Fprintln(os.Stderr, "error validating request:", err)
		os.Exit(1)
	}

	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "config error:", err)
		os.Exit(1)
	}

	aiClient := openrouter.NewClient(cfg.OpenRouterAPIKey, cfg.OpenRouterModel)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	responseData, err := aiClient.Complete(
		ctx,
		"Reply with exactly: connection works",
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Open router request failed:", err)
		os.Exit(1)
	}

	fmt.Println(string(responseData))

}
