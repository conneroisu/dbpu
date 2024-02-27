
package pkg

import (
	"fmt"
	"github.com/charmbracelet/log"
	"net/http"
)

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
