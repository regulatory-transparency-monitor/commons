package models

import "time"

type RawData map[string][]interface{}

type GraphMetadata struct {
	Source    string
	Version   string
	Timestamp time.Time
}

type CombinedResources struct {
	Source string
	Data   []ServiceData
}

type ServiceData struct {
	ServiceSource string
	Data          []interface{}
}
