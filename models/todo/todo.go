package todo

type GetTodoRequest struct {
	ID int32 `json:"int32"`
}

type Todo struct {
	UserId int32 `json:"userId"`
	ID int32 `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
	Response int16
}

type GetTodoResponse struct {
	UserId int32 `json:"userId"`
	ID int32 `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}
