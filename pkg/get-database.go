package pkg

import (
	"fmt"
	"github.com/bytedance/sonic/decoder"
	"github.com/charmbracelet/log"
	"io"
	"net/http"
)

func GetDatabase(organizationName string, organizationToken string, databaseName string) (Db, error) {
	url := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases/%s", organizationName, databaseName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("DBPU: GetDatabase: Error creating request. %v", err)
		return Db{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", organizationToken))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("DBPU: GetDatabase: Error sending request. %v", err)
		return Db{}, fmt.Errorf("Error reading response. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("DBPU: GetDatabase: Error reading body. %v", err)
		return Db{}, fmt.Errorf("Error reading body. %v", err)
	}
	var db Db
	err = decoder.NewDecoder(string(body)).Decode(&db)
	if err != nil {
		log.Errorf("DBPU: GetDatabase: Error decoding body. %v", err)
		return Db{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return db, nil
}
