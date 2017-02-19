package models

type Alarm struct {
	ID          string `json:"_id"`
	Enable      string `json:"enable"`
	Groups      string `json:"groups"`
	Alert       string `json:"alert"`
	Level       string `json:"level"`
	Version     string `json:"version"`
	DB          string `json:"db"`
	RP          string `json:"rp"`
	Measurement string `json:"measurement"`
	Name        string `json:"name"`
	Func        string `json:"func"`
	Value       string `json:"value"`
	Expression  string `json:"expression"`
	Period      string `json:"period"`
	Every       string `json:"every"`
	GroupBy     string `json:"groupby"`
	Where       string `json:"where"`
	Message     string `json:"message"`
	Trigger     string `json:"trigger"`
	Shift       string `json:"shift"`
	MD5         string `json:"md5"`
	Default     string `json:"default"`
}

var (
	VersionSep = "__"
	DBPrefix   = "collect."

	ThresHold = "threshold"
	Relative  = "relative"
	DeadMan   = "deadman"
)
