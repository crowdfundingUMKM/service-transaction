package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"service-transaction/helper"
)

func GetInvestorId(input InvestorIdInput) (string, error) {
	// check service admin

	err := CheckServiceUserInvestor()
	if err != nil {
		return "", err
	}

	investorID := helper.UserInvestor{}
	investorID.UnixInvestor = input.UnixID
	// fetch get /getAdminID from service api
	serviceInvestor := os.Getenv("SERVICE_INVESTOR")
	// if service admin is empty return error
	if serviceInvestor == "" {
		return investorID.UnixInvestor, errors.New("service user investor is empty")
	}
	resp, err := http.Get(serviceInvestor + "/api/v1/investor/getUserInvestorID/" + investorID.UnixInvestor)
	if err != nil {
		return investorID.UnixInvestor, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return investorID.UnixInvestor, errors.New("failed to get user investor status or investor not found")
	}

	var response helper.UserInvestorResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Meta.Code != 200 {
		return "", errors.New(response.Meta.Message)
	} else if response.Data.StatusAccountInvestor == "deactive" {
		return "", errors.New("investor account is deactive")
	} else if response.Data.StatusAccountInvestor == "active" {
		return investorID.UnixInvestor, nil
	} else {
		return "", errors.New("invalid investor status")
	}
}

// verify token from service user investor
func VerifyTokenInvestor(input string) (string, error) {

	err := CheckServiceUserInvestor()
	if err != nil {
		return "", err
	}

	// fetch get /verifyToken from service api
	serviceAdmin := os.Getenv("SERVICE_INVESTOR")
	// if service admin is empty return error
	if serviceAdmin == "" {
		return "", errors.New("service user investor is empty")
	}
	req, err := http.NewRequest("GET", serviceAdmin+"/api/v1/verifyTokenInvestor", nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+input)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("invalid token, account deactive or token expired")
	}

	var response helper.UserInvestorResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Meta.Code != 200 {
		return "", errors.New(response.Meta.Message)
	}

	return response.Data.UnixInvestor, nil

}
