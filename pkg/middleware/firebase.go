package middleware

import (
	"context"
	"fmt"
	"log"
	"native/crus/pkg/helper"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

var authClient *auth.Client

func init() {
	firebaseCred := "cred.json"

	ctx := context.Background()
	opt := option.WithCredentialsFile(firebaseCred)
	conf := &firebase.Config{
		ProjectID: "eduez-6ea18",
	}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		panic(err)
	}

	authClient, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		panic(err)
	}
}

func FirebaseAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")
		if header == "" {
			return helper.ResponseJson(c, 401, "Firebase 401 1", nil)
		}
		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return helper.ResponseJson(c, 401, "Firebase 401 2", nil)
		}
		token := parts[1]
		// fmt.Println(token)

		claims, err := authClient.VerifyIDToken(context.Background(), token)
		if err != nil {
			return helper.ResponseJson(c, 401, "Firebase 401 3", err.Error())
		}

		c.Locals("uid", claims.UID)

		// print id and name
		fmt.Println("Id " + claims.UID)

		return c.Next()
	}
}
