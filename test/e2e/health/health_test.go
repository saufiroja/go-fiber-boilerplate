package health_test

import (
	"net/http"
	"project/go-fiber-boilerplate/utils"
	"project/go-fiber-boilerplate/utils/constants"
	"testing"

	"github.com/stretchr/testify/suite"
)

type HealthSuiteTest struct {
	suite.Suite
	client utils.ISuite
	res    *constants.Response
}

func NewHealthSuiteTest() *HealthSuiteTest {
	return &HealthSuiteTest{
		client: utils.NewSuiteUtils(&http.Client{}),
		res:    &constants.Response{},
	}
}

func (suite *HealthSuiteTest) TestRegister() {
	res, resp, err := suite.client.Get("http://localhost:3000/health", suite.res)

	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)

	suite.Equal("welcome to go fiber boilerplate", res.Message)
}

func TestHealthSuiteTest(t *testing.T) {
	suite.Run(t, NewHealthSuiteTest())
}
