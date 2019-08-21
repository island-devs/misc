package main

import (
	"finanzas/misc/app"
	"os"

	"github.com/urfave/cli"
)

// TODO: Define config structure

// var (
// 	config *models.Config
// )

func main() {

	app := cli.NewApp()
	// config = &models.Config{}

	// app.Flags = getFlags()
	app.Name = "Finanzas Miscellaneous service"
	// app.Version = config.Version
	app.Usage = "Misc Service will host all miscellaneous information, from items used for billing to currency and conversion rate info."
	app.Action = run

	app.Run(os.Args)

}

// func getFlags() []cli.Flag {
// 	return []cli.Flag{
// 		cli.IntFlag{
// 			Name:        "port,p",
// 			Usage:       "Port for server to listen on.",
// 			Value:       9000,
// 			EnvVar:      "PORT",
// 			Destination: &config.Port,
// 		},
// 		cli.StringFlag{
// 			Name:        "conn-string,c",
// 			Usage:       "Database Connection String",
// 			EnvVar:      "CONN_STRING",
// 			Destination: &config.ConnectionString,
// 		},
// 		cli.StringFlag{
// 			Name:        "env,e",
// 			Usage:       "Environment",
// 			Value:       "DEV",
// 			EnvVar:      "ENV",
// 			Destination: &config.Env,
// 		},
// 		cli.StringFlag{
// 			Name:        "app-version,ver",
// 			Usage:       "Application Version",
// 			Value:       "0.0.0",
// 			EnvVar:      "VERSION",
// 			Destination: &config.Version,
// 		},
// 		cli.StringFlag{
// 			Name:        "db-name,d",
// 			Usage:       "Database Name",
// 			Value:       "cortes-programados",
// 			EnvVar:      "DB_NAME",
// 			Destination: &config.DatabaseName,
// 		},
// 		cli.StringFlag{
// 			Name:        "schedule,s",
// 			Usage:       "Cron Job Schedule",
// 			Value:       "@weekly",
// 			EnvVar:      "JOB_SCHEDULE",
// 			Destination: &config.Schedule,
// 		},
// 	}
// }

func run(c *cli.Context) error {

	return app.Main()
}
