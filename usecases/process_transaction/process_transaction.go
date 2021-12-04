package process_transaction

import "github.com/Daniel-Vinicius/golang-live/entities"

type ProcessTransaction struct {
	Repository entities.TransactionRepository
}

func NewProcessTransaction(repository entities.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := entities.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount

	invalidTransaction := transaction.IsValid()
	if invalidTransaction != nil {
		return p.rejectTransaction(transaction, invalidTransaction)
	}

	return p.approveTransaction(transaction)
}

func (p *ProcessTransaction) approveTransaction(transaction *entities.Transaction) (TransactionDtoOutput, error) {
	errorTryingInsertIntoDatabase := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "approved", "")
	if errorTryingInsertIntoDatabase != nil {
		return TransactionDtoOutput{}, errorTryingInsertIntoDatabase
	}

	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       "approved",
		ErrorMessage: "",
	}

	return output, nil
}

func (p *ProcessTransaction) rejectTransaction(transaction *entities.Transaction, invalidTransaction error) (TransactionDtoOutput, error) {
	errorTryingInsertIntoDatabase := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "rejected", invalidTransaction.Error())
	if errorTryingInsertIntoDatabase != nil {
		return TransactionDtoOutput{}, errorTryingInsertIntoDatabase
	}

	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       "rejected",
		ErrorMessage: invalidTransaction.Error(),
	}

	return output, nil
}
