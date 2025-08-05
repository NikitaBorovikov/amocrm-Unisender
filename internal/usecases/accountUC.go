package usecases

import (
	"context"
	"fmt"

	"amocrm2.0/internal/core/amocrm"
	pb "amocrm2.0/proto"
)

type AccountUC struct {
	pb.AccountServiceServer
	AccountRepo amocrm.AccountRepo
}

func NewAccountUC(accountRepo amocrm.AccountRepo) *AccountUC {
	return &AccountUC{
		AccountRepo: accountRepo,
	}
}

func (uc *AccountUC) Add(account *amocrm.Account) error {
	return uc.AccountRepo.Add(account)
}

func (uc *AccountUC) GetByID(accountID int) (*amocrm.Account, error) {
	account, err := uc.AccountRepo.GetByID(accountID)
	if err != nil || account == nil {
		return nil, err
	}
	return account, nil
}

func (uc *AccountUC) GetAll() ([]amocrm.Account, error) {
	accounts, err := uc.AccountRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (uc *AccountUC) Update(account *amocrm.Account) error {
	return uc.AccountRepo.Update(account)
}

func (uc *AccountUC) Delete(accountID int) error {
	return uc.AccountRepo.Delete(accountID)
}

func (uc *AccountUC) UpdateUnisenderKey(accountID int, key string) error {
	if !validateUnisenderKey(key) {
		return fmt.Errorf("invalid unisender API key")
	}
	return uc.AccountRepo.UpdateUnisenderKey(accountID, key)
}

func (uc *AccountUC) GetUnisenderKey(accountID int) (string, error) {
	key, err := uc.AccountRepo.GetUnisenderKey(accountID)
	if len(key) == 0 {
		return "", fmt.Errorf("key is ivalid")
	}
	return key, err
}

func (uc *AccountUC) Unsubscribe(ctx context.Context, req *pb.UnsubscribeRequest) (*pb.UnsubscribeResponse, error) {
	resp := &pb.UnsubscribeResponse{}
	if err := uc.AccountRepo.Delete(int(req.AccountId)); err != nil {
		resp.Message = err.Error()
		resp.Success = false
		return resp, err
	}
	resp.Message = fmt.Sprintf("Account was successfully deleted (accountID = %d)", req.AccountId)
	resp.Success = true
	return resp, nil
}

func validateUnisenderKey(key string) bool {
	if len(key) < 15 || len(key) > 100 {
		return false
	}
	return true
}
