# Treasure Bridged Token Registry

This repository contains the whitelisted token registries for the Treasure Bridge. It maintains lists of approved ERC20, ERC721, and ERC1155 tokens that can be displayed and bridged through the Treasure Bridge interface.

## Structure

The repository organizes token registries by network and token standard:

data/
├── mainnet/
│   ├── erc20.json   - Mainnet ERC20 token whitelist
│   ├── erc721.json  - Mainnet ERC721 token whitelist
│   └── erc1155.json - Mainnet ERC1155 token whitelist
└── testnet/
    ├── erc20.json   - Testnet ERC20 token whitelist
    ├── erc721.json  - Testnet ERC721 token whitelist
    └── erc1155.json - Testnet ERC1155 token whitelist

## Token Format

Each JSON file contains a `tokens` array with bridge-specific token configurations. The format varies by token standard:

### ERC20 Example
{
  "name": "Token Name",
  "srcChainId": 5,
  "destChainId": 80001,
  "srcTokenAddress": "0x...",
  "destTokenAddress": "0x...",
  "isNft": false,
  "isCollateral": true,
  "underlyingTokenAddress": "0x..."  // Optional: required if isCollateral is true
}

### ERC721 Example
{
  "name": "Collection Name",
  "srcChainId": 421613,
  "destChainId": 5,
  "srcTokenAddress": "0x...",
  "destTokenAddress": "0x...",
  "isNft": true,
  "nftStandard": "ERC721",
  "isCollateral": false,
  "underlyingTokenAddress": "0x...",  // Optional: required if isCollateral is true
  "nftBatchSignature": "send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes),(uint256,uint256),address)"
}

### ERC1155 Example
{
  "name": "Collection Name",
  "srcChainId": 5,
  "destChainId": 420,
  "srcTokenAddress": "0x...",
  "destTokenAddress": "0x...",
  "isNft": true,
  "nftStandard": "ERC1155",
  "isCollateral": true,
  "underlyingTokenAddress": "0x...",  // Optional: required if isCollateral is true
  "nftBatchSignature": "send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes),(uint256,uint256),address)"
}

## Validation

This repository includes a validation tool written in Go that verifies the integrity and format of all token configurations.

### Prerequisites
- Go 1.16 or later
- Git

### Running the Validator

1. Clone the repository:
   ```bash
   git clone https://github.com/TreasureProject/bridging-token-registry.git
   cd bridging-token-registry
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the validator:
   ```bash
   go run main.go
   ```

The validator performs the following checks:
- JSON syntax validation
- Schema validation for each token type
- Required fields presence
- Address format validation
- Chain ID validation
- Duplicate entry detection
- NFT standard compliance
- Batch signature format validation

If any issues are found, the validator will output detailed error messages indicating the problems and their locations in the JSON files.

## Contributing

We welcome contributions from the community! If you'd like to propose adding an ERC20, ERC721, or ERC1155 token to be whitelisted for display in the Treasure Bridge, please see our [CONTRIBUTING.md](CONTRIBUTING.md) guide.
