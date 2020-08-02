package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {

	// Initialization

	// Execution
	user, err := GetUser(0)

	// Validation
	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err, "We were expecting and error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode, "We were expecting error 404")
	assert.EqualValues(t, "not_found", err.Code, "We were expecting not_found err.Code")
	assert.EqualValues(t, "User 0 was not found", err.Message, "We were expecting an empty err.Message")
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err, "We expected nil err variable")
	assert.NotNil(t, user, "We expected a user object")
	assert.EqualValues(t, user.Id, 123, "We expected user.ID to equal 123")
	assert.EqualValues(t, user.FirstName, "Alex", "We expected user.Firstname to be Alex")
}
