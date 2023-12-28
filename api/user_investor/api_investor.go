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

	err := CheckServiceInvestor()
	if err != nil {
		return "", err
	}

	investorID := helper.UserInvestor{}
	investorID.UnixInvestor = input.UnixID
	// fetch get /getAdminID from service api
	serviceInvestor := os.Getenv("SERVICE_INVESTOR_HOST")
	// if service admin is empty return error
	if serviceInvestor == "" {
		return investorID.UnixInvestor, errors.New("service investor is empty")
	}
	resp, err := http.Get(serviceInvestor + "/api/v1/getInvestorID/" + investorID.UnixInvestor)
	if err != nil {
		return investorID.UnixInvestor, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return investorID.UnixInvestor, errors.New("failed to get user investor status or user investor not found")
	}

	var response helper.InvestorStatusResponse
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
		return "", errors.New("invalid user investor status")
	}
}

// verify token from service user admin
func VerifyTokenInvestor(input string) (string, error) {

	err := CheckServiceInvestor()
	if err != nil {
		return "", err
	}

	// fetch get /verifyToken from service api
	serviceInvestor := os.Getenv("SERVICE_INVESTOR_HOST")
	// if service admin is empty return error
	if serviceInvestor == "" {
		return "", errors.New("service user investor is empty")
	}
	req, err := http.NewRequest("GET", serviceInvestor+"/api/v1/verifyTokenInvestor", nil)

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

	var response helper.VerifyTokenApiInvestorResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Meta.Code != 200 {
		return "", errors.New(response.Meta.Message)
	}

	return response.Data.UnixInvestor, nil

}
