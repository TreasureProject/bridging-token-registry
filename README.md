# Treasure Bridged Token Registry

This repository maintains the official registry of bridged tokens in the Treasure ecosystem, supporting ERC20, ERC721, and ERC1155 tokens across multiple networks.

## Structure

The repository organizes token registries by network and token standard. In the `data` directory
you'll find

- `mainnet/`: Production token configurations
- `testnet/`: Test network token configurations

Each directory contains three JSON files:

- `erc20.json`: ERC20 token configurations
- `erc721.json`: ERC721 token configurations
- `erc1155.json`: ERC1155 token configurations

## Token Configuration

Each token entry requires the following mandatory fields:

```json
{
  "name": "Token Name",
  "srcChainId": 1234,
  "destChainId": 5678,
  "srcTokenAddress": "0x...",
  "destTokenAddress": "0x...",
  "isNft": false,
  "isCollateral": false
}
```

**Additional Required Fields:**

- For NFTs (`isNft: true`):
  - `nftStandard`: Either "ERC721" or "ERC1155"
  - `nftBatchSignature`: The function signature for batch operations
- For Collateral Tokens (`isCollateral: true`):
  - `underlyingTokenAddress`: The address of the underlying token

## Validation

The repository implements automatic validation through CI/CD pipelines. All token configurations must comply with the JSON schema defined in `schemas/tokenSchema.json`. The validation ensures:

- All required fields are present
- Addresses follow the correct format
- NFT configurations include necessary NFT-specific fields
- Collateral tokens specify their underlying token address

## Contributing

1. Fork the repository
2. Add your token configuration to the appropriate JSON file
3. Ensure your configuration passes schema validation
4. Submit a Pull Request
5. Contact your Treasure representative to review and approve your submission

## Local Validation

To validate configurations locally:

```bash
go run main.go
```

This command will check all JSON files against the schema and report any validation errors.

## Technical Requirements

- Go 1.22 or higher

## Support

For questions or assistance with token registration, please reach out to your Treasure point of contact. All submissions must go through proper review channels to ensure ecosystem security and stability.
