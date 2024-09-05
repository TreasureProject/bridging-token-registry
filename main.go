package main

import (
	"fmt"
	"os"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

type vaildationError struct {
	tokenType TokenType
	err       error
}

type schemaType struct {
	schemaLocation   string
	dataFileLocation [2]string
	compiledSchema   *jsonschema.Schema
}

type TokenType string

const (
	ERC20   TokenType = "erc20"
	ERC721  TokenType = "erc721"
	ERC1155 TokenType = "erc1155"
)

func main() {
	var errors []vaildationError
	comp := jsonschema.NewCompiler()

	schemaMap := map[TokenType]*schemaType{
		ERC20: {
			schemaLocation:   "schemas/erc20-schema.json",
			dataFileLocation: [2]string{"testnet/erc20-data.json", "mainnet/erc20-data.json"},
		},
		ERC721: {
			schemaLocation: "schemas/erc721-schema.json",
			dataFileLocation: [2]string{
				"testnet/erc721-data.json",
				"mainnet/erc721-data.json",
			},
		},
		ERC1155: {
			schemaLocation: "schemas/erc1155-schema.json",
			dataFileLocation: [2]string{
				"testnet/erc1155-data.json",
				"mainnet/erc1155-data.json",
			},
		},
	}

	for token, schema := range schemaMap {

		if _, err := os.Stat(schema.schemaLocation); os.IsNotExist(err) {
			errors = append(errors, vaildationError{
				tokenType: token,
				err:       err,
			})
			continue
		}

		fmt.Println("got all schemas")

		sch, err := comp.Compile(schema.schemaLocation)
		if err != nil {
			errors = append(errors, vaildationError{
				tokenType: token,
				err:       err,
			})
			continue
		}

		schemaMap[token].compiledSchema = sch

		for _, fileLocation := range schema.dataFileLocation {
			if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
				errors = append(errors, vaildationError{
					tokenType: token,
					err:       err,
				})
				continue
			}

			file, err := os.Open(fileLocation)
			if err != nil {
				errors = append(errors, vaildationError{
					tokenType: token,
					err:       err,
				})
				continue
			}
			defer file.Close()

			inst, err := jsonschema.UnmarshalJSON(file)
			if err != nil {
				errors = append(errors, vaildationError{
					tokenType: token,
					err:       err,
				})
				continue
			}

			err = schema.compiledSchema.Validate(inst)
			if err != nil {
				errors = append(errors, vaildationError{
					tokenType: token,
					err:       err,
				})
				continue
			}

		}

	}

	if len(errors) > 0 {
		fmt.Println("vaildation errors occured")
		for _, err := range errors {
			fmt.Println(err.err.Error())
		}

		os.Exit(1)
	}

	fmt.Println("validation successful")

}
