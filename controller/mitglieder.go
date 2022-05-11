package controller

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/thorstenkloehn/members/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Mitglieder struct {
	Mitglieder model.Mitglieders
}

func (start *Mitglieder) Test() string {
	dsn := viper.GetString("MemberZugang")
	fmt.Println(dsn)
	_, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		return "Kein Zugriff auf "
	} else {
		return "Sie haben vollen Zugriff"
	}
}

func (start Mitglieder) Index() *[]model.Mitglieders {
	dsn := viper.GetString("MemberZugang")
	datenbank, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Fatalln(error)
	}
	var mitglieder []model.Mitglieders
	datenbank.Find(&mitglieder)
	return &mitglieder

}
