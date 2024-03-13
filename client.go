package dbpu

type DbundToken struct {
	Database Db
	Token    string
}

type Client struct {
	ApiToken  string
	OrgToken  string
	Org       Org
	Databases []DbundToken
}
