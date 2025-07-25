package amocrm

type Contact struct {
	ContactID  int
	AccountID  int
	Name       string
	Email      string
	SyncStatus bool
}
