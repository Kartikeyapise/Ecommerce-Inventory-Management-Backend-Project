package view

import (
	"github.com/kartikeya/product_catalog_DIY/src/main/view"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponseMessage(t *testing.T) {

	var message view.ResponseMessage = view.ResponseMessage{
		Message: "message",
	}

	assert.Equal(t, message.Message, "message")
}

