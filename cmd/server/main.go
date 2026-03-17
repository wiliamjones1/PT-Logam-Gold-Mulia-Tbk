package main

import (
	"log"
	"os"

	"logam.gold/internal/config"
	"logam.gold/internal/handler"
	"logam.gold/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.Load()

	app := fiber.New(fiber.Config{
		Views:       handler.NewTemplateEngine(),
		ViewsLayout: "layouts/base",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			if code == fiber.StatusNotFound {
				return c.Status(code).Render("errors/404", fiber.Map{
					"Title": "Halaman Tidak Ditemukan - PT Logam Gold Mulia Tbk",
				})
			}
			return c.Status(code).Render("errors/500", fiber.Map{
				"Title": "Kesalahan Server - PT Logam Gold Mulia Tbk",
			})
		},
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(helmet.New(helmet.Config{
		CrossOriginEmbedderPolicy: "unsafe-none",
		CrossOriginResourcePolicy: "cross-origin",
	}))
	app.Use(middleware.SecurityHeaders())

	// Static files
	app.Static("/static", "./static", fiber.Static{
		Compress:  true,
		MaxAge:    86400,
		Browse:    false,
	})

	// Robots.txt & Sitemap
	app.Get("/robots.txt", func(c *fiber.Ctx) error {
		return c.SendFile("./static/robots.txt")
	})
	app.Get("/sitemap.xml", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/xml")
		return c.SendFile("./static/sitemap.xml")
	})

	// Page routes
	h := handler.New(cfg)
	app.Get("/", h.Home)
	app.Get("/tentang", h.Tentang)
	app.Get("/layanan", h.Layanan)
	app.Get("/komitmen", h.Komitmen)
	app.Get("/tata-kelola", h.TataKelola)
	app.Get("/berita", h.Berita)
	app.Get("/faq", h.FAQ)
	app.Get("/kontak", h.Kontak)

	port := cfg.Port
	if port == "" {
		port = "3000"
	}

	log.Printf("Server running on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
}
