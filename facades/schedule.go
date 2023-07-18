package facades

import (
	"github.com/mewway/go-laravel/contracts/schedule"
)

func Schedule() schedule.Schedule {
	return App().MakeSchedule()
}
