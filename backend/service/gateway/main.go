package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nats-io/nats.go"

	jwtware "github.com/gofiber/contrib/jwt"
)

func main() {
	log.Println("Gateway")

	nc, err := nats.Connect("nats://queue:4222", nats.Token("qwerty"))
	if err != nil {
		log.Fatal(err)
	}

	srv := fiber.New()

	srv.Post("/token", tokenHandler())

	srv.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	srv.Post("/rpc", rpcHandler(nc))

	if err := srv.Listen(":80"); err != nil {
		log.Fatalln(err)
	}

	// signalChan := make(chan os.Signal, 1)
	// signal.Notify(signalChan, os.Interrupt)

	// <-signalChan

	// fmt.Println("Done")
}

// rpcHandler ...
func rpcHandler(nc *nats.Conn) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Locals("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		userId := claims["sub"].(string)

		log.Println(userId)

		resp, err := nc.Request("test", ctx.Body(), 5*time.Second)
		if err != nil {
			log.Println(err)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		} else {
			return ctx.Send(resp.Data)
		}
	}
}

func tokenHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		claims := jwt.MapClaims{
			"sub":  "alex",
			"name": "Alex Matsuev",
			"exp":  time.Now().Add(72 * time.Hour).Unix(),
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		return ctx.JSON(fiber.Map{"token": t})
	}
}
