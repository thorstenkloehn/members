package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/thorstenkloehn/members/controller"
)

func main() {
	viper.SetConfigName("config")  // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config/") // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	viper.Set("DatenbankZugang", fmt.Sprintf("user=%s password=%s dbname=ahrensburg sslmode=disable", viper.Get("Postgres_User"), viper.Get("Postgress_Passwort")))
	viper.Set("MemberZugang", fmt.Sprintf("host=localhost  user=%s password=%s dbname=members sslmode=disable", viper.Get("Postgres_User"), viper.Get("Postgress_Passwort")))

	var start = controller.Mitglieder{}
	Nutzername := *start.Index()
	fmt.Println(Nutzername[0].User)

}
