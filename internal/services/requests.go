package services

import (
	"log"
	"net/http"
	"time"
)

type request struct{}

func NewRequest() *request {
	return &request{}
}

func (r *request) Get(requestUrl string) (http.Response, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(requestUrl)
	if err != nil {
		log.Println(err)

		return http.Response{}, err
	}

	return *resp, err
}
