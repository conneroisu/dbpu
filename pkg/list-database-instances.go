package pkg

import (
	"fmt"
	"github.com/bytedance/sonic/decoder"
	"github.com/charmbracelet/log"
	"io"
	"net/http"
)

func ListDatabaseInstances(organizationName string, databaseName string, token string) (DbInstance, error) {
	url := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases/%s/instances", organizationName, databaseName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error("Error on creating a request")
		return DbInstance{}, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Error on sending a request")
		return DbInstance{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("Error on reading response body")
		return DbInstance{}, err
	}
	var dbInstance DbInstance
	err = decoder.NewDecoder(string(body)).Decode(&dbInstance)
	if err != nil {
		log.Error("Error on decoding response body")
		return DbInstance{}, err
	}
	return dbInstance, nil
}
