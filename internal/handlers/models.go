package handlers

type Query struct {
	Id                string                 `json:"id" binding:"required"`
	Compliance        LocalCompliance        `json:"compliance"`
	ControlSet        []string               `json:"controlSet"`
	Fveys             Fveys                  `json:"fveys"`
	Datasource        string                 `json:"datasource"`
	DatasourceAdapter string                 `json:"datasourceAdapter"`
	Criteria          string                 `json:"criteria"`
	StartDate         string                 `json:"startDate"`
	EndDate           string                 `json:"endDate"`
	Params            map[string]interface{} `json:"params"`
	Name              string                 `json:"name"`
	MaxResults        int                    `json:"maxResults"`
	Submitted         string                 `json:"submitted"`
	User              User                   `json:"user"`
	System            string                 `json:"system"`
	Edh               Edh                    `json:"edh"`
	Authorization     string                 `json:"authorization"`
}

type LocalCompliance struct {
	Op                       string                   `json:"op"`
	NecessityProportionality NecessityProportionality `json:"necessityProportionality"`
	Location                 string                   `json:"location"`
	Nationality              string                   `json:"nationality"`
	Authorisation            Authorisation            `json:"authorisation"`
	Authorisations           []string                 `json:"authorisations"`
	Id                       string                   `json:"id"`
	TemplateId               string                   `json:"templateId"`
	ShortTitle               string                   `json:"shortTitle"`
	LastUsed                 string                   `json:"lastUsed"`
	ReviewDate               string                   `json:"reviewDate"`
	Edh                      Edh                      `json:"edh"`
}

type NecessityProportionality struct {
	Statement      string            `json:"statement"`
	Segments       map[string]string `json:"segments"`
	IntrusionLevel string            `json:"intrusionLevel"`
}

type Authorisation struct {
	Authorisations []string `json:"authorisations"`
}

type Fveys struct {
	Compliance map[string]interface{} `json:"compliance"`
	Ehd        Edh                    `json:"ehd"`
}

type User struct {
	Identity     string   `json:"identity"`
	Location     string   `json:"location"`
	Organisation string   `json:"organisation"`
	Roles        []string `json:"roles"`
}

type Edh struct {
	EdhSpecification     string   `json:"edhSpecification"`
	EdhVersion           string   `json:"edhVersion"`
	EdhIdentifier        string   `json:"edhIdentifier"`
	EdhCreateDateTime    string   `json:"edhCreateDateTime"`
	EdhResponsibleEntity []string `json:"edhResponsibleEntity"`
	EdhPolicyRef         string   `json:"edhPolicyRef"`
	EdhAuthRef           string   `json:"edhAuthRef"`
	EdhDataSet           string   `json:"edhDataSet"`
	EdhPolicy            string   `json:"edhPolicy"`
	EdhControlSet        []string `json:"edhControlSet"`
}
