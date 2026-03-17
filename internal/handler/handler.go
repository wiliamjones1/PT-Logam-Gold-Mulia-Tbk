package handler

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"logam.gold/internal/config"
)

type Handler struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Handler {
	return &Handler{cfg: cfg}
}

func NewTemplateEngine() *html.Engine {
	engine := html.New("./templates", ".html")
	engine.AddFunc("safe", func(s string) template.HTML {
		return template.HTML(s)
	})
	engine.AddFunc("year", func() int {
		return 2026
	})
	engine.Reload(true)
	return engine
}

func (h *Handler) Home(c *fiber.Ctx) error {
	return c.Render("pages/home", fiber.Map{
		"Title":       "PT Logam Gold Mulia Tbk — Mitra Terpercaya dalam Industri Logam Mulia",
		"Description": "PT Logam Gold Mulia Tbk adalah perusahaan terbuka yang berkomitmen dalam industri logam mulia dengan mengedepankan kualitas, integritas, dan kepercayaan.",
		"Canonical":   h.cfg.BaseURL,
		"Page":        "home",
	})
}

func (h *Handler) Tentang(c *fiber.Ctx) error {
	return c.Render("pages/tentang", fiber.Map{
		"Title":       "Tentang Perusahaan — PT Logam Gold Mulia Tbk",
		"Description": "Pelajari lebih lanjut tentang PT Logam Gold Mulia Tbk, visi misi, nilai-nilai perusahaan, dan komitmen kami dalam industri logam mulia.",
		"Canonical":   h.cfg.BaseURL + "/tentang",
		"Page":        "tentang",
	})
}

func (h *Handler) Layanan(c *fiber.Ctx) error {
	return c.Render("pages/layanan", fiber.Map{
		"Title":       "Layanan & Bisnis Kami — PT Logam Gold Mulia Tbk",
		"Description": "Jelajahi layanan dan lini bisnis PT Logam Gold Mulia Tbk dalam perdagangan, distribusi, dan solusi kemitraan logam mulia.",
		"Canonical":   h.cfg.BaseURL + "/layanan",
		"Page":        "layanan",
	})
}

func (h *Handler) Komitmen(c *fiber.Ctx) error {
	return c.Render("pages/komitmen", fiber.Map{
		"Title":       "Komitmen Kami — PT Logam Gold Mulia Tbk",
		"Description": "Komitmen PT Logam Gold Mulia Tbk terhadap kualitas, kepercayaan nasabah, etika bisnis, dan tata kelola perusahaan yang baik.",
		"Canonical":   h.cfg.BaseURL + "/komitmen",
		"Page":        "komitmen",
	})
}

func (h *Handler) TataKelola(c *fiber.Ctx) error {
	return c.Render("pages/tata-kelola", fiber.Map{
		"Title":       "Tata Kelola Perusahaan — PT Logam Gold Mulia Tbk",
		"Description": "Informasi tata kelola perusahaan yang baik (GCG) PT Logam Gold Mulia Tbk, mencakup prinsip transparansi, akuntabilitas, dan kepatuhan.",
		"Canonical":   h.cfg.BaseURL + "/tata-kelola",
		"Page":        "tata-kelola",
	})
}

func (h *Handler) Berita(c *fiber.Ctx) error {
	return c.Render("pages/berita", fiber.Map{
		"Title":       "Berita & Wawasan — PT Logam Gold Mulia Tbk",
		"Description": "Berita terbaru, wawasan industri, dan informasi korporasi dari PT Logam Gold Mulia Tbk.",
		"Canonical":   h.cfg.BaseURL + "/berita",
		"Page":        "berita",
	})
}

func (h *Handler) FAQ(c *fiber.Ctx) error {
	return c.Render("pages/faq", fiber.Map{
		"Title":       "FAQ — PT Logam Gold Mulia Tbk",
		"Description": "Pertanyaan yang sering diajukan seputar PT Logam Gold Mulia Tbk, layanan, dan informasi perusahaan.",
		"Canonical":   h.cfg.BaseURL + "/faq",
		"Page":        "faq",
	})
}

func (h *Handler) Kontak(c *fiber.Ctx) error {
	return c.Render("pages/kontak", fiber.Map{
		"Title":        "Hubungi Kami — PT Logam Gold Mulia Tbk",
		"Description":  "Hubungi PT Logam Gold Mulia Tbk untuk informasi lebih lanjut, kerja sama, atau pertanyaan seputar layanan kami.",
		"Canonical":    h.cfg.BaseURL + "/kontak",
		"Page":         "kontak",
		"Web3FormsKey": h.cfg.Web3FormsKey,
	})
}


