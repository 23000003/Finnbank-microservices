package entities

import (
	"github.com/graphql-go/graphql"
)

func GetProductEntityType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Product",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"info": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Float,
				},
			},
		},
	)
}

func GetAccountEntityType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Account",
			Fields: graphql.Fields{
				"account_id": &graphql.Field{
					Type: graphql.String, // UUID in DB
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"full_name": &graphql.Field{
					Type: graphql.String,
				},
				"phone_number": &graphql.Field{
					Type: graphql.String,
				},
				"password": &graphql.Field{
					Type: graphql.String, // encrypted
				},
				"date_created": &graphql.Field{
					Type: graphql.DateTime,
				},
				"account_number": &graphql.Field{
					Type: graphql.String,
				},
				"has_card": &graphql.Field{
					Type: graphql.Boolean,
				},
				"address": &graphql.Field{
					Type: graphql.String,
				},
				"balance": &graphql.Field{
					Type: graphql.Float,
				},
				"account_type": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}

func GetBankCardEntityType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "BankCard",
			Fields: graphql.Fields{
				"card_number": &graphql.Field{
					Type: graphql.String, // UUID in DB
				},
				"expiry": &graphql.Field{
					Type: graphql.DateTime,
				},
				"account_id": &graphql.Field{
					Type: graphql.Int,
				},
				"cvv": &graphql.Field{
					Type: graphql.Int,
				},
				"pin_number": &graphql.Field{
					Type: graphql.String, // encrypted
				},
				"date_created": &graphql.Field{
					Type: graphql.DateTime,
				},
				"card_type": &graphql.Field{
					Type: graphql.EnumValueType,
				},
			},
		},
	)
}

func GetTransactionEntityType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Transaction",
			Fields: graphql.Fields{
				"ref_number": &graphql.Field{
					Type: graphql.String, // UUID in DB
				},
				"sender": &graphql.Field{
					Type: graphql.Int,
				},
				"receiver": &graphql.Field{
					Type: graphql.Int,
				},
				"transaction_type": &graphql.Field{
					Type: graphql.EnumValueType,
				},
				"amount": &graphql.Field{
					Type: graphql.Int,
				},
				"date_created": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		},
	)
}

func GetNotificationEntityType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Notification",
			Fields: graphql.Fields{
				"notif_id": &graphql.Field{
					Type: graphql.Int,
				},
				"notif_type": &graphql.Field{
					Type: graphql.EnumValueType,
				},
				"account_id": &graphql.Field{
					Type: graphql.Int,
				},
				"redirect_url": &graphql.Field{
					Type: graphql.String,
				},
				"date_notified": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		},
	)
}
