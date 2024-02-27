package pkg

type Jwt struct {
	Jwt string `json:"jwt"` // jwt
}
type TopQueries struct {
	Query      string `json:"query"`       // query
	RowsRead   int    `json:"rows_read"`   // number of rows read by query
	RowsWriten int    `json:"rows_writen"` // number of rows writen by query
}

type DbUsage struct {
	RowsRead     int `json:"rows_read"`     // number of rows read
	RowsWriten   int `json:"rows_writen"`   // number of rows writen
	StorageBytes int `json:"storage_bytes"` // storage in bytes
}

type DbInstance struct {
	Usage    DbUsage `json:"usage"`    // db usage
	Uuid     string  `json:"uuid"`     // db instance uuid
	Hostname string  `json:"hostname"` // db instance hostname
	Name     string  `json:"name"`     // db instance name
	Region   string  `json:"region"`   // db instance region
	Type     string  `json:"type"`     // object type
}

type DbTotal struct {
	RowsRead     int `json:"rows_read"`     // number of rows read
	RowsWriten   int `json:"rows_writen"`   // number of rows writen
	StorageBytes int `json:"storage_bytes"` // storage in bytes
}

type Db struct {
	Id            string       `json:"DbId"`          // db id
	Hostname      string       `json:"DbHostname"`    // db hostname
	DbName        string       `json:"Name"`          // db name
	Regions       []string     `json:"Regions"`       // regions where db deployed
	PrimaryRegion string       `json:"primaryRegion"` // primary region
	Group         string       `json:"Group"`         // db owned by group
	Type          string       `json:"type"`          // object type
	Version       string       `json:"version"`       // libSQL version
	Instances     []DbInstance `json:"instances"`     // db instances
	Total         DbTotal      `json:"total"`         // db total usage
	Uuid          string       `json:"uuid"`          // db uuid
}
