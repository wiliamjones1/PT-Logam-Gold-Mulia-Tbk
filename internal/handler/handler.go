package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"logam.gold/internal/config"
)

type Handler struct {
	cfg *config.Config
}

type ContactForm struct {
	Nama       string `form:"nama"`
	Email      string `form:"email"`
	Telepon    string `form:"telepon"`
	Perusahaan string `form:"perusahaan"`
	Pesan      string `form:"pesan"`
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
		"Title":       "Hubungi Kami — PT Logam Gold Mulia Tbk",
		"Description": "Hubungi PT Logam Gold Mulia Tbk untuk informasi lebih lanjut, kerja sama, atau pertanyaan seputar layanan kami.",
		"Canonical":   h.cfg.BaseURL + "/kontak",
		"Page":        "kontak",
	})
}

func (h *Handler) KontakSubmit(c *fiber.Ctx) error {
	form := new(ContactForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(http.StatusBadRequest).Render("pages/kontak", fiber.Map{
			"Title":       "Hubungi Kami — PT Logam Gold Mulia Tbk",
			"Description": "Hubungi PT Logam Gold Mulia Tbk untuk informasi lebih lanjut.",
			"Canonical":   h.cfg.BaseURL + "/kontak",
			"Page":        "kontak",
			"Error":       "Terjadi kesalahan saat memproses formulir. Silakan coba lagi.",
		})
	}

	// Validate required fields
	if form.Nama == "" || form.Email == "" || form.Pesan == "" {
		return c.Status(http.StatusBadRequest).Render("pages/kontak", fiber.Map{
			"Title":       "Hubungi Kami — PT Logam Gold Mulia Tbk",
			"Description": "Hubungi PT Logam Gold Mulia Tbk untuk informasi lebih lanjut.",
			"Canonical":   h.cfg.BaseURL + "/kontak",
			"Page":        "kontak",
			"Error":       "Mohon lengkapi semua kolom yang wajib diisi.",
			"Form":        form,
		})
	}

	// Validate email format (basic)
	if !isValidEmail(form.Email) {
		return c.Status(http.StatusBadRequest).Render("pages/kontak", fiber.Map{
			"Title":       "Hubungi Kami — PT Logam Gold Mulia Tbk",
			"Description": "Hubungi PT Logam Gold Mulia Tbk untuk informasi lebih lanjut.",
			"Canonical":   h.cfg.BaseURL + "/kontak",
			"Page":        "kontak",
			"Error":       "Format alamat email tidak valid.",
			"Form":        form,
		})
	}

	// Send email notification
	if h.cfg.ResendAPIKey != "" {
		log.Printf("Sending contact email to %s via Resend API", h.cfg.ContactEmail)
		if err := h.sendContactEmail(form); err != nil {
			log.Printf("Failed to send contact email: %v", err)
			return c.Render("pages/kontak", fiber.Map{
				"Title":       "Hubungi Kami — PT Logam Gold Mulia Tbk",
				"Description": "Hubungi PT Logam Gold Mulia Tbk untuk informasi lebih lanjut.",
				"Canonical":   h.cfg.BaseURL + "/kontak",
				"Page":        "kontak",
				"Error":       "Maaf, terjadi kesalahan saat mengirim pesan. Silakan coba lagi nanti.",
				"Form":        form,
			})
		}
	} else {
		log.Printf("Contact form submission (email not configured): Nama=%s, Email=%s, Pesan=%s", form.Nama, form.Email, form.Pesan)
	}

	return c.Render("pages/kontak", fiber.Map{
		"Title":       "Hubungi Kami — PT Logam Gold Mulia Tbk",
		"Description": "Hubungi PT Logam Gold Mulia Tbk untuk informasi lebih lanjut.",
		"Canonical":   h.cfg.BaseURL + "/kontak",
		"Page":        "kontak",
		"Success":     "Terima kasih! Pesan Anda telah berhasil dikirim. Tim kami akan menghubungi Anda dalam waktu dekat.",
	})
}

func (h *Handler) sendContactEmail(form *ContactForm) error {
	to := h.cfg.ContactEmail
	subject := fmt.Sprintf("[Logam Gold] Pesan dari %s", form.Nama)

	var body strings.Builder
	body.WriteString(fmt.Sprintf("Nama: %s\n", form.Nama))
	body.WriteString(fmt.Sprintf("Email: %s\n", form.Email))
	if form.Telepon != "" {
		body.WriteString(fmt.Sprintf("Telepon: %s\n", form.Telepon))
	}
	if form.Perusahaan != "" {
		body.WriteString(fmt.Sprintf("Perusahaan: %s\n", form.Perusahaan))
	}
	body.WriteString(fmt.Sprintf("\nPesan:\n%s\n", form.Pesan))

	payload := map[string]interface{}{
		"from":    h.cfg.EmailFrom,
		"to":      []string{to},
		"subject": subject,
		"text":    body.String(),
		"reply_to": form.Email,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.resend.com/emails", bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+h.cfg.ResendAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("resend API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	log.Printf("Email sent successfully to %s", to)
	return nil
}

func isValidEmail(email string) bool {
	if len(email) < 5 || len(email) > 254 {
		return false
	}
	at := false
	dot := false
	for i, c := range email {
		if c == '@' {
			if at || i == 0 {
				return false
			}
			at = true
		}
		if at && c == '.' && i > 0 {
			dot = true
		}
	}
	return at && dot
}
