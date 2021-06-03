package polo

import (
	"github.com/rprajapati0067/golang-microservices/introduction/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}

func TestPolo(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/polo", nil)
	response := httptest.NewRecorder()

	c := test_utils.GetMockContext(request, response)

	Polo(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())
}
