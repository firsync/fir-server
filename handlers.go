package main

import (
	"crypto/ed25519"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/jmoiron/sqlx"
)

func signupHandler(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	var jsonData map[string]string
	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	publicKey := jsonData["public_key"]
	signature := jsonData["signature"]

	if err := registerPublicKey(publicKey, signature, db); err != nil {
		http.Error(w, "Error registering public key: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Public key registered successfully"))
}

func registerPublicKey(publicKey string, signature string, db *sqlx.DB) error {
	// Verify that the provided public key is valid
	clientPubKey, _ := hex.DecodeString(publicKey)
	signedKey, _ := hex.DecodeString(signature)
	if !ed25519.Verify(clientPubKey, signedKey, []byte(signature)) {
		return fmt.Errorf("provided public key is invalid")
	}

	// Check if the public key is already registered
	var id int
	err := db.Get(&id, "SELECT id FROM users WHERE public_key=$1", publicKey)
	if err == nil {
		return fmt.Errorf("public key already registered")
	}
	if err != sql.ErrNoRows {
		return fmt.Errorf("error checking for existing key: %v", err)
	}

	// Call the combined script, passing in the public key and home directory as arguments
	createUser := exec.Command("create_user.sh", publicKey, "/home")
	createUser.Run()

	return nil
}
