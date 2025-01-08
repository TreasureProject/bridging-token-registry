# Treasure Bridged Token Registry

This repository maintains the official registry of bridged tokens in the Treasure ecosystem, supporting ERC20, ERC721, and ERC1155 tokens across multiple networks.

## Structure

The registry is organized into two main directories:

- `mainnet/`: Production token configurations
- `testnet/`: Test network token configurations

Each directory contains three JSON files:

- `erc20-data.json`: ERC20 token configurations
- `erc721-data.json`: ERC721 token configurations
- `erc1155-data.json`: ERC1155 token configurations

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

The repository implements automatic validation through CI/CD pipelines[1]. All token configurations must comply with the JSON schema defined in `schemas/tokenSchema.json`[1]. The validation ensures:

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

This command will check all JSON files against the schema and report any validation errors[1].

## Technical Requirements

- Go 1.22.1 or higher
- JSON Schema validation using `github.com/santhosh-tekuri/jsonschema/v6`

## Support

For questions or assistance with token registration, please reach out to your Treasure point of contact. All submissions must go through proper review channels to ensure ecosystem security and stability.

Citations:
[1] https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/40858233/ae4cd89f-e486-4c57-aa23-2073ef8db4bd/paste.txt
