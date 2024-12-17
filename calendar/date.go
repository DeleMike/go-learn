package calendar

import "errors"

type Date struct {
	Year, Month, Day int
}

func (d *Date) SetYear(year int) error {
	if year < 1901 {
		return errors.New("year out of range")
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("month out of range")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("day out of range")
	}
	d.Day = day
	return nil
}
