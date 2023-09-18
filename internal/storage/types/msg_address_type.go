package types

// swagger:enum MsgAddressType
/*
	ENUM(
		validatorAddress,
		delegatorAddress,
		validatorSrcAddress,
		validatorDstAddress,
		fromAddress,
		toAddress,
		grantee,
		granter,
		signer,
		withdraw,

		voter,
		proposer,
	)
*/
//go:generate go-enum --marshal --sql --values
type MsgAddressType string
