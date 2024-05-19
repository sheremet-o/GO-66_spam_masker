package masker_test

import (
	"masker"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProducer struct {
	mock.Mock
}

func (mp *MockProducer) Produce() ([]string, error) {
	args := mp.Called()
	return args.Get(0).([]string), args.Error(1)
}

type MockPresenter struct {
	mock.Mock
}

func (mp *MockPresenter) Present(data []string) error {
	args := mp.Called(data)
	return args.Error(0)
}

func TestMaskingService_Run(t *testing.T) {
	mockProducer := new(MockProducer)
	mockProducer.On("Produce").Return([]string{"http://example.com", "This is a message"}, nil)

	mockPresenter := new(MockPresenter)
	mockPresenter.On("Present", []string{"****example.com", "This is a message"}).Return(nil)

	ms := masker.NewMaskingService(mockProducer, mockPresenter)
	ms.Run()

	mockProducer.AssertExpectations(t)
	mockPresenter.AssertExpectations(t)
}

func TestMaskingService_Masker(t *testing.T) {
	ms := masker.MaskingService{}

	input := "http://example.com is a link"
	expectedOutput := "****example.com is a link"

	assert.Equal(t, expectedOutput, ms.Masker(input))
}
