package main

import (
	"fmt"
	"time"
)

type Period struct {
	StartDate time.Time
	EndDate   time.Time
}

func main() {
	start := "2021-01-15"
	end := "2021-04-11"

	periodList, err := ParceRange(start, end)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, period := range periodList {
			fmt.Println(fmt.Sprintf("%s ~ %s", period.StartDate.Format("2006-01-02"), period.EndDate.Format("2006-01-02")))
			//fmt.Println(period)
		}
	}
}

func ParceRange(start, end string) ([]Period, error) {
	var (
		startDate, endDate time.Time
		periodList         []Period
		err                error
	)
	startDate, err = time.Parse("2006-01-02", start)
	if err != nil {
		return nil, err
	}
	endDate, err = time.Parse("2006-01-02", end)
	if err != nil {
		return nil, err
	}

	firstPeriod := Period{
		StartDate: startDate,
		EndDate:   time.Date(startDate.Year(), startDate.Month(), 1, 0, 0, 0, 0, startDate.Location()).AddDate(0, 1, -1),
	}
	periodList = append(periodList, firstPeriod)
	startDate = time.Date(startDate.Year(), startDate.Month(), 1, 0, 0, 0, 0, startDate.Location())
	for {

		startDate = startDate.AddDate(0, 1, 0)
		if Equal(startDate, endDate) {
			periodList = append(periodList, formPeriod(startDate, endDate))
			break
		} else {
			periodList = append(periodList, formPeriod(startDate, startDate.AddDate(0, 1, -1)))
		}
	}
	return periodList, nil
}

func formPeriod(start, end time.Time) Period {
	return Period{
		StartDate: start,
		EndDate:   end,
	}
}

func Equal(start, end time.Time) bool {
	return start.Year() == end.Year() && start.Month() == end.Month()
}
