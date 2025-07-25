package inmemorydb

type InmemoryDB struct {
	AccountRepo     *AccountRepo
	IntegrationRepo *IntegrationRepo
	ContactRepo     *ContactRepo
}

func NewInmomryDB() *InmemoryDB {
	return &InmemoryDB{
		AccountRepo:     NewAccountRepo(),
		IntegrationRepo: NewIntegrationRepo(),
		ContactRepo:     NewContactRepo(),
	}
}
