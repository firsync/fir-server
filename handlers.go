package main

import (
	"crypto/ed25519"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func signupHandler(w http.ResponseWriter, r *http.Request, keys *ED25519Keys, db *sqlx.DB) {
	clientPubKey, _ := hex.DecodeString(keys.publicKey)
	signature := r.Header.Get("Signature")
	signedKey, _ := hex.DecodeString(keys.signedKey)
	if !ed25519.Verify(clientPubKey, signedKey, []byte(signature)) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}
	if err := registerPublicKey(keys, db); err != nil {
		http.Error(w, "Error registering public key: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Public key registered successfully"))
}

func registerPublicKey(keys *ED25519Keys, db *sqlx.DB) error {
	// Verify that the provided public key is valid
	if !ed25519.Verify(ed25519.PublicKey([]byte(keys.publicKey)), []byte(keys.publicKey), []byte(keys.signedKey)) {
		return fmt.Errorf("provided public key is invalid")
	}

	// Check if the public key is already registered
	var id int
	err := db.Get(&id, "SELECT id FROM users WHERE public_key=$1", keys.publicKey)
	if err == nil {
		return fmt.Errorf("public key already registered")
	}
	if err != sql.ErrNoRows {
		return fmt.Errorf("error checking for existing key: %v", err)
	}

	// Add the public key to the list of registered keys
	_, err = db.Exec("INSERT INTO users (public_key) VALUES ($1)", keys.publicKey)
	if err != nil {
		return fmt.Errorf("error registering public key: %v", err)
	}
	return nil
}
