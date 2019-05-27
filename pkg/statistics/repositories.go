package statistics

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"testServerStats/pkg/db"
)

type UserRepository interface {
	checkUserExists(userId int) (bool, error)
	createUser(user *User) error
}

type StatRepository interface {
	createStat(stat *Stat) error
	getAccumulateStats(date1, date2, action string, limit int) (map[string]interface{}, error)
	getStatsPerDay(date1, date2, action string, limit int) (map[string]interface{}, error)
}

type UserRepositoryModel struct{
	conn *sqlx.DB
}

type StatRepositoryModel struct{
	conn *sqlx.DB
}

func (u *UserRepositoryModel) checkUserExists(userId int) (bool, error) {
	var user User
	e := db.Db.Get(&user, "SELECT * FROM users where id=$1", userId)
	if e != nil {
		// TODO: find more proper way
		if e.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false, e
	}
	if user.Id != 0 {
		return true, nil
	}
	return false, nil
}

func (u *UserRepositoryModel) createUser(user *User) error {
	createStatsQuery := "INSERT INTO users (id, age, sex) VALUES ($1, $2, $3)"
	db.Db.MustExec(createStatsQuery, user.Id, user.Age, user.Sex)
	return nil
}

func (s *StatRepositoryModel) createStat(stat *Stat) error {
	createStatsQuery := "INSERT INTO stats_values (user_id, action, ts) VALUES ($1, $2, $3)"
	db.Db.MustExec(createStatsQuery, stat.User, stat.Action, stat.Ts)
	return nil
}

func (s *StatRepositoryModel) getAccumulateStats(date1, date2, action string, limit int) (map[string]interface{}, error) {
	getStatsQuery := `
			WITH zero_actions AS (
				SELECT DISTINCT
					sv.user_id,
					0 as sum_actions,
					to_char(s.ts, 'YYYY-MM-DD') as day 
				FROM stats_values s
				CROSS JOIN stats_values sv
				where s.ts >= $1
					AND s.ts < $2 AND s.action = $3
				ORDER BY 3, 1
			),
			user_actions AS (
				select 
					user_id, 
					count(action) sum_actions,
					to_char(ts, 'YYYY-MM-DD') as day 
				from stats_values
				where ts >= $1
					AND ts < $2 AND stats_values.action = $3
				group by to_char(ts, 'YYYY-MM-DD'), user_id
			),
			finsh_actions AS (
				SELECT
					za.user_id,
					za.sum_actions + COALESCE(ua.sum_actions, 0) as sum_actions,
					za.day
				FROM zero_actions za
				LEFT JOIN user_actions ua ON za.user_id = ua.user_id and za.day = ua.day
			),
			stat AS (
				select
					d.day,
					d.sum_actions,
					SUM(d.sum_actions) OVER (w ORDER BY d.day) as total_count_from_begin_to_currnt_day,
					d.user_id
				from finsh_actions d
				WINDOW w AS (PARTITION BY d.user_id)
				ORDER BY d.day
			)
			select
				jsonb_build_object(
					'rows', jsonb_agg(
						jsonb_build_object(
							'id', u.id,
							'age', u.age,
							'sex', u.sex,
							'count', a.total_count_from_begin_to_currnt_day
						) ORDER BY a.total_count_from_begin_to_currnt_day desc
					),
					'date', a.day 
				)
			from (
				SELECT
					row_number() OVER (w ORDER BY d.total_count_from_begin_to_currnt_day desc) as crow_number,
					d.day,
					user_id,
					d.total_count_from_begin_to_currnt_day
				from stat d
				WINDOW w AS (PARTITION BY d.day)) a 
			JOIN users u ON a.user_id = u.id
			where a.crow_number <= $4
			GROUP BY a.day
			ORDER BY a.day
			`
	var dbResult []interface{}

	err := db.Db.Select(&dbResult, getStatsQuery, date1, date2, action, limit)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})

	items := make([]interface{}, len(dbResult))
	for i, v := range dbResult {
		err = json.Unmarshal(v.([]byte), &items[i])
		if err != nil {
			return nil, err
		}
	}
	result["items"] = items
	return result, err
}

func (s *StatRepositoryModel) getStatsPerDay(date1, date2, action string, limit int) (map[string]interface{}, error) {
	getStatsQuery := `
		WITH user_actions AS (
			select 
				user_id, 
				count(action) as total_actions_per_day,
				to_char(ts, 'YYYY-MM-DD') as day 
			from stats_values
			where ts >= $1
				AND ts < $2 AND stats_values.action = $3
			group by to_char(ts, 'YYYY-MM-DD'), user_id
		)
		
		select
			jsonb_build_object(
				'rows', jsonb_agg(
					jsonb_build_object(
						'id', u.id,
						'age', u.age,
						'sex', u.sex,
						'count', a.total_actions_per_day
					) ORDER BY a.total_actions_per_day desc
				),
				'date', a.day 
			)
		from (
			SELECT
				row_number() OVER (w ORDER BY d.total_actions_per_day desc) as crow_number,
				d.day,
				user_id,
				d.total_actions_per_day
			from user_actions d
			WINDOW w AS (PARTITION BY d.day)) a 
		JOIN users u ON a.user_id = u.id
		where a.crow_number <= $4
		GROUP BY a.day
		ORDER BY a.day
	`

	var dbResult []interface{}

	err := db.Db.Select(&dbResult, getStatsQuery, date1, date2, action, limit)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})

	items := make([]interface{}, len(dbResult))
	for i, v := range dbResult {
		err = json.Unmarshal(v.([]byte), &items[i])
		if err != nil {
			return nil, err
		}
	}
	result["items"] = items
	return result, err
}
