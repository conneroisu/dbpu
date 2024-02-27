package pkg

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bytedance/sonic/decoder"
	"github.com/charmbracelet/log"
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
	var db Db
	err = decoder.NewDecoder(string(body)).Decode(&db)
	if err != nil {
		log.Errorf("DBPU: CreateDatabase: Error decoding body. %v", err)
		return fmt.Errorf("Error decoding body. %v", err)
	}
	return nil
}

func CreateDatabaseToken(organizationName string, databaseName string, token string, expiration string, authorization string) (Jwt, error) {
	url := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases/%s/auth/tokens?expiration=%s&authorization=%s", organizationName, databaseName, expiration, authorization)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Errorf("DBPU: CreateDatabaseToken: failed to create database token: %v", err)
		return Jwt{}, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("DBPU: CreateDatabaseToken: failed to create database token: %v", err)
		return Jwt{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Errorf("DBPU: CreateDatabaseToken: failed to create database token: %s", resp.Status)
		return Jwt{}, fmt.Errorf("failed to create database token: %s", resp.Status)
	}
	var jwt Jwt
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("DBPU: CreateDatabaseToken: failed to read response body: %v", err)
		return Jwt{}, err
	}
	err = decoder.NewDecoder(string(body)).Decode(&jwt)
	if err != nil {
		log.Errorf("DBPU: CreateDatabaseToken: failed to decode response body: %v", err)
		return Jwt{}, err
	}
	return jwt, nil
}

func DeleteDatabaseTokens(organizationName string, databaseName string, token string) error {
	url := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases/%s/auth/rotate", organizationName, databaseName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Errorf("DBPU: DeleteDatabaseTokens: failed to delete database tokens: %v", err)
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("DBPU: DeleteDatabaseTokens: failed to delete database tokens: %v", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Errorf("DBPU: DeleteDatabaseTokens: failed to delete database tokens: %s", resp.Status)
		return fmt.Errorf("failed to delete database tokens: %s", resp.Status)
	}
	return nil
}


func DeleteDatabase(organizationName string, organizationToken string, databaseName string) (string, error) {
	url := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases/%s", organizationName, databaseName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Errorf("DBPU: DeleteDatabase: Error creating request. %v", err)
		return "", fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", organizationToken))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("DBPU: DeleteDatabase: Error sending request. %v", err)
		return "", fmt.Errorf("Error reading response. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("DBPU: DeleteDatabase: Error reading body. %v", err)
		return "", fmt.Errorf("Error reading body. %v", err)
	}
	var db Db
	err = decoder.NewDecoder(string(body)).Decode(&db)
	if err != nil {
		log.Errorf("DBPU: DeleteDatabase: Error decoding body. %v", err)
		return "", fmt.Errorf("Error decoding body. %v", err)
	}
	return db.DbName, nil
}

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

func GetDatabaseStats(organizationName string, organizationToken string, databaseName string) (Db, error) {
	url := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases/%s/stats", organizationName, databaseName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("DBPU: GetDatabaseStats: Error creating request. %v", err)
		return Db{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", organizationToken))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("DBPU: GetDatabaseStats: Error sending request. %v", err)
		return Db{}, fmt.Errorf("Error reading response. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("DBPU: GetDatabaseStats: Error reading body. %v", err)
		return Db{}, fmt.Errorf("Error reading body. %v", err)
	}
	var db Db
	err = decoder.NewDecoder(string(body)).Decode(&db)
	if err != nil {
		log.Errorf("DBPU: GetDatabaseStats: Error decoding body. %v", err)
		return Db{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return db, nil
}

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

func GetDatabaseUsage(organizationName string, organizationToken string, databaseName string, from string, to string) (Db, error) {
	url := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases/%s/usage?from=%s&to=%s", organizationName, databaseName, from, to)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("DBPU: GetDatabaseDb: Error creating request. %v", err)
		return Db{}, fmt.Errorf("Error reading request. %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", organizationToken))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("DBPU: GetDatabaseDb: Error sending request. %v", err)
		return Db{}, fmt.Errorf("Error reading response. %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("DBPU: GetDatabaseDb: Error reading body. %v", err)
		return Db{}, fmt.Errorf("Error reading body. %v", err)
	}
	var usage Db
	err = decoder.NewDecoder(string(body)).Decode(string(body))
	if err != nil {
		log.Errorf("DBPU: GetDatabaseDb: Error decoding body. %v", err)
		return Db{}, fmt.Errorf("Error decoding body. %v", err)
	}
	return usage, nil
}

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

func UploadDump(organizationName string, databaseName string, token string, dump io.Reader) error {
	url := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases/%s/dump", organizationName, databaseName)
	req, err := http.NewRequest("POST", url, dump)
	if err != nil {
		log.Errorf("DBPU: UploadDump: failed to upload dump: %v", err)
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("DBPU: UploadDump: failed to upload dump: %v", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Errorf("DBPU: UploadDump: failed to upload dump: %s", resp.Status)
		return fmt.Errorf("failed to upload dump: %s", resp.Status)
	}
	return nil
}
