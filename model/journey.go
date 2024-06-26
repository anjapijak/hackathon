package model

import "time"

// var CategoryList map[string]struct{} = map[string]struct{}{
// 	"food and drink":    {},
// 	"turist attraction": {},
// 	"entertainment":     {},
// }

type Journey struct {
	Map
	ID       string
	Activity []Activity
	City     string

	Date   time.Time
	Images []string
	Diary  string
	Stay   string //place to sleep
}

func NewJourney(Map Map, Diary, Stay, ID, City string, Activity []Activity) Journey {
	return Journey{
		ID:       ID,
		Map:      Map,
		Activity: Activity,
		Date:     time.Now(),
		City:     City,
		Images:   nil,
		Diary:    Diary,
		Stay:     Stay,
	}
}

func EmptyJourney() Journey {
	return Journey{}
}

func (j *Journey) AppendActivity(Activity Activity) {
	j.Activity = append(j.Activity, Activity)
}

func (j *Journey) AppendImage(url string) {
	j.Images = append(j.Images, url)
}
