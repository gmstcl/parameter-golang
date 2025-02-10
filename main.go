package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

var (
	framework string
	port      = flag.Int("p", 8888, "Listen port")
)

func init() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatal("e.g. http, echo, fiber")
	}
	framework = flag.Arg(0)
}
func main() {
	switch framework {
	case "http":
		RunNewHttpServer()
	case "fiber":
		RunNewFiberServer()
	default:
		log.Fatal("Not framework.")
	}
}

func RunNewHttpServer() {
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Server is listening %d", *port)
	http.HandleFunc("/check", func(w http.ResponseWriter, req *http.Request) {
		if _, err := w.Write([]byte("by net/http\n")); err != nil {
			log.Print(err)
		}
	})

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Print(err)
	}
}

func RunNewFiberServer() {
	addr := fmt.Sprintf(":%d", *port)
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pingpong by fiber\n")
	})
	log.Printf("Server is listening %d", *port)
	if err := app.Listen(addr); err != nil {
		log.Print(err)
	}
}