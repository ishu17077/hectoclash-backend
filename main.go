package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	// "github.com/ishu17077/hectoclash-backend/middlewares"

	"github.com/ishu17077/hectoclash-backend/middlewares"
	"github.com/ishu17077/hectoclash-backend/routes"
	// calculator "github.com/mnogu/go-calculator"
)

func main() {
	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, os.Interrupt, os.Kill)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middlewares.CORSMiddleware())
	routes.UserRoutes(router)
	//? get user should not be called without a refreshtoken
	routes.MatchRoutes(router)
	routes.ProblemRoutes(router)
	routes.MatchProblemRoutes(router)
	routes.PlayerScorecardRoutes(router)
	routes.ProblemStatusRoutes(router)
	routes.ProblemTimeRoutes(router)
	// calculator.Calculate()
	router.SetTrustedProxies([]string{"127.0.0.1", "localhost", "192.168.232.61"}) //! Remember to change this
	router.Run(":" + port)
	<-stopChannel
	log.Print("Shutting down Server at 8080")
	/*
	* Code here to reserve previous matches
	*/

}
