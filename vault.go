package main

import "time"

type Vault struct {
    BcryptHash string    `json:"bcrypt_hash"`
    Salt       string    `json:"salt"`
    CreatedAt  time.Time `json:"created_at"`
    Entries    []Entry   `json:"entries"`
}

type Entry struct {
    Site      string `json:"site"`
    Username  string `json:"username"`
    Password  string `json:"password"`
    Nonce     string `json:"nonce"`
    CreatedAt string `json:"created_at"`
}

func (v *Vault) AddEntry(e Entry) {
    v.Entries = append(v.Entries, e)
}

func (v *Vault) GetEntry(site string) *Entry {
    for _, e := range v.Entries {
        if e.Site == site {
            return &e
        }
    }
    return nil
}
