package createaccount

import (
	"testing"

	"github.com/mrpsousa/internal/entity"
	"github.com/mrpsousa/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountSuccess(t *testing.T) {
	client, _ := entity.NewClient("roger", "eu@email.com")
	clientMock := &mocks.ClientGateway{}
	clientMock.On("Get", mock.Anything).Return(client, nil)

	accountMock := &mocks.AccountGateway{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountMock, clientMock)
	inputDto := &CreateAccountInputDTO{
		ClientID: client.ID,
	}
	output, err := uc.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)

}
