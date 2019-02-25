package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Handle(req []byte) bool {

	var statusPageNotification *StatusPageNotification
	err := json.Unmarshal(req, &statusPageNotification)
	if err != nil {
		panic(err)
	}

	attachment := []MMAttachment{}

	component := MMAttachment{}
	if statusPageNotification.Component != nil {
		msg := fmt.Sprintf("Status: %s\nDescription: %s", statusPageNotification.Component.Status, statusPageNotification.Component.Description)
		component = *component.AddField(MMField{Title: statusPageNotification.Component.Name, Value: msg})

		if statusPageNotification.ComponentUpdate != nil {
			msg := fmt.Sprintf("Old Status: %s\nNew Status: %s", statusPageNotification.ComponentUpdate.OldStatus, statusPageNotification.ComponentUpdate.NewStatus)
			component = *component.AddField(MMField{Value: msg, Short: true})
		}
	}

	if statusPageNotification.Incident != nil {
		component = *component.AddField(MMField{Title: statusPageNotification.Incident.Name, Value: statusPageNotification.Incident.Status})
		component = *component.AddField(MMField{Title: "Impact", Value: statusPageNotification.Incident.Impact, Short: true})
		component = *component.AddField(MMField{Title: "Link", Value: statusPageNotification.Incident.Shortlink, Short: true})
		createdAt := statusPageNotification.Incident.CreatedAt
		updatedAt := statusPageNotification.Incident.UpdatedAt
		component = *component.AddField(MMField{Title: "Created At", Value: createdAt.String(), Short: true})
		component = *component.AddField(MMField{Title: "Updated At", Value: updatedAt.String(), Short: true})

		for _, incidentUpdate := range statusPageNotification.Incident.IncidentUpdates {
			msg := fmt.Sprintf("Status: %s\nDescription: %s\nUpdatedAt: %s", incidentUpdate.Status, incidentUpdate.Body, incidentUpdate.UpdatedAt.String())
			component = *component.AddField(MMField{Title: "Incident Update", Value: msg})
		}

	}

	attachment = append(attachment, component)
	payload := MMSlashResponse{
		Username:    "StatusPageBot",
		IconUrl:     "https://pbs.twimg.com/profile_images/963832478728314880/QoqF8Db1_400x400.jpg",
		Attachments: attachment,
	}
	mmHook := os.Getenv("MATTERMOST_HOOK")
	log.Println(payload.ToJson())
	if mmHook != "" {
		send(mmHook, payload)
	}

	return true

}

func send(webhookURL string, payload MMSlashResponse) {
	marshalContent, _ := json.Marshal(payload)
	var jsonStr = []byte(marshalContent)
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "statusPageBot")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}
