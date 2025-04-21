package main

import (
	"flag"
	"fmt"
	"log"
	"yamanmnur/simple-dashboard/internal/injectors"
	"yamanmnur/simple-dashboard/internal/middlewares"
	"yamanmnur/simple-dashboard/pkg/db"
	"yamanmnur/simple-dashboard/pkg/util"
	"yamanmnur/simple-dashboard/seeder"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	flag.Parse()

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	dbUrl := viper.Get("DB_URL").(string)
	APP_PORT := viper.Get("APP_PORT").(string)

	dbPgsql, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	dbHandler := db.Init(dbPgsql)

	client := util.InitMinio()

	db.InitInstanceDbHandler(&db.IDbHandler{
		DB: dbHandler,
	})

	app := fiber.New()
	app.Use(cors.New())
	app.Use(func(ctx *fiber.Ctx) error {
		return middlewares.ErrorHandler(ctx, ctx.Next())
	})

	injectors.InitDI(app, db.GetDbHandler(), client)

	seeder.InitSeed()

	app.Listen(fmt.Sprintf(":%s", APP_PORT))
}
