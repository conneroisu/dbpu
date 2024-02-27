package pkg

import (
	"fmt"
	"github.com/bytedance/sonic/decoder"
	"github.com/charmbracelet/log"
	"io"
	"net/http"
)

func ListDatabases(organizationName string, organizationToken string) (Db, error) {
	url := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases", organizationName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("DBPU: ListDatabases: Error creating request. %v", err)
		return Db{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", organizationToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("DBPU: ListDatabases: Error sending request. %v", err)
		return Db{}, fmt.Errorf("Error reading response. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("DBPU: ListDatabases: Error reading body. %v", err)
		return Db{}, fmt.Errorf("Error reading body. %v", err)
	}
	var db Db
	err = decoder.NewDecoder(string(body)).Decode(&db)
	if err != nil {
		log.Errorf("DBPU: ListDatabases: Error decoding body. %v", err)
		return Db{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return db, nil
}
