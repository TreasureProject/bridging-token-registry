package main

import (
	"fmt"
	"os"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

type validationError struct {
	tokenType TokenType
	err       error
	file      string
}

type TokenType string

const (
	ERC20   TokenType = "erc20"
	ERC721  TokenType = "erc721"
	ERC1155 TokenType = "erc1155"
)

func main() {
	var errors []validationError
	comp := jsonschema.NewCompiler()

	// Single schema for all token types
	schemaLocation := "schemas/tokenSchema.json"

	// Data file locations by token type
	dataFiles := map[TokenType][2]string{
		ERC20: {
			"testnet/erc20-data.json",
			"mainnet/erc20-data.json",
		},
		ERC721: {
			"testnet/erc721-data.json",
			"mainnet/erc721-data.json",
		},
		ERC1155: {
			"testnet/erc1155-data.json",
			"mainnet/erc1155-data.json",
		},
	}

	// Compile the single schema
	if _, err := os.Stat(schemaLocation); os.IsNotExist(err) {
		fmt.Printf("Schema file not found: %s\n", err)
		os.Exit(1)
	}

	schema, err := comp.Compile(schemaLocation)
	if err != nil {
		fmt.Printf("Schema compilation error: %s\n", err)
		os.Exit(1)
	}

	// Validate each data file against the schema
	for tokenType, locations := range dataFiles {
		for _, fileLocation := range locations {
			if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
				errors = append(errors, validationError{
					tokenType: tokenType,
					err:       err,
					file:      fileLocation,
				})
				continue
			}

			file, err := os.Open(fileLocation)
			if err != nil {
				errors = append(errors, validationError{
					tokenType: tokenType,
					err:       err,
					file:      fileLocation,
				})
				continue
			}

			inst, err := jsonschema.UnmarshalJSON(file)
			if err != nil {
				errors = append(errors, validationError{
					tokenType: tokenType,
					err:       err,
					file:      fileLocation,
				})
				file.Close()
				continue
			}

			file.Close()

			err = schema.Validate(inst)
			if err != nil {
				errors = append(errors, validationError{
					tokenType: tokenType,
					err:       err,
					file:      fileLocation,
				})
			}
		}
	}

	if len(errors) > 0 {
		fmt.Println("Validation errors occurred:")
		for _, err := range errors {
			fmt.Printf("%s: %s in file %s\n", err.tokenType, err.err.Error(), err.file)
		}
		os.Exit(1)
	}

	fmt.Println("All token configurations validated successfully")
}
