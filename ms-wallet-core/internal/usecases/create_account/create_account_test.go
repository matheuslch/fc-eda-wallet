package create_account

import (
	"testing"

	"github.com/matheuslch/fc-ms-wallet-core/internal/entity"
	"github.com/matheuslch/fc-ms-wallet-core/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "john@doe.com")
	clientMock := &mocks.ClientGatewayMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	useCase := NewCreateAccountUseCase(accountMock, clientMock)
	output, err := useCase.Execute(CreateAccountInputDTO{
		ClientID: client.ID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	clientMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}
