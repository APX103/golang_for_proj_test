package cmd

// ----- Define type for cobra command line parser
type CobraParam struct {
	Key      string `json:"key" bson:"key"`
	Type     string `json:"type" bson:"type"`
	Default  string `json:"default,omitempty" bson:"default,omitempty"`
	Required bool   `json:"required" bson:"required"`
	Help     string `json:"help,omitempty" bson:"help,omitempty"`
}

type CobraCMD struct {
	CMD    string       `json:"cmd" bson:"cmd"`
	SubCMD []CobraCMD   `json:"sub-cmd,omitempty" bson:"cmd,omitempty"`
	Params []CobraParam `json:"params,omitempty" bson:"params,omitempty"`
	Short  string       `json:"short,omitempty" bson:"short,omitempty"`
	Long   string       `json:"long,omitempty" bson:"long,omitempty"`
}

// ----- End define

// +++++ Define type enumerate for task

type ParamEnum string

const String ParamEnum = "String"
const StringToString ParamEnum = "StringToString"

type ParamStruct struct {
	Type  ParamEnum
	Value any
}

type TaskCmd struct {
	SubCmd map[string]*TaskCmd
	Enable bool
	Params map[string]*ParamStruct
}

// +++++ End define
