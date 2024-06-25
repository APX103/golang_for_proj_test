package cmd

type CobraParam struct {
	Key      string `json:"key" bson:"key"`
	Type     string `json:"type" bson:"type"`
	Default  string `json:"default,omitempty" bson:"default,omitempty"`
	Required bool   `json:"required" bson:"required"`
	Help     string `json:"help,omitempty" bson:"help,omitempty"`
}

// ,omitempty

type CobraCMD struct {
	CMD    string       `json:"cmd" bson:"cmd"`
	SubCMD []CobraCMD   `json:"sub-cmd,omitempty" bson:"cmd,omitempty"`
	Params []CobraParam `json:"params,omitempty" bson:"params,omitempty"`
	Short  string       `json:"short,omitempty" bson:"short,omitempty"`
	Long   string       `json:"long,omitempty" bson:"long,omitempty"`
}
