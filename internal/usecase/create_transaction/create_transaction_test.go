package createtransaction

import (
	"testing"

	"github.com/mrpsousa/internal/entity"
	"github.com/mrpsousa/internal/event"
	"github.com/mrpsousa/mocks"
	"github.com/mrpsousa/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("client1", "eu@email.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("client2", "eu2@email.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockAccount := &mocks.AccountGateway{}
	mockAccount.On("FindById", account1.ID).Return(account1, nil)
	mockAccount.On("FindById", account2.ID).Return(account2, nil)

	mockTransaction := &mocks.TransactionGateway{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	dispatcher := *events.NewEventDispatcher()
	event := event.NewTransactionCreated()

	uc := NewCreateTransactionUseCase(mockTransaction, mockAccount, dispatcher, event)
	output, err := uc.Execute(&inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockAccount.AssertNumberOfCalls(t, "FindById", 2)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)
}
