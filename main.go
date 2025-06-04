package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/ishu17077/hectoclash-backend/middlewares"

	"github.com/ishu17077/hectoclash-backend/controllers"
	"github.com/ishu17077/hectoclash-backend/middlewares"
	"github.com/ishu17077/hectoclash-backend/routes"
	// calculator "github.com/mnogu/go-calculator"
)

func main() {
	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

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
	router.SetTrustedProxies([]string{"127.0.0.1", "localhost", "192.168.232.61"}) //! Remember to change this

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	<-stopChannel

	ctx, _ := context.WithTimeout(context.TODO(), 600*time.Second)
	fmt.Println("Shutting down Server at 8080")
	fmt.Println("Attempting to pause all current matches")
	//? A logic to prevent condition to pause a game if server becomes unusable for some reason.
	if controllers.AddPauseToCurrentMatches() {
		fmt.Println("All Matches Paused Successfully")
	} else {
		fmt.Println("All Matches Pause Unsuccessful")
	}

	if err := server.Shutdown(ctx); err != nil {
		if err == http.ErrServerClosed {
			fmt.Printf("Server closed under request")
		} else {
			log.Fatal("Server connection closed unexpectedly")
		}
	} else {
		server.Close()
	}
	//* Gin router may not be necessarily required to shut gracefully as it has its own in built protection
}
