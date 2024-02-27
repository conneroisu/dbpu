

package pkg

import (
	"fmt"
	"github.com/charmbracelet/log"
	"io"
	"net/http"
)

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
