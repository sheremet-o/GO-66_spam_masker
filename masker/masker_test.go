package masker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskingService_Masker(t *testing.T) {
	service := NewMaskingService(nil, nil)

	assert.Equal(t, "http://***********", service.Masker("http://example.com"))

	assert.Equal(t, "www.example.com", service.Masker("www.example.com"))

	assert.Equal(t, " ", service.Masker(" "))

	assert.Equal(t, "http://*********** and http://********", service.Masker("http://example.com and http://test.com"))
}
