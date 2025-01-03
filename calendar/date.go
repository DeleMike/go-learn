package calendar

import "errors"

type Date struct {
	year, month, day int
}

func (d *Date) SetYear(year int) error {
	if year < 1901 {
		return errors.New("year out of range")
	}
	d.year = year
	return nil
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("month out of range")
	}
	d.month = month
	return nil
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("day out of range")
	}
	d.day = day
	return nil
}

func (d *Date) Day() int {
	return d.day
}
