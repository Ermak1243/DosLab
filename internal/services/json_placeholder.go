package services

import (
	"encoding/json"
	"fmt"
	"main/internal/models"
)

type placeholder struct {
	commentsUrl string
	request     *request
}

func NewPlaceholder() *placeholder {
	return &placeholder{
		commentsUrl: "https://jsonplaceholder.typicode.com/comments",
		request:     NewRequest(),
	}
}

func (p *placeholder) Comments() []models.Comment {
	var model []models.Comment

	resp, err := p.request.Get(p.commentsUrl)
	if err != nil {
		return model
	}

	defer resp.Body.Close()

	if resp.Body == nil {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Response Body = nil!")

		return model
	}

	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		fmt.Println("JSON decode error: ", err)

		return model
	}

	return model
}
