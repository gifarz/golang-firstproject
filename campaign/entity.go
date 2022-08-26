package campaign

import (
	"time"
)

type Campaign struct {
	Id                int
	Name              string
	User_ID           string
	Description       string
	Short_description string
	Goal_amount       string
	Current_amount    string
	Perks             string
	Backer_count      string
	Slug              string
	Created_at        time.Time
	Updated_at        time.Time
}
