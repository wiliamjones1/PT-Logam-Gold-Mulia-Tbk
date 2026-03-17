package company

// Data holds all configurable corporate information.
// Fields left empty will trigger graceful fallback text in templates.
type Data struct {
	// Legal identity
	LegalName    string
	ShortName    string
	Ticker       string
	ListingBoard string
	Status       string
	Sector       string
	SubSector    string
	YearEstd     string
	NIB          string
	NPWP         string

	// Contact
	Email             string
	InvestorEmail     string
	CorpSecEmail      string
	Phone             string
	Website           string
	WhatsApp          string
	BusinessHoursText string

	// Addresses
	HeadOffice string
	OpsOffice  string
	MapEmbedURL string

	// People
	Directors    []Person
	Commissioners []Person
	AuditCommittee []Person
	CorpSecretary string

	// External parties
	ShareRegistrar string
	ExternalAuditor string

	// Documents
	Documents []Document

	// Public disclosures
	Disclosures []Disclosure

	// External links (IDX, etc.)
	ExternalLinks []ExternalLink

	// News / articles
	News []NewsItem

	// FAQ
	FAQs []FAQ

	// Company overview text blocks
	OverviewShort  string
	OverviewLong   string
	VisionText     string
	MissionItems   []string
	Values         []Value
	Milestones     []Milestone

	// Business lines
	BusinessLines []BusinessLine
}

// Person represents a board member or committee member.
type Person struct {
	Name     string
	Position string
	Brief    string
}

// Document represents a downloadable corporate document.
type Document struct {
	Title       string
	Category    string // "annual-report", "financial", "governance", "prospectus", "other"
	Period      string
	Description string
	URL         string
}

// Disclosure represents a public information announcement.
type Disclosure struct {
	Title string
	Date  string
	Type  string
	URL   string
}

// ExternalLink represents a link to official external resources.
type ExternalLink struct {
	Title       string
	Description string
	URL         string
}

// NewsItem represents a news article or press release.
type NewsItem struct {
	Title   string
	Summary string
	Date    string
	URL     string
}

// FAQ represents a frequently asked question.
type FAQ struct {
	Question string
	Answer   string
}

// Value represents a corporate value.
type Value struct {
	Title       string
	Description string
	Icon        string // SVG path content
}

// Milestone represents a company milestone.
type Milestone struct {
	Year  string
	Title string
	Desc  string
}

// BusinessLine represents a line of business.
type BusinessLine struct {
	Title       string
	Description string
	Features    []string
	Icon        string
}

// Default returns the corporate data with current known information.
// Empty strings signal "not yet available" and templates will show fallback text.
func Default() *Data {
	return &Data{
		LegalName:    "PT Logam Gold Mulia Tbk",
		ShortName:    "Logam Gold",
		Ticker:       "",
		ListingBoard: "IDX",
		Status:       "Perseroan Terbuka (Tbk)",
		Sector:       "Logam Mulia",
		SubSector:    "Perdagangan & Distribusi Logam Mulia",
		YearEstd:     "",
		NIB:          "",
		NPWP:         "",

		Email:             "info@logam.gold",
		InvestorEmail:     "ir@logam.gold",
		CorpSecEmail:      "corsec@logam.gold",
		Phone:             "",
		Website:           "logam.gold",
		WhatsApp:          "",
		BusinessHoursText: "Senin — Jumat, 09.00 — 17.00 WIB",

		HeadOffice:  "",
		OpsOffice:   "",
		MapEmbedURL: "",

		Directors:       nil,
		Commissioners:   nil,
		AuditCommittee:  nil,
		CorpSecretary:   "",
		ShareRegistrar:  "",
		ExternalAuditor: "",

		Documents:     nil,
		Disclosures:   nil,
		ExternalLinks: nil,
		News:          nil,

		OverviewShort: "PT Logam Gold Mulia Tbk adalah perusahaan terbuka yang bergerak di bidang industri logam mulia. Perseroan berkomitmen untuk menjalankan kegiatan usaha dengan mengedepankan prinsip tata kelola perusahaan yang baik, transparansi, serta integritas dalam setiap aspek operasional.",
		OverviewLong:  "Sebagai entitas yang tercatat di Bursa Efek Indonesia, PT Logam Gold Mulia Tbk memposisikan diri sebagai pelaku usaha yang bertanggung jawab di sektor logam mulia. Perseroan menjalankan aktivitas bisnis dengan berlandaskan pada kepatuhan terhadap peraturan yang berlaku, prinsip kehati-hatian, serta komitmen terhadap kepentingan seluruh pemangku kepentingan. Informasi korporasi yang lebih detail akan disampaikan melalui kanal resmi Perseroan sesuai dengan ketentuan keterbukaan informasi yang berlaku.",

		VisionText: "Menjadi perusahaan terdepan dan terpercaya di industri logam mulia Indonesia yang memberikan nilai tambah berkelanjutan bagi seluruh pemangku kepentingan.",
		MissionItems: []string{
			"Menjalankan kegiatan usaha di bidang logam mulia dengan standar profesionalisme dan integritas tertinggi.",
			"Menerapkan prinsip tata kelola perusahaan yang baik secara konsisten dan transparan.",
			"Membangun kepercayaan jangka panjang dengan seluruh mitra bisnis, investor, dan pemangku kepentingan.",
			"Berkontribusi positif terhadap perkembangan industri logam mulia nasional.",
		},

		Values: []Value{
			{Title: "Integritas", Description: "Menjunjung tinggi kejujuran, etika, dan kepatuhan dalam setiap aspek operasional perusahaan.", Icon: "M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"},
			{Title: "Transparansi", Description: "Berkomitmen pada keterbukaan informasi dan komunikasi yang jelas kepada seluruh pemangku kepentingan.", Icon: "M15 12a3 3 0 11-6 0 3 3 0 016 0z M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"},
			{Title: "Profesionalisme", Description: "Menjalankan setiap kegiatan usaha dengan standar profesional dan kompetensi yang tinggi.", Icon: "M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z"},
			{Title: "Akuntabilitas", Description: "Bertanggung jawab penuh atas setiap keputusan dan tindakan perusahaan kepada seluruh pihak terkait.", Icon: "M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4"},
			{Title: "Kehati-hatian", Description: "Menerapkan prinsip prudensial dalam setiap pengambilan keputusan strategis perusahaan.", Icon: "M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"},
			{Title: "Keberlanjutan", Description: "Mengutamakan pertumbuhan usaha yang berkelanjutan dan bertanggung jawab dalam jangka panjang.", Icon: "M3.055 11H5a2 2 0 012 2v1a2 2 0 002 2 2 2 0 012 2v2.945M8 3.935V5.5A2.5 2.5 0 0010.5 8h.5a2 2 0 012 2 2 2 0 104 0 2 2 0 012-2h1.064M15 20.488V18a2 2 0 012-2h3.064M21 12a9 9 0 11-18 0 9 9 0 0118 0z"},
		},

		Milestones: nil,

		BusinessLines: []BusinessLine{
			{
				Title:       "Perdagangan Logam Mulia",
				Description: "Kegiatan usaha di bidang perdagangan logam mulia dengan standar kualitas dan kepatuhan terhadap regulasi yang berlaku.",
				Features:    []string{"Perdagangan emas dan logam mulia", "Penjaminan kualitas produk", "Kepatuhan regulasi perdagangan"},
				Icon:        "M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4",
			},
			{
				Title:       "Distribusi & Rantai Pasok",
				Description: "Pengelolaan rantai distribusi logam mulia yang aman, terukur, dan dapat dipertanggungjawabkan.",
				Features:    []string{"Jaringan distribusi terintegrasi", "Sistem keamanan berlapis", "Pelacakan dan audit rantai pasok"},
				Icon:        "M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4",
			},
			{
				Title:       "Kemitraan Strategis",
				Description: "Pengembangan kerja sama strategis dengan berbagai pihak di ekosistem industri logam mulia.",
				Features:    []string{"Kolaborasi lintas sektor", "Pengembangan pasar bersama", "Transfer pengetahuan industri"},
				Icon:        "M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z",
			},
		},

		FAQs: []FAQ{
			{
				Question: "Apa itu PT Logam Gold Mulia Tbk?",
				Answer:   "PT Logam Gold Mulia Tbk adalah perusahaan terbuka (Tbk) yang bergerak di bidang industri logam mulia. Sebagai perusahaan tercatat di Bursa Efek Indonesia, Perseroan menjalankan kegiatan usaha dengan mengedepankan prinsip tata kelola perusahaan yang baik dan keterbukaan informasi.",
			},
			{
				Question: "Apakah PT Logam Gold Mulia Tbk merupakan perusahaan terdaftar di Bursa Efek Indonesia?",
				Answer:   "Ya, PT Logam Gold Mulia Tbk berstatus sebagai Perseroan Terbuka (Tbk) yang tercatat di Bursa Efek Indonesia (IDX). Informasi lebih lanjut mengenai profil emiten dapat diakses melalui situs resmi Bursa Efek Indonesia.",
			},
			{
				Question: "Bagaimana cara menghubungi bagian Investor Relations?",
				Answer:   "Untuk pertanyaan terkait hubungan investor, Anda dapat mengirimkan email ke ir@logam.gold atau menghubungi Sekretaris Perusahaan melalui corsec@logam.gold. Tim kami akan merespons sesuai dengan jam operasional kantor.",
			},
			{
				Question: "Di mana saya dapat mengakses laporan keuangan dan laporan tahunan perusahaan?",
				Answer:   "Laporan keuangan dan laporan tahunan akan dipublikasikan melalui situs resmi Perseroan di halaman Dokumen & Laporan, serta melalui kanal keterbukaan informasi resmi sesuai ketentuan yang berlaku.",
			},
			{
				Question: "Bagaimana cara mengajukan kerja sama bisnis dengan PT Logam Gold Mulia Tbk?",
				Answer:   "Untuk mengajukan proposal kerja sama bisnis, silakan menghubungi kami melalui halaman Kontak di situs ini atau mengirimkan email ke info@logam.gold dengan menyertakan informasi perusahaan dan proposal singkat kerja sama yang diinginkan.",
			},
			{
				Question: "Apa saja prinsip tata kelola perusahaan yang diterapkan?",
				Answer:   "Perseroan menerapkan prinsip-prinsip Good Corporate Governance (GCG) yang meliputi Transparansi, Akuntabilitas, Responsibilitas, Independensi, dan Kewajaran (TARIF). Penerapan prinsip-prinsip ini dijabarkan lebih lanjut di halaman Tata Kelola Perusahaan.",
			},
			{
				Question: "Apakah informasi di situs ini merupakan informasi resmi perusahaan?",
				Answer:   "Informasi yang disajikan di situs logam.gold bersifat umum dan informatif. Untuk informasi material dan keputusan investasi, kami menyarankan untuk merujuk pada keterbukaan informasi resmi yang dipublikasikan melalui kanal-kanal yang ditentukan oleh otoritas pasar modal.",
			},
		},
	}
}
