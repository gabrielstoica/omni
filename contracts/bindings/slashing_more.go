package bindings

import (
	_ "embed"
)

const (
	SlashingDeployedBytecode = "0x6080604052348015600f57600080fd5b506004361060285760003560e01c8063f679d30514602d575b600080fd5b60336035565b005b60405133907fc3ef55ddda4bc9300706e15ab3aed03c762d8afd43a7d358a7b9503cb39f281b90600090a256fea264697066735822122046125cf3aa1770dee7f65080efcf7560c19924e2c15e3b3fec4cb1deb26fda5164736f6c63430008180033"
)

//go:embed slashing_storage_layout.json
var slashingStorageLayoutJSON []byte

var SlashingStorageLayout = mustGetStorageLayout(slashingStorageLayoutJSON)
