package pkg

import (
	"fmt"
	"github.com/bytedance/sonic/decoder"
	"github.com/charmbracelet/log"
	"io"
	"net/http"
)

//	curl -L https://api.turso.tech/v1/organizations/{organizationName}/databases/{databaseName}/instances/{instanceName} \
//	  -H 'Authorization: Bearer TOKEN'
func GetDatabaseInstance(instanceName string, organizationName string, databaseName string, token string) (DbInstance, error) {
	url := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases/%s/instances/%s", organizationName, databaseName, instanceName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error("DBPU: GetDatabaseInstance: Error on creating a request")
		return DbInstance{}, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("DBPU: GetDatabaseInstance: Error on sending a request")
		return DbInstance{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Errorf("DBPU: GetDatabaseInstance: failed to get database instance: %s", resp.Status)
		return DbInstance{}, fmt.Errorf("failed to get database instance: %s", resp.Status)
	}
	var dbInstance DbInstance
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("DBPU: GetDatabaseInstance: Error on reading response body")
		return DbInstance{}, err
	}
	err = decoder.NewDecoder(string(body)).Decode(&dbInstance)
	if err != nil {
		log.Error("DBPU: GetDatabaseInstance: Error on decoding response body")
		return DbInstance{}, err
	}
	return dbInstance, nil
}
