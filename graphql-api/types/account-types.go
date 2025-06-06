package types

import (
	"time"

	"github.com/graphql-go/graphql"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents a login response.
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int32  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	AuthID       string `json:"auth_id"`
	// below are values for auth context in FE
	DisplayName   string `json:"full_name"`
	AccountId     string `json:"account_id"`
	AccountStatus string `json:"account_status"`
	AccountType 	string `json:"account_type"`
}

// Account represents an account.
type Account struct {
	ID            string    `json:"account_id"`
	Email         string    `json:"email"`
	FirstName     string    `json:"first_name"`
	MiddleName    string    `json:"middle_name"`
	LastName      string    `json:"last_name"`
	PhoneNumber   string    `json:"phone_number"`
	Address       string    `json:"address"`
	AccountType   string    `json:"account_type"`
	AccountNumber string    `json:"account_number"`
	AuthID        string    `json:"auth_id"`
	HasCard       bool      `json:"has_card"`
	DateCreated   time.Time `json:"date_created"`
	DateUpdated   time.Time `json:"date_updated"`
	Nationality   string    `json:"nationality"`
	BirthDate     time.Time `json:"birthdate"`
	NationalID    string    `json:"national_id"`
	AccountStatus string    `json:"account_status"`
}

type AccountsResponse struct {
	Accounts []Account `json:"accounts"`
}

// AccountResponse represents a response containing a single account.
type AccountResponse struct {
	Account Account `json:"account"`
}

// UpdateRequest represents a request to update account details.
type UpdateRequest struct {
	AccountNumber string `json:"account_number"`
	FullName      string `json:"full_name"`
	PhoneNumber   string `json:"phone_number"`
	Address       string `json:"address"`
}

// CardUpdateRequest represents a request to update card details.
type CardUpdateRequest struct {
	AccountNumber string `json:"account_number"`
}

// CardUpdateResponse represents a response for a card update request.
type CardUpdateResponse struct {
	Status string `json:"status"`
}

// UpdatePasswordRequest represents a request to update a password.
type UpdatePasswordRequest struct {
	AccountID   string
	OldPassword string
	NewPassword string
}

// UpdatePasswordResponse represents a response for a password update request.
type UpdatePasswordResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// DeleteUserRequest represents a request to delete a user.
type DeleteUserRequest struct {
	AccountNumber string `json:"account_number"`
}

// DeleteUserResponse represents a response for a user deletion request.
type DeleteUserResponse struct {
	Message string `json:"message"`
}

// AddAccountRequest represents a request to add a new account.
type AddAccountRequest struct {
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	MiddleName  string    `json:"middle_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	Address     string    `json:"address"`
	AccountType string    `json:"account_type"`
	Nationality string    `json:"nationality"`
	BirthDate   time.Time `json:"birthdate"`
	NationalID  string    `json:"national_id"`
}

// AddAccountResponse represents a response for adding a new account.
type AddAccountResponse struct {
	Account Account `json:"account"`
}

type UpdateAccountRequest struct {
	AccountID  string `json:"account_id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
}

type UpdateAccountDetailsRequest struct {
	Type      string `json:"type"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	AccountID string `json:"account_id"`
}
type UpdateAccountStatusRequest struct {
	AccountID string `json:"account_id"`
	Type      string `json:"type"`
}

// For signup input
var AccountInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "AccountInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"first_name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"middle_name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"last_name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"phone_number": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"address": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"account_type": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"nationality": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"birthdate": &graphql.InputObjectFieldConfig{
				Type: graphql.DateTime,
			},
			"national_id": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

// using this for Login
var LoginInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "LoginInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

// using this for Update Account
var UpdatePasswordInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UpdatePasswordInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"auth_id": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"old_password": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"new_password": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)
var UpdateAccountInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UpdateAccountInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"account_id": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"first_name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"middle_name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"last_name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"phone": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"address": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)
var UpdateAccountDetailsInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UpdateAccountDetailsInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"account_id": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"phone": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"address": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"type": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)
var UpdateAccountStatusInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UpdateAccountStatusInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"account_id": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"type": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

var AuthResponseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AuthResponse",
		Fields: graphql.Fields{
			"access_token":   &graphql.Field{Type: graphql.String},
			"token_type":     &graphql.Field{Type: graphql.String},
			"expires_in":     &graphql.Field{Type: graphql.Int},
			"refresh_token":  &graphql.Field{Type: graphql.String},
			"auth_id":        &graphql.Field{Type: graphql.String},
			"full_name":      &graphql.Field{Type: graphql.String},
			"account_id":     &graphql.Field{Type: graphql.String},
			"account_status": &graphql.Field{Type: graphql.String},
			"account_type":   &graphql.Field{Type: graphql.String},
		},
	},
)
