# TVDB Go Client

[![API Version](https://img.shields.io/badge/TVDB%20API-v4.7.10-blue)](https://thetvdb.github.io/v4-api/)

A Go client library for [TheTVDB API V4](https://thetvdb.com/).

## Installation

```bash
go get github.com/mxrk/tvdb
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mxrk/tvdb"
)

func main() {
	// Create a new client with your API key
	client := tvdb.New("your-api-key")

	// Login to get a bearer token (valid for 1 month)
	ctx := context.Background()
	if err := client.Login(ctx); err != nil {
		log.Fatal(err)
	}

	// Search for a series
	results, err := client.Search(ctx, tvdb.SearchOptions{
		Query: "Breaking Bad",
		Type:  "series",
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range results.Data {
		fmt.Printf("Found: %s (%s)\n", result.Name, result.Year)
	}

	// Get series details
	series, err := client.GetSeriesBase(ctx, 81189)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Series: %s\n", series.Name)

	// Get extended series info with translations
	extended, err := client.GetSeriesExtended(ctx, 81189, &tvdb.SeriesExtendedOptions{
		Meta: "translations",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Genres: %v\n", extended.Genres)
}
```

## Options Pattern

Many methods accept optional parameters:

```go
// Get series with short response (no characters/artworks)
series, err := client.GetSeriesExtended(ctx, id, &tvdb.SeriesExtendedOptions{
	Short: true,
})

// Filter movies
movies, err := client.GetMoviesFilter(ctx, tvdb.MovieFilterOptions{
	Country: "usa",
	Lang:    "eng",
	Genre:   &genreId,
	Year:    &year,
})
```
