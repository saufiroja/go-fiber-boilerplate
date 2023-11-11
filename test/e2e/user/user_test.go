package user_test

// import (
// 	"fmt"
// 	"net/http"
// 	"project/go-fiber-boilerplate/utils"
// 	"project/go-fiber-boilerplate/utils/constants"
// 	"testing"

// 	"github.com/stretchr/testify/suite"
// )

// type UserSuiteTest struct {
// 	suite.Suite
// 	client utils.ISuite
// 	res    *constants.Response
// }

// func (suite *UserSuiteTest) SetupSuite() {
// 	suite.client = utils.NewSuiteUtils(&http.Client{})
// 	suite.res = &constants.Response{}
// }

// func (suite *UserSuiteTest) TestFindAllUsers() {
// 	res, resp, err := suite.client.Get("http://localhost:3000/user", suite.res, "")

// 	suite.NoError(err)
// 	suite.Equal(http.StatusOK, resp.StatusCode)

// 	fmt.Println(res)
// 	suite.Equal("success get all users", res.Message)
// }

// func (suite *UserSuiteTest) TestFindUserByID() {
// 	req := dto.Login{
// 		Email:    "test@gmail.com",
// 		Password: "12345678",
// 	}

// 	jsonBytes2, err := json.Marshal(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	reqBody2 := bytes.NewBuffer(jsonBytes2)

// 	res1, resp1, err := suite.client.Post("http://localhost:3000/auth/login", reqBody2, suite.res, "")

// 	suite.NoError(err)
// 	suite.Equal(http.StatusOK, resp1.StatusCode)

// 	fmt.Println(res1)

// 	id := "1"
// 	res, resp, err := suite.client.Get(
// 		fmt.Sprintf("http://localhost:3000/user/%s", id),
// 		suite.res,
// 		"",
// 	)

// 	suite.NoError(err)
// 	suite.Equal(http.StatusOK, resp.StatusCode)

// 	fmt.Println(res)
// 	suite.Equal("success get user by id", res.Message)
// 	// suite.NotEmpty(res.Result)
// }

// func TestUserTestSuite(t *testing.T) {
// 	suite.Run(t, new(UserSuiteTest))
// }
