package travis

// requestBody is the top level object received from a Travis CI webhook
type requestBody struct {
	Payload *Payload `json:"payload"`
}

// Payload contains the actual data from Travis
type Payload struct {
	ID             string `json:"id"`
	Number         string `json:"number"`
	Status         string `json:"status"`
	StartedAt      string `json:"started_at"`
	FinishedAt     string `json:"finished_at"`
	StatusMessage  string `json:"status_message"`
	Commit         string `json:"commit"`
	Branch         string `json:"branch"`
	Message        string `json:"message"`
	CompareURL     string `json:"compare_url"`
	CommittedAt    string `json:"committed_at"`
	CommitterName  string `json:"committer_name"`
	CommitterEmail string `json:"committer_email"`
	AuthorName     string `json:"author_name"`
	AuthorEmail    string `json:"author_email"`
	Type           string `json:"type"`
	BuildURL       string `json:"build_url"`

	Repository *Repository      `json:"repository"`
	Config     *Config          `json:"config"`
	Matrix     []*MatrixElement `json:"matrix"`
}

// Repository contains the repo data
type Repository struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	OwnerName string `json:"owner_name"`
	URL       string `json:"url"`
}

// Config contains the notification data
type Config struct {
	Notifications *Notifications `json:"notifications"`
}

// Notifications contains a list of webhooks
type Notifications struct {
	WebHooks []string `json:"webhooks"`
}

// MatrixElement contains the data for each Travis Matrix object
type MatrixElement struct {
	ID             string `json:"id"`
	RepositoryID   string `json:"repository_id"`
	Number         string `json:"number"`
	State          string `json:"state"`
	StartedAt      string `json:"started_at"`
	FinishedAt     string `json:"finished_at"`
	Status         string `json:"status"`
	Log            string `json:"log"`
	Result         string `json:"result"`
	ParentID       string `json:"parent_id"`
	Commit         string `json:"commit"`
	Branch         string `json:"branch"`
	Message        string `json:"message"`
	CommittedAt    string `json:"committed_at"`
	CommitterName  string `json:"committer_name"`
	CommitterEmail string `json:"committer_email"`
	AuthorName     string `json:"author_name"`
	AuthorEmail    string `json:"author_email"`
	CompareURL     string `json:"compare_url"`

	Config Config `json:"config"`
}
