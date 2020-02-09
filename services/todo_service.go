package services

import (
	"encoding/json"
	"fmt"
	"github.com/mersanuzun/go-bff-boilerplate/config"
	"github.com/mersanuzun/go-bff-boilerplate/helpers"
	"github.com/mersanuzun/go-bff-boilerplate/models/todo"
	"github.com/parnurzeal/gorequest"
)

func GetTodo(requestDto *todo.GetTodoRequest) todo.GetTodoResponse {
	var todoResponse todo.GetTodoResponse

	endpoint := fmt.Sprint(config.JsonPlaceHolderApi(), "/todos/", requestDto.ID)
	request := gorequest.New().
		Get(endpoint)

	_, body := helpers.HandleClientResponse(request, 200)

	if err := json.Unmarshal([]byte(body), &todoResponse); err != nil {
		panic(err)
	}

	return todoResponse
}