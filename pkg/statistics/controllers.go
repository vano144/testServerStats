package statistics

import (
	"testServerStats/pkg/db"
)

func createUserController(user *User) error {
	userRepository := &UserRepositoryModel{db.Db}
	isExists, err := userRepository.checkUserExists(user.Id)
	if isExists {
		return &EntityAlreadyExists{id: user.Id, modelName:"user"}
	}
	if err != nil {
		return err
	}
	err = userRepository.createUser(user)
	if err != nil {
		return err
	}
	return nil
}

func createStatController(stat *Stat) error {
	userRepository := &UserRepositoryModel{db.Db}
	isExists, err := userRepository.checkUserExists(stat.User)
	if  !isExists {
		return &EntityDoesNotAlreadyExists{id: stat.User, modelName:"user"}
	}
	if err != nil {
		return err
	}
	statRepository := &StatRepositoryModel{db.Db}
	err = statRepository.createStat(stat)
	if err != nil {
		return err
	}
	return nil
}

func getAccumulateStatController(date1, date2, action string, limit int) (map[string]interface{}, error) {
	statRepository := &StatRepositoryModel{db.Db}
	result, err := statRepository.getAccumulateStats(date1, date2, action, limit)
	return result, err
}

func getStatPerDayController(date1, date2, action string, limit int) (map[string]interface{}, error) {
	statRepository := &StatRepositoryModel{db.Db}
	result, err := statRepository.getStatsPerDay(date1, date2, action, limit)
	return result, err
}
