package queue

import "amocrm2.0/internal/core/amocrm"

type SyncContactsTask struct {
	AccountID int              `json:"account_id"`
	TaskType  string           `json:"task_type"`  //firstSync or webhookSync
	EventType string           `json:"event_type"` //add or update
	Contacts  []amocrm.Contact `json:"contacts"`
}
