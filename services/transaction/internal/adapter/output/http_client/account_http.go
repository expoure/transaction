package http_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/expoure/pismo/transaction/internal/adapter/output/model/response_api"
	"github.com/expoure/pismo/transaction/internal/application/constants"
	"github.com/expoure/pismo/transaction/internal/application/port/output"
	"github.com/expoure/pismo/transaction/internal/configuration/customized_errors"
)

func NewAccountHttpClient() output.AccountHttpClient {
	return &accountHttpClientImpl{
		httpClient: &http.Client{},
	}
}

type accountHttpClientImpl struct {
	httpClient *http.Client
}

func (c *accountHttpClientImpl) GetAccount(id string) (*response_api.Account, error) {
	url := os.Getenv("ACCOUNT_SERVICE_URL") + "/v1/accounts/" + id
	request, _ := http.NewRequest(http.MethodGet, url, nil)

	result, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, customized_errors.NewBadRequestError(constants.ErrWasNotPossibleToCreateTransaction)
	}

	defer result.Body.Close()
	body, err := ioutil.ReadAll(result.Body)

	var account response_api.Account
	if err := json.Unmarshal(body, &account); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return &account, nil
}
