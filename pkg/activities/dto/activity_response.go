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

type ActivityDeleteResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    *Blank `json:"data"`
}

type Blank struct {
	ID int64 `json:"id,omitempty"`
}
