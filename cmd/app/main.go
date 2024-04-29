package main

import (
	"github.com/banggibima/go-fiber-jwt-rbac/config"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/interface/http"
	"github.com/banggibima/go-fiber-jwt-rbac/pkg/fiber"
	"github.com/banggibima/go-fiber-jwt-rbac/pkg/gorm"
	"github.com/banggibima/go-fiber-jwt-rbac/pkg/postgres"
	"github.com/banggibima/go-fiber-jwt-rbac/pkg/redis"
	"github.com/banggibima/go-fiber-jwt-rbac/pkg/viper"
)

func main() {
	v, err := viper.New()
	if err != nil {
		panic(err)
	}

	c, err := config.Init(v)
	if err != nil {
		panic(err)
	}

	p, err := postgres.New(c)
	if err != nil {
		panic(err)
	}

	g, err := gorm.New(p)
	if err != nil {
		panic(err)
	}

	r, err := redis.New(c)
	if err != nil {
		panic(err)
	}

	f, err := fiber.New(c)
	if err != nil {
		panic(err)
	}

	if err := http.NewServer(c, g, r, f).Start(); err != nil {
		panic(err)
	}
}
