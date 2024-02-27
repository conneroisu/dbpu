package pkg

import (
	"fmt"
	"github.com/bytedance/sonic/decoder"
	"github.com/charmbracelet/log"
	"io"
	"net/http"
)

func CreateDatabase(organizationName string, organizationToken string, databaseName string, databaseGroup string) error {
	url := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases", organizationName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Errorf("DBPU: CreateDatabase: Error creating request. %v", err)
		return fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", organizationToken))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("DBPU: CreateDatabase: Error sending request. %v", err)
		return fmt.Errorf("Error reading response. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("DBPU: CreateDatabase: Error reading body. %v", err)
		return fmt.Errorf("Error reading body. %v", err)
	}
	decoder := decoder.NewDecoder(string(body))
	var db Db
	err = decoder.Decode(&db)
	if err != nil {
		log.Errorf("DBPU: CreateDatabase: Error decoding body. %v", err)
		return fmt.Errorf("Error decoding body. %v", err)
	}
	return nil
}
