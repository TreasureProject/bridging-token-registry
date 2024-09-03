package main

import (
	"fmt"
	"log"
	"os"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

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
			log.Fatal("could not find all schemas")
		}

		fmt.Println("got all schemas")

		sch, err := comp.Compile(schema.schemaLocation)
		if err != nil {
			log.Fatalf("could not compile schema for %s", token)
		}

		schemaMap[token].compiledSchema = sch

		for _, fileLocation := range schema.dataFileLocation {
			if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
				log.Fatalf("cannot open data file for type %s", token)
			}

			file, err := os.Open(fileLocation)
			if err != nil {
				log.Fatal("could not open data file")
			}
			defer file.Close()

			inst, err := jsonschema.UnmarshalJSON(file)
			if err != nil {
				log.Fatal("could not unmarshall json")
			}

			schemaErr := schema.compiledSchema.Validate(inst)
			if schemaErr != nil {
				log.Fatal(schemaErr)
			}

		}

	}

}
