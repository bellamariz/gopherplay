package ingests

type Ingest struct {
	Signal       string   `json:"signal"`
	Packagers    []string `json:"packagers"`
	LastReported string   `json:"last_reported"`
}

func GetIngests() ([]Ingest, error) {
	return []Ingest{}, nil
}

func UpdateIngest(ing Ingest) error {
	return nil
}
