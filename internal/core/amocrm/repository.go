package amocrm

type AccountRepo interface {
	Add(account *Account) error
	GetByID(accountID int) (*Account, error)
	GetAll() ([]Account, error)
	Update(account *Account) error
	Delete(accountID int) error
	UpdateUnisenderKey(accountID int, key string) error
}

type IntegrationRepo interface {
	Add(integration *Integration) error
	GetByID(integrationID int) (*Integration, error)
	GetAll() ([]Integration, error)
	Update(integration *Integration) error
	Delete(integrationID int) error
}

type ContactRepo interface {
	Add(contact *Contact) error
	GetByID(contactID int) (*Contact, error)
	GetAll() ([]Contact, error)
	Update(contact *Contact) error
	Delete(contactID int) error
}
