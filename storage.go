package main

import (
    "crypto/rand"
    "encoding/json"
    "io/ioutil"
    "os"
)

const vaultFile = "vault.json"

func LoadVault() (Vault, error) {
    data, err := ioutil.ReadFile(vaultFile)
    if err != nil {
        return Vault{}, err
    }

    var v Vault
    err = json.Unmarshal(data, &v)
    return v, err
}

func SaveVault(v Vault) error {
    data, err := json.MarshalIndent(v, "", "  ")
    if err != nil {
        return err
    }

    return ioutil.WriteFile(vaultFile, data, 0600)
}

func GenerateRandomSalt() ([]byte, error) {
    salt := make([]byte, 16)
    _, err := rand.Read(salt)
    return salt, err
}

func VaultExists() bool {
    _, err := os.Stat(vaultFile)
    return !os.IsNotExist(err)
}
