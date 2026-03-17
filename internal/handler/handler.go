package handler

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"logam.gold/internal/company"
	"logam.gold/internal/config"
)

type Handler struct {
	cfg  *config.Config
	corp *company.Data
}

func New(cfg *config.Config) *Handler {
	return &Handler{
		cfg:  cfg,
		corp: company.Default(),
	}
}

func NewTemplateEngine() *html.Engine {
	engine := html.New("./templates", ".html")
	engine.AddFunc("safe", func(s string) template.HTML {
		return template.HTML(s)
	})
	engine.AddFunc("year", func() int {
		return 2026
	})
	engine.AddFunc("fallback", func(val, fb string) string {
		if val == "" {
			return fb
		}
		return val
	})
	engine.AddFunc("hasItems", func(items interface{}) bool {
		switch v := items.(type) {
		case []company.Person:
			return len(v) > 0
		case []company.Document:
			return len(v) > 0
		case []company.Disclosure:
			return len(v) > 0
		case []company.ExternalLink:
			return len(v) > 0
		case []company.NewsItem:
			return len(v) > 0
		case []company.FAQ:
			return len(v) > 0
		case []string:
			return len(v) > 0
		default:
			return false
		}
	})
	engine.Reload(true)
	return engine
}

func (h *Handler) base(page string, extra fiber.Map) fiber.Map {
	m := fiber.Map{
		"Page": page,
		"Corp": h.corp,
		"Cfg":  h.cfg,
	}
	for k, v := range extra {
		m[k] = v
	}
	return m
}

func (h *Handler) Home(c *fiber.Ctx) error {
	return c.Render("pages/home", h.base("home", fiber.Map{
		"Title":       "PT Logam Gold Mulia Tbk — Perusahaan Terbuka Industri Logam Mulia",
		"Description": "PT Logam Gold Mulia Tbk adalah perusahaan terbuka yang berkomitmen di industri logam mulia dengan prinsip tata kelola yang baik, transparansi, dan integritas.",
		"Canonical":   h.cfg.BaseURL,
	}))
}

func (h *Handler) Tentang(c *fiber.Ctx) error {
	return c.Render("pages/tentang", h.base("tentang", fiber.Map{
		"Title":       "Profil Perusahaan — PT Logam Gold Mulia Tbk",
		"Description": "Profil lengkap PT Logam Gold Mulia Tbk meliputi identitas korporasi, visi misi, nilai-nilai perusahaan, dan informasi entitas hukum.",
		"Canonical":   h.cfg.BaseURL + "/tentang",
	}))
}

func (h *Handler) Layanan(c *fiber.Ctx) error {
	return c.Render("pages/layanan", h.base("layanan", fiber.Map{
		"Title":       "Lini Bisnis & Layanan — PT Logam Gold Mulia Tbk",
		"Description": "Informasi lini bisnis dan layanan PT Logam Gold Mulia Tbk dalam perdagangan, distribusi, dan kemitraan strategis logam mulia.",
		"Canonical":   h.cfg.BaseURL + "/layanan",
	}))
}

func (h *Handler) TataKelola(c *fiber.Ctx) error {
	return c.Render("pages/tata-kelola", h.base("tata-kelola", fiber.Map{
		"Title":       "Tata Kelola Perusahaan — PT Logam Gold Mulia Tbk",
		"Description": "Penerapan prinsip Good Corporate Governance (GCG) PT Logam Gold Mulia Tbk: transparansi, akuntabilitas, responsibilitas, independensi, dan kewajaran.",
		"Canonical":   h.cfg.BaseURL + "/tata-kelola",
	}))
}

func (h *Handler) InvestorRelations(c *fiber.Ctx) error {
	return c.Render("pages/investor", h.base("investor", fiber.Map{
		"Title":       "Hubungan Investor — PT Logam Gold Mulia Tbk",
		"Description": "Informasi emiten, keterbukaan informasi, dan dokumen korporasi PT Logam Gold Mulia Tbk untuk investor dan pemangku kepentingan.",
		"Canonical":   h.cfg.BaseURL + "/investor",
	}))
}

func (h *Handler) Dokumen(c *fiber.Ctx) error {
	return c.Render("pages/dokumen", h.base("dokumen", fiber.Map{
		"Title":       "Dokumen & Laporan — PT Logam Gold Mulia Tbk",
		"Description": "Pusat dokumen resmi PT Logam Gold Mulia Tbk meliputi laporan tahunan, laporan keuangan, dan dokumen tata kelola perusahaan.",
		"Canonical":   h.cfg.BaseURL + "/dokumen",
	}))
}

func (h *Handler) Berita(c *fiber.Ctx) error {
	return c.Render("pages/berita", h.base("berita", fiber.Map{
		"Title":       "Berita & Pengumuman — PT Logam Gold Mulia Tbk",
		"Description": "Berita terbaru, pengumuman korporasi, dan keterbukaan informasi dari PT Logam Gold Mulia Tbk.",
		"Canonical":   h.cfg.BaseURL + "/berita",
	}))
}

func (h *Handler) FAQ(c *fiber.Ctx) error {
	return c.Render("pages/faq", h.base("faq", fiber.Map{
		"Title":       "Pertanyaan Umum (FAQ) — PT Logam Gold Mulia Tbk",
		"Description": "Pertanyaan yang sering diajukan seputar PT Logam Gold Mulia Tbk, layanan, tata kelola, dan informasi korporasi.",
		"Canonical":   h.cfg.BaseURL + "/faq",
	}))
}

func (h *Handler) Kontak(c *fiber.Ctx) error {
	return c.Render("pages/kontak", h.base("kontak", fiber.Map{
		"Title":       "Hubungi Kami — PT Logam Gold Mulia Tbk",
		"Description": "Hubungi PT Logam Gold Mulia Tbk untuk pertanyaan umum, kerja sama bisnis, hubungan investor, atau informasi korporasi.",
		"Canonical":   h.cfg.BaseURL + "/kontak",
	}))
}

func (h *Handler) KebijakanPrivasi(c *fiber.Ctx) error {
	return c.Render("pages/privasi", h.base("privasi", fiber.Map{
		"Title":       "Kebijakan Privasi — PT Logam Gold Mulia Tbk",
		"Description": "Kebijakan privasi PT Logam Gold Mulia Tbk terkait pengelolaan data dan informasi pengguna situs.",
		"Canonical":   h.cfg.BaseURL + "/kebijakan-privasi",
	}))
}

func (h *Handler) SyaratPenggunaan(c *fiber.Ctx) error {
	return c.Render("pages/syarat", h.base("syarat", fiber.Map{
		"Title":       "Syarat & Ketentuan — PT Logam Gold Mulia Tbk",
		"Description": "Syarat dan ketentuan penggunaan situs resmi PT Logam Gold Mulia Tbk.",
		"Canonical":   h.cfg.BaseURL + "/syarat-penggunaan",
	}))
}

func (h *Handler) Disclaimer(c *fiber.Ctx) error {
	return c.Render("pages/disclaimer", h.base("disclaimer", fiber.Map{
		"Title":       "Disclaimer — PT Logam Gold Mulia Tbk",
		"Description": "Disclaimer informasi korporasi PT Logam Gold Mulia Tbk.",
		"Canonical":   h.cfg.BaseURL + "/disclaimer",
	}))
}


