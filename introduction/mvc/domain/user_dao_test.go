package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {

	user, err := UserDao.GetUser(0)

	assert.Nil(t, user, "we are not accepting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)

	//if user != nil {
	//	t.Error("we are not excepting a user with id 0")
	//}
	//if err == nil {
	//	t.Error("we are excepting a error when user id is 0")
	//}
	//if err.StatusCode != http.StatusNotFound {
	//	t.Error("We are excepting 404 when user is not found")
	//}

}

func TestGetUserNoError(t *testing.T) {
	// Initialization
	user, err := UserDao.GetUser(123)

	UserDao.GetUser(123)


	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, user, "we are excepting response")
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Ravi", user.FirstName)
	assert.EqualValues(t, "Shankar", user.LastName)
	assert.EqualValues(t, "rprajapati067@gmail.com", user.Email)

}
