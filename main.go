package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	cmd := os.Args[1]

	switch cmd {

	// Create new vault
	case "init":
		initVault()

	// Add new entry
	case "add":
		addEntry()

	// List all
	case "list":
		listEntries()

	// Get entry
	case "get":
		if len(os.Args) < 3 {
			fmt.Println("Usage: get <site>")
			return
		}
		getEntry(os.Args[2])

	default:
		showUsage()
	}
}

// -------------------- COMMANDS --------------------

func showUsage() {
	fmt.Println("Commands:")
	fmt.Println("  init       Create new vault")
	fmt.Println("  add        Add a password entry")
	fmt.Println("  list       List all entries")
	fmt.Println("  get <site> View one entry")
}

// --- init command ---

func initVault() {
	if VaultExists() {
		fmt.Println("Vault already exists.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Create master password: ")
	pw, _ := reader.ReadString('\n')
	pw = strings.TrimSpace(pw)

	hash, _ := HashMasterPassword(pw)
	salt, _ := GenerateRandomSalt()

	v := Vault{
		BcryptHash: hash,
		Salt:       EncodeBase64(salt),
		CreatedAt:  time.Now(),
		Entries:    []Entry{},
	}

	SaveVault(v)
	fmt.Println("Vault created successfully!")
}

// --- add command ---

func addEntry() {
	v, _ := LoadVault()

	reader := bufio.NewReader(os.Stdin)

	// verify master pw
	fmt.Print("Master password: ")
	pw, _ := reader.ReadString('\n')
	pw = strings.TrimSpace(pw)

	err := VerifyMasterPassword(v.BcryptHash, pw)
	if err != nil {
		fmt.Println("Incorrect password.")
		return
	}

	// Derive AES key
	salt, _ := DecodeBase64(v.Salt)
	key := DeriveKey(pw, salt)

	// Collect new entry data
	fmt.Print("Site: ")
	site, _ := reader.ReadString('\n')
	site = strings.TrimSpace(site)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	pass, _ := reader.ReadString('\n')
	pass = strings.TrimSpace(pass)

	// Encrypt password
	ct, nonce, _ := EncryptAES([]byte(pass), key)

	e := Entry{
		Site:      site,
		Username:  username,
		Password:  EncodeBase64(ct),
		Nonce:     EncodeBase64(nonce),
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	v.AddEntry(e)
	SaveVault(v)

	fmt.Println("Entry saved!")
}

// --- list command ---

func listEntries() {
	v, _ := LoadVault()
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Master password: ")
	pw, _ := reader.ReadString('\n')
	pw = strings.TrimSpace(pw)

	err := VerifyMasterPassword(v.BcryptHash, pw)
	if err != nil {
		fmt.Println("Incorrect master password.")
		return
	}

	// derive key
	salt, _ := DecodeBase64(v.Salt)
	key := DeriveKey(pw, salt)

	for _, e := range v.Entries {
		ct, _ := base64.StdEncoding.DecodeString(e.Password)
		nonce, _ := base64.StdEncoding.DecodeString(e.Nonce)

		pt, _ := DecryptAES(ct, nonce, key)

		fmt.Println("-----")
		fmt.Println("Site:", e.Site)
		fmt.Println("Username:", e.Username)
		fmt.Println("Password:", string(pt))
		fmt.Println("Created:", e.CreatedAt)
	}
}

// --- get command ---

func getEntry(site string) {
	v, _ := LoadVault()
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Master password: ")
	pw, _ := reader.ReadString('\n')
	pw = strings.TrimSpace(pw)

	err := VerifyMasterPassword(v.BcryptHash, pw)
	if err != nil {
		fmt.Println("Incorrect master password.")
		return
	}

	salt, _ := DecodeBase64(v.Salt)
	key := DeriveKey(pw, salt)

	for _, e := range v.Entries {
		if e.Site == site {
			ct, _ := base64.StdEncoding.DecodeString(e.Password)
			nonce, _ := base64.StdEncoding.DecodeString(e.Nonce)

			pt, _ := DecryptAES(ct, nonce, key)

			fmt.Println("Site:", e.Site)
			fmt.Println("Username:", e.Username)
			fmt.Println("Password:", string(pt))
			return
		}
	}

	fmt.Println("Entry not found:", site)
}
