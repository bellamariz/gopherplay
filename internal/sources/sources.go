package sources

type Ingest struct {
	Signal       string   `json:"signal"`
	Packagers    []string `json:"packagers"`
	LastReported string   `json:"last_reported"`
}

type Source struct {
	Signal string `json:"signal"`
	Server string `json:"server"`
}
