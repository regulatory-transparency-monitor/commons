package models

/* type ServiceData interface {
	GetServiceSource() string
	GetData() []interface{}
}

type CombinedResources interface {
	GetSource() string
	GetData() []ServiceData
} */

type CombinedResources struct {
	Source string
	Data   []ServiceData
}

type ServiceData struct {
	ServiceSource string
	Data          []interface{}
}
