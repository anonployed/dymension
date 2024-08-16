package keeper

import (
	"fmt"
	"strings"

	errorsmod "cosmossdk.io/errors"
	"github.com/dymensionxyz/gerr-cosmos/gerrc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	dymnstypes "github.com/dymensionxyz/dymension/v3/x/dymns/types"
	dymnsutils "github.com/dymensionxyz/dymension/v3/x/dymns/utils"
)

// SetDymName stores a Dym-Name into the KVStore.
//
// Important Note:
//  1. Must call BeforeDymNameOwnerChanged and AfterDymNameOwnerChanged before and after calling this function when updating owner.
//  2. Must call BeforeDymNameConfigChanged and AfterDymNameConfigChanged before and after calling this function when updating configuration.
func (k Keeper) SetDymName(ctx sdk.Context, dymName dymnstypes.DymName) error {
	if err := dymName.Validate(); err != nil {
		return err
	}

	// persist record
	store := ctx.KVStore(k.storeKey)
	dymNameKey := dymnstypes.DymNameKey(dymName.Name)
	bz := k.cdc.MustMarshal(&dymName)
	store.Set(dymNameKey, bz)
	ctx.EventManager().EmitEvent(dymName.GetSdkEvent())

	return nil
}

// BeforeDymNameOwnerChanged must be called before updating the owner of a Dym-Name.
// This function will remove the reverse mapping from the old owner to the Dym-Name.
func (k Keeper) BeforeDymNameOwnerChanged(ctx sdk.Context, name string) error {
	dymName := k.GetDymName(ctx, name)
	if dymName == nil {
		return nil
	}

	return k.RemoveReverseMappingOwnerToOwnedDymName(ctx, dymName.Owner, dymName.Name)
}

// AfterDymNameOwnerChanged must be called after the owner of a Dym-Name is changed.
// This function will add the reverse mapping from the new owner to the Dym-Name.
func (k Keeper) AfterDymNameOwnerChanged(ctx sdk.Context, name string) error {
	dymName := k.GetDymName(ctx, name)
	if dymName == nil {
		return errorsmod.Wrapf(gerrc.ErrNotFound, "Dym-Name: %s", name)
	}

	return k.AddReverseMappingOwnerToOwnedDymName(ctx, dymName.Owner, name)
}

// BeforeDymNameConfigChanged must be called before updating the configuration of a Dym-Name.
// This function will remove the reverse mapping from the configured addresses and fallback addresses to the Dym-Name.
func (k Keeper) BeforeDymNameConfigChanged(ctx sdk.Context, name string) error {
	dymName := k.GetDymName(ctx, name)
	if dymName == nil {
		return nil
	}

	configuredAddresses, fallbackAddresses := dymName.GetAddressesForReverseMapping()
	for _, configuredAddress := range dymnsutils.GetSortedStringKeys(configuredAddresses) {
		if err := k.RemoveReverseMappingConfiguredAddressToDymName(ctx, configuredAddress, name); err != nil {
			return err
		}
	}
	for _, fallbackAddress := range dymnsutils.GetSortedStringKeys(fallbackAddresses) {
		bz := dymnsutils.GetBytesFromHexAddress(fallbackAddress)
		if err := k.RemoveReverseMappingFallbackAddressToDymName(ctx, bz, name); err != nil {
			return err
		}
	}

	return nil
}

// AfterDymNameConfigChanged must be called after the configuration of a Dym-Name is changed.
// This function will add the reverse mapping from the configured addresses and fallback addresses to the Dym-Name.
func (k Keeper) AfterDymNameConfigChanged(ctx sdk.Context, name string) error {
	dymName := k.GetDymName(ctx, name)
	if dymName == nil {
		return errorsmod.Wrapf(gerrc.ErrNotFound, "Dym-Name: %s", name)
	}

	configuredAddresses, fallbackAddresses := dymName.GetAddressesForReverseMapping()
	for _, configuredAddress := range dymnsutils.GetSortedStringKeys(configuredAddresses) {
		if err := k.AddReverseMappingConfiguredAddressToDymName(ctx, configuredAddress, name); err != nil {
			return err
		}
	}
	for _, fallbackAddrAsHex := range dymnsutils.GetSortedStringKeys(fallbackAddresses) {
		bz := dymnsutils.GetBytesFromHexAddress(fallbackAddrAsHex)
		if err := k.AddReverseMappingFallbackAddressToDymName(ctx, bz, name); err != nil {
			return err
		}
	}

	return nil
}

// GetDymName returns a Dym-Name from the KVStore.
func (k Keeper) GetDymName(ctx sdk.Context, name string) *dymnstypes.DymName {
	store := ctx.KVStore(k.storeKey)
	dymNameKey := dymnstypes.DymNameKey(name)

	bz := store.Get(dymNameKey)
	if bz == nil {
		return nil
	}

	var dymName dymnstypes.DymName
	k.cdc.MustUnmarshal(bz, &dymName)

	return &dymName
}

// GetDymNameWithExpirationCheck returns a Dym-Name from the KVStore, if the Dym-Name is not expired.
// Returns nil if Dym-Name does not exist or is expired.
func (k Keeper) GetDymNameWithExpirationCheck(ctx sdk.Context, name string) *dymnstypes.DymName {
	// Legacy TODO DymNS: always use this on queries
	dymName := k.GetDymName(ctx, name)
	if dymName == nil {
		return nil
	}

	if dymName.IsExpiredAtCtx(ctx) {
		return nil
	}

	return dymName
}

// DeleteDymName removes a Dym-Name from the KVStore.
// This function will remove the Dym-Name record as well as the reverse mappings records.
func (k Keeper) DeleteDymName(ctx sdk.Context, name string) error {
	if err := k.BeforeDymNameOwnerChanged(ctx, name); err != nil {
		return err
	}

	if err := k.BeforeDymNameConfigChanged(ctx, name); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	dymNameKey := dymnstypes.DymNameKey(name)
	store.Delete(dymNameKey)

	return nil
}

// GetAllNonExpiredDymNames returns all non-expired Dym-Names from the KVStore.
// This function will filter out all expired Dym-Names based on the time of the context.
func (k Keeper) GetAllNonExpiredDymNames(ctx sdk.Context) (list []dymnstypes.DymName) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, dymnstypes.KeyPrefixDymName)
	defer func() {
		_ = iterator.Close()
	}()

	for ; iterator.Valid(); iterator.Next() {
		var dymName dymnstypes.DymName
		k.cdc.MustUnmarshal(iterator.Value(), &dymName)

		if dymName.IsExpiredAtCtx(ctx) {
			continue
		}

		list = append(list, dymName)
	}

	return
}

// GetAllDymNames returns all Dym-Names from the KVStore.
// No filter applied, to be used in genesis export.
func (k Keeper) GetAllDymNames(ctx sdk.Context) (list []dymnstypes.DymName) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, dymnstypes.KeyPrefixDymName)
	defer func() {
		_ = iterator.Close()
	}()

	for ; iterator.Valid(); iterator.Next() {
		var dymName dymnstypes.DymName
		k.cdc.MustUnmarshal(iterator.Value(), &dymName)
		list = append(list, dymName)
	}

	return list
}

// PruneDymName removes a Dym-Name from the KVStore, as well as all related records.
func (k Keeper) PruneDymName(ctx sdk.Context, name string) error {
	// remove SO (force, ignore active SO)
	k.DeleteSellOrder(ctx, name, dymnstypes.TypeName)

	dymName := k.GetDymName(ctx, name)
	if dymName == nil {
		return nil
	}

	// remove config
	// This seems not necessary because we are going to remove the record anyway,
	// but just let it here to clear the business logic
	dymName.Configs = nil   // all configuration should be removed
	dymName.Owner = ""      // no one owns it anyone
	dymName.Controller = "" // no one controls it anyone

	return k.DeleteDymName(ctx, name)
}

// ResolveByDymNameAddress resolves a Dym-Name-Address into an output address.
//
// For example:
//   - "my-name@dym" => "dym1a..."
//   - "another.my-name@dym" => "dym1b..."
//   - "my-name@nim" => "nim1..."
//   - (extra format) "0x1234...6789@nim" => "nim1..."
//   - (extra format) "dym1a...@nim" => "nim1..."
func (k Keeper) ResolveByDymNameAddress(ctx sdk.Context, dymNameAddress string) (outputAddress string, err error) {
	subName, name, chainIdOrAlias, parseErr := ParseDymNameAddress(dymNameAddress)
	if parseErr != nil {
		ctx.Logger().Debug("failed to parse Dym-Name address", "dym-name-address", dymNameAddress, "error", parseErr)
		err = parseErr
		return
	}

	dymName := k.GetDymNameWithExpirationCheck(ctx, name)
	if dymName == nil {
		err = errorsmod.Wrapf(gerrc.ErrNotFound, "Dym-Name: %s", name)

		// Dym-Name not found, in this case, there are 3 possible reasons:
		// 1. Dym-Name does not exist
		// 2. Dym-Name was expired
		// 3. Resolve extra format 0x1234...6789@nim and dym1...@nim
		// First two cases, stop here.
		// If it is the third case, we need to resolve it.

		if subName != "" {
			return
		}

		outputAddressFromExtraFormat, success := k.resolveByDymNameAddressInExtraFormat(ctx, name, chainIdOrAlias)
		if !success {
			return
		}

		outputAddress = outputAddressFromExtraFormat
		err = nil
		return
	}

	defer func() {
		if outputAddress == "" {
			err = errorsmod.Wrapf(gerrc.ErrNotFound, "no resolution found")
			return
		}
	}()

	tryResolveFromConfig := func(lookupChainIdConfig string) (value string, found bool) {
		if lookupChainIdConfig == ctx.ChainID() {
			lookupChainIdConfig = ""
		}

		for _, config := range dymName.Configs {
			if config.Type != dymnstypes.DymNameConfigType_DCT_NAME {
				continue
			}

			if config.ChainId != lookupChainIdConfig {
				continue
			}

			if config.Path != subName {
				continue
			}

			return config.Value, true
		}

		return "", false
	}

	// first attempt to resolve, to see if the chain-id/alias is user-configured
	var found bool
	outputAddress, found = tryResolveFromConfig(chainIdOrAlias)
	if found {
		return
	}

	// end of first attempt

	var resolvedToChainId string

	if chainId, success := k.tryResolveChainIdOrAliasToChainId(ctx, chainIdOrAlias); success {
		resolvedToChainId = chainId
	} else {
		// treat it as chain-id
		resolvedToChainId = chainIdOrAlias
	}

	// second attempt to resolve
	outputAddress, found = tryResolveFromConfig(resolvedToChainId)
	if found {
		return
	}
	// end of second attempt

	// no more attempts to resolve from config

	// try fallback

	if subName != "" {
		// no need to fallback
		return
	}

	if resolvedToChainId == ctx.ChainID() {
		outputAddress = dymName.Owner
		return
	}

	resolveToAddress := dymName.Owner
	for _, config := range dymnstypes.DymNameConfigs(dymName.Configs).DefaultNameConfigs(true) {
		resolveToAddress = config.Value
		break
	}

	// not the host chain, try to resolve if is a RollApp

	isRollAppId := k.IsRollAppId(ctx, resolvedToChainId)
	if !isRollAppId {
		// fallback does not apply for non-RollApp
		return
	}

	rollAppBech32Prefix, found := k.GetRollAppBech32Prefix(ctx, resolvedToChainId)
	if !found {
		// fallback does not apply for RollApp does not have this metadata
		return
	}

	accAddr := sdk.MustAccAddressFromBech32(resolveToAddress)
	rollAppBasedBech32Addr, convertErr := bech32.ConvertAndEncode(rollAppBech32Prefix, accAddr)
	if convertErr != nil {
		err = errorsmod.Wrapf(
			gerrc.ErrUnknown,
			"failed to convert address to RollApp-based address: %s", resolveToAddress,
		)
		return
	}

	outputAddress = rollAppBasedBech32Addr
	return
}

func (k Keeper) resolveByDymNameAddressInExtraFormat(ctx sdk.Context, anyAddress, chainIdOrAlias string) (outputAddress string, success bool) {
	var accAddr sdk.AccAddress
	if dymnsutils.IsValidHexAddress(anyAddress) {
		accAddr = dymnsutils.GetBytesFromHexAddress(anyAddress)
	} else if dymnsutils.IsValidBech32AccountAddress(anyAddress, false) {
		_, bz, errDecode := bech32.DecodeAndConvert(anyAddress)
		if errDecode != nil {
			// silent error on purpose, just skip it
			return
		}

		accAddr = bz
	} else {
		// neither hex address nor bech32 account address
		return
	}

	chainId, resolveSuccess := k.tryResolveChainIdOrAliasToChainId(ctx, chainIdOrAlias)
	if !resolveSuccess {
		// not a known chain-id or alias
		return
	}

	// only accept resolve for host chain or RollApp

	if chainId == ctx.ChainID() {
		// is host chain
		outputAddress = accAddr.String()
		success = true
		return
	}

	if !k.IsRollAppId(ctx, chainId) {
		// don't do bech32 conversion for non-RollApp because we don't know bech32 prefix
		return
	}

	bech32Prefix, found := k.GetRollAppBech32Prefix(ctx, chainId)
	if !found {
		// no bech32 prefix configured for this RollApp
		return
	}
	rollAppBasedBech32Addr, convertErr := bech32.ConvertAndEncode(bech32Prefix, accAddr)
	if convertErr != nil {
		return
	}
	fmt.Println(rollAppBasedBech32Addr)

	outputAddress = rollAppBasedBech32Addr
	success = true
	return
}

// ReplaceChainIdWithAliasIfPossible replaces the chain-id with alias if possible.
func (k Keeper) tryResolveChainIdOrAliasToChainId(ctx sdk.Context, chainIdOrAlias string) (resolvedToChainId string, success bool) {
	if chainIdOrAlias == ctx.ChainID() {
		return chainIdOrAlias, true
	}

	chainsParams := k.ChainsParams(ctx)
	if len(chainsParams.AliasesOfChainIds) > 0 {
		for _, record := range chainsParams.AliasesOfChainIds {
			if chainIdOrAlias == record.ChainId {
				return record.ChainId, true
			}

			for _, alias := range record.Aliases {
				if alias == chainIdOrAlias {
					return record.ChainId, true
				}
			}
		}
	}

	if isRollAppId := k.IsRollAppId(ctx, chainIdOrAlias); isRollAppId {
		return chainIdOrAlias, true
	}

	if rollAppId, found := k.GetRollAppIdByAlias(ctx, chainIdOrAlias); found {
		return rollAppId, true
	}

	return
}

// ParseDymNameAddress parses a Dym-Name address into its components.
//
// The components are:
//  1. Sub-Name is the part before the last '.' before the '@'.
//  2. Dym-Name is the part before the last '@'.
//  3. Chain-ID or Alias is the part after the last '@'.
//
// The '@' can replaced with '.'.
func ParseDymNameAddress(
	dymNameAddress string,
) (
	subName string, dymName string, chainIdOrAlias string, err error,
) {
	dymNameAddress = strings.ToLower(strings.TrimSpace(dymNameAddress))

	lastDotIndex := strings.LastIndex(dymNameAddress, ".")
	lastAtIndex := strings.LastIndex(dymNameAddress, "@")

	if lastAtIndex > -1 && lastDotIndex > -1 {
		if lastDotIndex > lastAtIndex {
			// do not accept '.' at chain-id/alias part
			err = errorsmod.Wrap(dymnstypes.ErrBadDymNameAddress, "misplaced '.'")
			return
		}
	}

	firstDotIndex := strings.IndexRune(dymNameAddress, '.')
	firstAtIndex := strings.IndexRune(dymNameAddress, '@')
	if firstAtIndex > -1 {
		if firstAtIndex != lastAtIndex {
			err = errorsmod.Wrap(dymnstypes.ErrBadDymNameAddress, "multiple '@' found")
			return
		}
	}

	if firstDotIndex == 0 || firstAtIndex == 0 {
		err = dymnstypes.ErrBadDymNameAddress
		return
	}

	if lastCharIdx := len(dymNameAddress) - 1; firstDotIndex == lastCharIdx ||
		firstAtIndex == lastCharIdx ||
		lastDotIndex == lastCharIdx ||
		lastAtIndex == lastCharIdx {
		err = dymnstypes.ErrBadDymNameAddress
		return
	}

	if strings.Contains(strings.ReplaceAll(strings.ReplaceAll(dymNameAddress, ".", "|"), "@", "|"), "||") {
		err = dymnstypes.ErrBadDymNameAddress
		return
	}

	chunks := strings.FieldsFunc(dymNameAddress, func(r rune) bool {
		return r == '.' || r == '@'
	})
	for i, chunk := range chunks {
		normalizedChunk := strings.TrimSpace(chunk)
		if normalizedChunk != chunk {
			err = dymnstypes.ErrBadDymNameAddress
			return
		}
		chunks[i] = normalizedChunk
	}

	if len(chunks) == 1 {
		// only Dym-Name, without chain-id/alias,... That is not accepted
		err = dymnstypes.ErrBadDymNameAddress
		return
	}

	chainIdOrAlias = chunks[len(chunks)-1]
	dymName = chunks[len(chunks)-2]
	if len(chunks) > 2 {
		subNameParts := chunks[:len(chunks)-2]
		for _, subNamePart := range subNameParts {
			if !dymnsutils.IsValidDymName(subNamePart) {
				err = errorsmod.Wrapf(
					dymnstypes.ErrBadDymNameAddress,
					"Sub-Dym-Name part is not well-formed: %s", subNamePart,
				)
				return
			}
		}
		subName = strings.Join(subNameParts, ".")
	}

	if !dymnsutils.IsValidChainIdFormat(chainIdOrAlias) &&
		!dymnsutils.IsValidAlias(chainIdOrAlias) {
		err = errorsmod.Wrapf(dymnstypes.ErrBadDymNameAddress, "chain-id/alias is not well-formed: %s", chainIdOrAlias)
		return
	}

	if subName == "" {
		// when no sub-name, we have 2 valid formats that won't pass Dym-Name validation

		// 0x1234...6789@nim
		if dymnsutils.IsValidHexAddress(dymName) {
			return
		} else if dymnsutils.IsValidBech32AccountAddress(dymName, false) {
			return
		}
	}

	if !dymnsutils.IsValidDymName(dymName) {
		err = errorsmod.Wrapf(dymnstypes.ErrBadDymNameAddress, "Dym-Name is not well-formed: %s", dymName)
		return
	}

	return
}

// ReverseResolveDymNameAddress resolves an address into a Dym-Name-Address which points to it.
// This function may return multiple possible Dym-Name-Addresses those point to the input address.
//
// For example: when we have "my-name@dym" resolves to "dym1a..."
// so reverse resolve will return "my-name@dym" when input is "dym1a..."
func (k Keeper) ReverseResolveDymNameAddress(ctx sdk.Context, inputAddress, workingChainId string) (outputDymNameAddresses dymnstypes.ReverseResolvedDymNameAddresses, err error) {
	if !dymnsutils.PossibleAccountRegardlessChain(inputAddress) {
		return nil, errorsmod.Wrapf(gerrc.ErrInvalidArgument, "not supported address format: %s", inputAddress)
	}

	isBech32Addr := dymnsutils.IsValidBech32AccountAddress(inputAddress, false)
	is0xAddr := dymnsutils.IsValidHexAddress(inputAddress)

	if !dymnsutils.IsValidChainIdFormat(workingChainId) {
		return nil, errorsmod.Wrapf(gerrc.ErrInvalidArgument, "invalid chain-id format: %s", workingChainId)
	}

	workingChainIdIsHostChain := workingChainId == ctx.ChainID()
	workingChainIdIsRollApp := !workingChainIdIsHostChain && k.IsRollAppId(ctx, workingChainId)

	if workingChainIdIsHostChain || workingChainIdIsRollApp {
		if !isBech32Addr && !is0xAddr {
			return nil, errorsmod.Wrapf(
				gerrc.ErrInvalidArgument,
				"not supported address format in host-chain and RollApp: %s", inputAddress,
			)
		}

		// on host-chain and RollApps, guarantee of case-insensitive address
		inputAddress = strings.ToLower(inputAddress)
	} else if dymnsutils.IsValidHexAddress(inputAddress) {
		// if the address is hex format, then treat the chain is case-insensitive address,
		// like Ethereum, where the address is case-insensitive and checksum address contains mixed case
		inputAddress = strings.ToLower(inputAddress)
	}

	defer func() {
		outputDymNameAddresses = outputDymNameAddresses.Distinct()
		outputDymNameAddresses = k.ReplaceChainIdWithAliasIfPossible(ctx, outputDymNameAddresses)
		outputDymNameAddresses.Sort()
	}()

	if is0xAddr {
		hexAddr := inputAddress

		outputDymNameAddresses, err = k.reverseResolveDymNameAddressUsingHexAddress(
			ctx,
			hexAddr,
			workingChainId,
			workingChainIdIsHostChain, workingChainIdIsRollApp,
		)
		if err != nil {
			return nil, err
		}

		if len(outputDymNameAddresses) > 0 {
			// there is at least one result, can stop here
			return
		}

		return k.fallbackReverseResolveDymNameAddress(
			ctx,
			dymnsutils.GetBytesFromHexAddress(hexAddr),
			workingChainId,
			workingChainIdIsHostChain, workingChainIdIsRollApp,
		)
	}

	// try lookup using configured address

	outputDymNameAddresses, err = k.reverseResolveDymNameAddressUsingConfiguredAddress(
		ctx,
		inputAddress,
		workingChainId,
	)
	if err != nil {
		return nil, err
	}

	if len(outputDymNameAddresses) > 0 {
		// there is at least one result, can stop here
		return
	}

	// There is no matching result from lookup by configured-address,
	// we are going to give it one more try using fallback-lookup.
	// If the working chain-id is a coin-type-60 chain-id.
	// If the working chain-id is NOT a coin-type-60 chain-id, does not satisfy the condition.
	// Only host-chain and RollApp are coin-type-60 chains and support reverse-resolve using fallback lookup,
	// while other chains are ignored for safety purpose, as we don't if they are coin-type-60 chains or not.

	if !workingChainIdIsHostChain && !workingChainIdIsRollApp {
		return
	}

	if !isBech32Addr {
		// not a bech32 address, we can't do fallback lookup
		return
	}

	bech32Addr := inputAddress

	_, bz, err2 := bech32.DecodeAndConvert(bech32Addr)
	if err2 != nil {
		// this can not happen because we have checked the format before
		panic("failed to decode bech32 address: " + bech32Addr)
	}

	return k.fallbackReverseResolveDymNameAddress(
		ctx,
		bz,
		workingChainId,
		workingChainIdIsHostChain, workingChainIdIsRollApp,
	)
}

// reverseResolveDymNameAddressUsingConfiguredAddress resolves an address into a Dym-Name-Address
// which points to its address.
func (k Keeper) reverseResolveDymNameAddressUsingConfiguredAddress(
	ctx sdk.Context,
	inputAddress,
	workingChainId string,
) (outputDymNameAddresses dymnstypes.ReverseResolvedDymNameAddresses, err error) {
	dymNames, err1 := k.GetDymNamesContainsConfiguredAddress(ctx, inputAddress)
	if err1 != nil {
		return nil, err1
	}

	for _, dymName := range dymNames {
		configuredAddresses, _ := dymName.GetAddressesForReverseMapping()
		configs := configuredAddresses[inputAddress]
		outputDymNameAddresses = outputDymNameAddresses.AppendConfigs(ctx, dymName,
			configs, func(address dymnstypes.ReverseResolvedDymNameAddress) bool {
				return address.ChainIdOrAlias == workingChainId
			},
		)
	}

	return
}

// reverseResolveDymNameAddressUsingHexAddress resolves a hex address into a Dym-Name-Address
// which points to its corresponding bech32 address.
func (k Keeper) reverseResolveDymNameAddressUsingHexAddress(
	ctx sdk.Context,
	hexAddr,
	workingChainId string,
	workingChainIdIsHostChain, workingChainIdIsRollApp bool,
) (outputDymNameAddresses dymnstypes.ReverseResolvedDymNameAddresses, err error) {
	if !dymnsutils.IsValidHexAddress(hexAddr) {
		return nil, errorsmod.Wrapf(gerrc.ErrInvalidArgument, "invalid hex address: %s", hexAddr)
	}

	bzAddr := dymnsutils.GetBytesFromHexAddress(hexAddr)

	// When the input address is a hex address,
	// we can assume the query comes from a coin-type-60 chain.
	// Only host-chain and RollApp are coin-type-60 chains and support reverse lookup,
	// while other chains are ignored.
	//
	// Should do a fallback lookup?
	// Depend on the working chain-id is recognized or not.
	// - If the working chain-id is recognized then DO a fallback lookup.
	// - If the working chain-id is NOT recognized, DO NOT do a fallback lookup.

	// But first, let try to convert it into a bech32 address to see if any result is available.
	var bech32Hrp string
	if workingChainIdIsHostChain {
		bech32Hrp = sdk.GetConfig().GetBech32AccountAddrPrefix()
	} else if rollappBech32Hrp, found := k.GetRollAppBech32Prefix(ctx, workingChainId); found {
		bech32Hrp = rollappBech32Hrp
	}

	{
		lookupKey := hexAddr
		if bech32Hrp != "" {
			// found bech32 prefix configured for this chain-id
			lookupKey = sdk.MustBech32ifyAddressBytes(bech32Hrp, bzAddr)
		}

		dymNames, err1 := k.GetDymNamesContainsConfiguredAddress(ctx, lookupKey)
		if err1 != nil {
			return nil, err1
		}

		for _, dymName := range dymNames {
			configuredAddresses, _ := dymName.GetAddressesForReverseMapping()
			configs := configuredAddresses[lookupKey]
			outputDymNameAddresses = outputDymNameAddresses.AppendConfigs(ctx, dymName,
				configs, func(address dymnstypes.ReverseResolvedDymNameAddress) bool {
					return address.ChainIdOrAlias == workingChainId
				},
			)
		}

		if len(outputDymNameAddresses) > 0 {
			// there is at least one result, can stop here
			return
		}
	}

	return k.fallbackReverseResolveDymNameAddress(
		ctx,
		bzAddr,
		workingChainId,
		workingChainIdIsHostChain, workingChainIdIsRollApp,
	)
}

// fallbackReverseResolveDymNameAddress performs a fallback lookup for reverse resolve Dym-Name-Address.
func (k Keeper) fallbackReverseResolveDymNameAddress(
	ctx sdk.Context,
	bzAddr []byte,
	workingChainId string,
	workingChainIdIsHostChain, workingChainIdIsRollApp bool,
) (outputDymNameAddresses dymnstypes.ReverseResolvedDymNameAddresses, err error) {
	// check if we should do a fallback lookup

	if !workingChainIdIsHostChain && !workingChainIdIsRollApp {
		// we don't do fallback lookup for this case, just for safety purpose
		return
	}

	dymNames, err2 := k.GetDymNamesContainsFallbackAddress(ctx, bzAddr)
	if err2 != nil {
		return nil, err2
	}

	fallbackAddr := dymnstypes.FallbackAddress(bzAddr)

	for _, dymName := range dymNames {
		_, fallbackAddresses := dymName.GetAddressesForReverseMapping()
		configs := fallbackAddresses[fallbackAddr.String()]

		// only accept fallback for the case of default config
		for range dymnstypes.DymNameConfigs(configs).DefaultNameConfigs(true) {
			outputDymNameAddresses = append(outputDymNameAddresses, dymnstypes.ReverseResolvedDymNameAddress{
				SubName:        "",
				Name:           dymName.Name,
				ChainIdOrAlias: workingChainId, // fallback
			})

			break // just take the first one
		}
	}

	return
}

// ReplaceChainIdWithAliasIfPossible replaces the chain-id with alias if possible for the reverse resolved records.
func (k Keeper) ReplaceChainIdWithAliasIfPossible(ctx sdk.Context, reverseResolvedRecords dymnstypes.ReverseResolvedDymNameAddresses) []dymnstypes.ReverseResolvedDymNameAddress {
	if len(reverseResolvedRecords) < 1 {
		return reverseResolvedRecords
	}

	resolvedCache := make(map[string]string)
	// Describe usage of Go Map: used for caching purpose, no iteration.

	for i, reverseResolvedRecord := range reverseResolvedRecords {
		chainIdOrAlias := reverseResolvedRecord.ChainIdOrAlias

		if chainIdOrAlias == "" {
			chainIdOrAlias = ctx.ChainID()
			reverseResolvedRecords[i].ChainIdOrAlias = chainIdOrAlias
		}

		if resolvedTo, found := resolvedCache[chainIdOrAlias]; found {
			if resolvedTo != chainIdOrAlias {
				reverseResolvedRecords[i].ChainIdOrAlias = resolvedTo
			}
			continue
		}

		aliases := k.GetEffectiveAliasesByChainId(ctx, chainIdOrAlias)

		if len(aliases) < 1 {
			// no alias found, no need to replace
			resolvedCache[chainIdOrAlias] = chainIdOrAlias // cache it
			continue
		}

		defaultAlias := aliases[0]

		reverseResolvedRecords[i].ChainIdOrAlias = defaultAlias
		resolvedCache[chainIdOrAlias] = defaultAlias // cache it
	}

	return reverseResolvedRecords
}
