package dtoActivity

import entityActivity "github.com/nach9/go-todolist/pkg/activities/entity"

type ActivityResponse struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Data    entityActivity.Activity `json:"data"`
}

type ActivityListResponse struct {
	Status  string                    `json:"status"`
	Message string                    `json:"message"`
	Data    []entityActivity.Activity `json:"data"`
}
