package models

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
	CityID   *uint
	City     *City
}
