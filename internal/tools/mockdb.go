package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"test": {
		AuthToken: "test",
		Username:  "test",
	},
	"test2": {
		AuthToken: "test2",
		Username:  "test2",
	},
	"test3": {
		AuthToken: "test3",
		Username:  "test3",
	},
	"test4": {
		AuthToken: "test4",
		Username:  "test4",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"coin1": {
		Coins:    100,
		Username: "test",
	},
	"coin2": {
		Coins:    200,
		Username: "test2",
	},
	"coin3": {
		Coins:    300,
		Username: "test3",
	},
	"coin4": {
		Coins:    400,
		Username: "test4",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
