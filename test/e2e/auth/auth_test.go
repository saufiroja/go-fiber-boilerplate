package auth_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"project/go-fiber-boilerplate/models/dto"
	"project/go-fiber-boilerplate/utils"
	"project/go-fiber-boilerplate/utils/constants"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	suite.Suite
	client utils.ISuite
	res    *constants.Response
}

func NewAuthTestSuite() *AuthTestSuite {
	return &AuthTestSuite{
		client: utils.NewSuiteUtils(&http.Client{}),
		res:    &constants.Response{},
	}
}

func (suite *AuthTestSuite) TestRegister() {
	input := dto.Register{
		FullName:  "test",
		Email:     "test@gmail.com",
		Password:  "12345678",
		IsMale:    true,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	jsonBytes, err := json.Marshal(input)
	if err != nil {
		log.Fatal(err)
	}

	byt := bytes.NewBuffer(jsonBytes)

	res, resp, err := suite.client.Post("http://localhost:3000/auth/register", byt, suite.res)
	suite.NoError(err)
	suite.Equal(http.StatusCreated, resp.StatusCode)

	suite.Equal("register success", res.Message)
}

func (suite *AuthTestSuite) TestLogin() {
	input := dto.Register{
		FullName:  "halo",
		Email:     "halo@gmail.com",
		Password:  "12345678",
		IsMale:    true,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	jsonBytes, err := json.Marshal(input)
	if err != nil {
		log.Fatal(err)
	}

	byt := bytes.NewBuffer(jsonBytes)

	res, resp, err := suite.client.Post("http://localhost:3000/auth/register", byt, suite.res)
	suite.NoError(err)
	suite.Equal(http.StatusCreated, resp.StatusCode)

	suite.Equal("register success", res.Message)

	time.Sleep(1 * time.Second)

	req := dto.Login{
		Email:    "halo@gmail.com",
		Password: "12345678",
	}

	jsonBytes2, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	reqBody2 := bytes.NewBuffer(jsonBytes2)

	res2, resp2, err := suite.client.Post("http://localhost:3000/auth/login", reqBody2, suite.res)

	suite.NoError(err)
	suite.Equal(http.StatusOK, resp2.StatusCode)

	suite.Equal("login success", res2.Message)
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, NewAuthTestSuite())
}
