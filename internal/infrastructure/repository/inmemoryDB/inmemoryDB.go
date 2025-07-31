package inmemorydb

type InmemoryDB struct {
	AccountRepo *AccountRepo
	ContactRepo *ContactRepo
}

func NewInmomryDB() *InmemoryDB {
	return &InmemoryDB{
		AccountRepo: NewAccountRepo(),
		ContactRepo: NewContactRepo(),
	}
}
