package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pavlo-v-chernykh/keystore-go/v4"
)

func ReadKeyStore(filename string, password []byte) (keystore.KeyStore, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	ks := keystore.New()
	if err := ks.Load(f, password); err != nil {
		// log.Fatal(err) // nolint: gocritic
		return keystore.KeyStore{}, err
	}

	return ks, nil
}

func WriteKeyStore(ks keystore.KeyStore, filename string, password []byte) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	err = ks.Store(f, password)
	if err != nil {
		log.Fatal(err) // nolint: gocritic
	}
}

func zeroing(buf []byte) {
	for i := range buf {
		buf[i] = 0
	}
}

func main() {
	passwords := []string{"password1", "password2", "password3", "password4", "password5", "password6"}
	var result keystore.KeyStore = keystore.KeyStore{}

	for i := 0; i < len(passwords); i++ {
		password := []byte(passwords[i])
		defer zeroing(password)
	
		ks, err := ReadKeyStore("./keystores/keystore", password); if err != nil {
			continue
		}
		fmt.Println("Password found: ", password)
		result = ks
	}

	log.Println(result)

	

	// writeKeyStore(ks1, "keystore2.jks", password)

	// ks2 := readKeyStore("keystore2.jks", password)

	// log.Printf("is equal: %v\n", reflect.DeepEqual(ks1, ks2))
}