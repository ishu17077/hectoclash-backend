package main

import (
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/ishu17077/hectoclash-backend/middlewares"

	"github.com/ishu17077/hectoclash-backend/middlewares"
	"github.com/ishu17077/hectoclash-backend/routes"
	// calculator "github.com/mnogu/go-calculator"
)

func main() {

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

}
