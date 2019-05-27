# Statistic server

# How to run

1. run  `docker-compose up -d postgres` and wait
2. docker-compose up -d flyway
3. docker-compose up -d server-stats-backend


# Configuration

.env is supported

all params in config.go(replace by link)


# Api

1. GET /healthcheck

    Response example:
    ```
    ok
    ```

2. POST /api/users 
    
    Create user
    request data:
    ```
    {
      "id": 33,
      "sex": "M",
      "age": 1
    }
    ```
    sex filed:
    * M
    * W
    
    age is between 1 and 100
    
    id should be greater than 1
    
    returns 201 status code in success case

3. POST /api/users/stats

    Create user's stat
    ```
    {
      "user": 1,
      "action": "like",
      "ts": "2018-11-30T18:12:34"
    }
    ```
    Possible action values:
    * login
    * logout
    * like
    * comments
    
    returns 201 status code in success case
    
4. GET /api/users/stats/top?date1=2017-06-30&date2=2019-06-30&action=like&limit=1
    
    All query params are required. 
    Get most active by action top=`limit` users per day, statistics is considered for each day separately 
    
    Response example:
    ```
    {
        "items":
            [
                {
                    "date":"2017-06-30",
                    "rows":[{"age":1,"count":4,"id":2,"sex":"M"}]
                },
                {
                    "date":"2018-06-30",
                    "rows":[{"age":1,"count":2,"id":3,"sex":"M"}]
                },
                {
                    "date":"2018-07-30",
                    "rows":[{"age":1,"count":1,"id":3,"sex":"M"}]
                },
                {
                    "date":"2018-08-30",
                    "rows":[{"age":1,"count":1,"id":3,"sex":"M"}]},
                },
                ...
            ]
    }
    ```
    
5. GET /api/users/stats/topAccum?date1=2011-06-20&date2=2019-06-30&action=like&limit=1
    
    All query params are required. 
    Get most active by action top=`limit` users per day, statistics is accumulated
F.E in first day I have 1 actions of type login, in next day I have 2 actions of type login. Then the answer is wiil be like 
for first day is 1, for second day is 3. 

    Response example:
    ```
    {
        "items":
            [
                {
                    "date":"2017-06-30",
                    "rows":[{"age":1,"count":4,"id":2,"sex":"M"}]
                },
                {
                    "date":"2018-06-30",
                    "rows":[{"age":1,"count":4,"id":2,"sex":"M"}]
                },
                {
                    "date":"2018-07-30",
                    "rows":[{"age":1,"count":4,"id":3,"sex":"M"}]
                },
                ...
            ]
    }
    ```