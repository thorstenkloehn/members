package model

type Mitglieders struct {
	ID       uint `gorm:"primaryKey"`
	User     string
	Passwort string
	Vorname  string
	Nachname string
}

type Session struct {
	ID               uint `gorm:"primaryKey"`
	MitgliedersRefer uint
	Mitglieders      Mitglieders `gorm:"foreignKey:MitgliedersRefer"`
	Session_Name     string
}
