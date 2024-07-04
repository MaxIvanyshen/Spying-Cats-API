package validation

import (
    "github.com/go-resty/resty/v2"
    "log"
    "sync"
    "encoding/json"
)

var (
    breeds []string
    once   sync.Once
)

func FetchBreeds() ([]string, error) {
    once.Do(func() {
        client := resty.New()
        resp, err := client.R().
            SetHeader("Content-Type", "application/json").
            Get("https://api.thecatapi.com/v1/breeds")
        if err != nil {
            log.Fatalf("Failed to fetch cat breeds: %v", err)
            return
        }

        var result []struct {
            Name string `json:"name"`
        }
        err = json.Unmarshal(resp.Body(), &result)
        if err != nil {
            log.Fatalf("Failed to parse cat breeds: %v", err)
            return
        }

        for _, breed := range result {
            breeds = append(breeds, breed.Name)
        }
    })

    return breeds, nil
}

func IsValidBreed(breed string) (bool, error) {
    breeds, err := FetchBreeds()
    if err != nil {
        return false, err
    }
    for _, b := range breeds {
        if b == breed {
            return true, nil
        }
    }
    return false, nil
}
