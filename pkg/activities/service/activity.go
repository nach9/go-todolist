package serviceActivity

import (
	dtoActivity "github.com/nach9/go-todolist/pkg/activities/dto"
	entityActivity "github.com/nach9/go-todolist/pkg/activities/entity"
	repositoryActivity "github.com/nach9/go-todolist/pkg/activities/repository"
)

type ActivityService interface {
	GetAll() []entityActivity.Activity
	GetById(id int64) (entityActivity.Activity, error)
	Create(body dtoActivity.CreateActivityBody) (entityActivity.Activity, error)
	UpdateById(id int64, body dtoActivity.UpdateActivityBody) (entityActivity.Activity, error)
	DeleteById(id int64) (entityActivity.Activity, error)
}

type activityService struct {
	repo repositoryActivity.ActivityRepo
}

func NewActivityService(repo repositoryActivity.ActivityRepo) ActivityService {
	return &activityService{repo}
}

func (s activityService) GetAll() []entityActivity.Activity {
	return s.repo.FindAll()
}

func (s activityService) GetById(id int64) (entityActivity.Activity, error) {
	return s.repo.FindById(id)
}

func (s activityService) Create(body dtoActivity.CreateActivityBody) (entityActivity.Activity, error) {
	newActivity := entityActivity.NewActivity(body.Title, body.Email)

	return s.repo.Save(newActivity)
}

func (s activityService) UpdateById(id int64, body dtoActivity.UpdateActivityBody) (entityActivity.Activity, error) {
	activity, err := s.repo.FindById(id)

	if err != nil {
		return activity, err
	}

	activity.Title = body.Title

	return s.repo.Save(activity)
}

func (s activityService) DeleteById(id int64) (entityActivity.Activity, error) {
	activity, err := s.repo.FindById(id)

	if err != nil {
		return activity, err
	}

	errDelete := s.repo.Delete(id)

	if errDelete != nil {
		return activity, errDelete
	}

	return activity, nil
}
