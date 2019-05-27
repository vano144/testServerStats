package statistics

type User struct {
	Id  int    `json:"id"`
	Age int    `json:"age"`
	Sex string `json:"sex"`
}

type Stat struct {
	User   int    `json:"user"`
	Action string `json:"action"`
	Ts     string `json:"ts"`
}


