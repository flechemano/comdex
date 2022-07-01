package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewMsgWhitelistAppID(appID uint64, from string) *WhitelistAppId {
	return &WhitelistAppId{
		AppId: appID,
		From:  from,
	}
}

func (m *WhitelistAppId) Route() string {
	return RouterKey
}

func (m *WhitelistAppId) Type() string {
	return ModuleName
}

func (m *WhitelistAppId) ValidateBasic() error {
	return nil
}

func (m *WhitelistAppId) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *WhitelistAppId) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgRemoveWhitelistAsset(appID uint64, from string) *RemoveWhitelistAppId {
	return &RemoveWhitelistAppId{
		AppId: appID,
		From:  from,
	}
}

func (m *RemoveWhitelistAppId) Route() string {
	return RouterKey
}

func (m *RemoveWhitelistAppId) Type() string {
	return ModuleName
}

func (m *RemoveWhitelistAppId) ValidateBasic() error {
	return nil
}

func (m *RemoveWhitelistAppId) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *RemoveWhitelistAppId) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
