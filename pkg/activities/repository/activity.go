package repositoryActivity

import (
	entityActivity "github.com/nach9/go-todolist/pkg/activities/entity"
	"gorm.io/gorm"
)

type ActivityRepo interface {
	FindAll() []entityActivity.Activity
	FindById(id int64) (entityActivity.Activity, error)
	Save(activity entityActivity.Activity) (entityActivity.Activity, error)
	Delete(id int64) error
}

type activityRepo struct {
	DB *gorm.DB
}

func NewActivityrepo(DB *gorm.DB) ActivityRepo {
	return &activityRepo{DB}
}

func (r *activityRepo) FindAll() []entityActivity.Activity {
	var activities []entityActivity.Activity

	r.DB.Find(&activities)

	return activities
}

func (r *activityRepo) FindById(id int64) (entityActivity.Activity, error) {
	var activity entityActivity.Activity
	result := r.DB.First(&activity, id)

	if result.Error != nil {
		return activity, result.Error
	}

	return activity, nil
}

func (r *activityRepo) Save(activity entityActivity.Activity) (entityActivity.Activity, error) {
	result := r.DB.Save(&activity)

	if result.Error != nil {
		return activity, result.Error
	}

	return activity, nil
}

func (r *activityRepo) Delete(id int64) error {
	var activity entityActivity.Activity

	result := r.DB.Delete(&activity, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
