package main // import "github.com/jiaxuyang/go-vue-starter"

import (
	"github.com/jiaxuyang/go-vue-starter/api"
	"github.com/jiaxuyang/go-vue-starter/models"
	"github.com/jiaxuyang/go-vue-starter/routes"
	"github.com/urfave/negroni"
)

func main() {
	db := models.NewSqliteDB("data.db")
	api := api.NewAPI(db)
	routes := routes.NewRoutes(api)
	n := negroni.Classic()
	n.UseHandler(routes)
	n.Run(":8081")
}
