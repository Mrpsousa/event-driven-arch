package createclient

import (
	"testing"

	"github.com/mrpsousa/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClientSuccess(t *testing.T) {
	gateWayMock := &mocks.ClientGateway{}
	gateWayMock.On("Save", mock.Anything).Return(nil)
	uc := NewCreateClientUseCase(gateWayMock)

	output, err := uc.Execute(CreateClientInputDTO{
		Name:  "name",
		Email: "email@email.com",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "name", output.Name)
	assert.Equal(t, "email@email.com", output.Email)
	gateWayMock.AssertNumberOfCalls(t, "Save", 1)
}
