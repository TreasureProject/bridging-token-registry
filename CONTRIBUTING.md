# Contributing to Treasure Bridged Token Registry

This guide explains how to propose adding new tokens to the Treasure Bridge whitelist registry.

## Token Submission Process

1. Fork this repository
2. Add your token details to the appropriate JSON file in either `data/mainnet/` or `data/testnet/`:
   - ERC20 tokens: `erc20.json`
   - ERC721 tokens: `erc721.json`
   - ERC1155 tokens: `erc1155.json`
3. Create a Pull Request

## Token Information Requirements

All tokens should be added to the `tokens` array in the respective JSON file.

### ERC20 Tokens
Add your token to `erc20.json` using the following format:

{
  "name": "Token Name",
  "srcChainId": 5,                 // Source chain ID
  "destChainId": 80001,           // Destination chain ID
  "srcTokenAddress": "0x...",     // Token address on source chain
  "destTokenAddress": "0x...",    // Token address on destination chain
  "isNft": false,
  "isCollateral": true,           // Optional: true if token is collateral
  "underlyingTokenAddress": "0x..." // Optional: required if isCollateral is true
}

### ERC721 Tokens
Add your NFT collection to `erc721.json` using the following format:

{
  "name": "Collection Name",
  "srcChainId": 421613,
  "destChainId": 5,
  "srcTokenAddress": "0x...",
  "destTokenAddress": "0x...",
  "isNft": true,
  "nftStandard": "ERC721",
  "isCollateral": false,
  "underlyingTokenAddress": "0x...", // Optional: required if isCollateral is true
  "nftBatchSignature": "send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes),(uint256,uint256),address)"
}

### ERC1155 Tokens
Add your collection to `erc1155.json` using the following format:

{
  "name": "Collection Name",
  "srcChainId": 5,
  "destChainId": 420,
  "srcTokenAddress": "0x...",
  "destTokenAddress": "0x...",
  "isNft": true,
  "nftStandard": "ERC1155",
  "isCollateral": true,
  "underlyingTokenAddress": "0x...", // Optional: required if isCollateral is true
  "nftBatchSignature": "send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes),(uint256,uint256),address)"
}

## Submission Requirements

1. **Contract Verification**
   - Contract must be verified on the respective network's block explorer
   - Must be a valid implementation of the respective token standard (ERC20/721/1155)

2. **Security**
   - Contract should have no known vulnerabilities
   - Preferably audited by a reputable security firm

3. **PR Description Requirements**
   - Source and destination chain IDs
   - Contract addresses on both chains
   - Brief description of the token/collection
   - Link to project website
   - Link to verified contracts on block explorers
   - Link to audit report (if available)

## Pull Request Process

1. Ensure your token information follows the correct JSON format
2. Create a PR with a clear title (e.g., "Add MAGIC to ERC20 whitelist")
3. Fill out all required information in the PR description
4. Wait for review from maintainers
5. Address any requested changes
6. Once approved, your token will be added to the registry

## Questions or Issues?

If you have questions about the contribution process or encounter any issues, please open a GitHub issue in this repository. 