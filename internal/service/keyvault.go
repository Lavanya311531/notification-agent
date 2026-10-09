package service

import "fmt"

func GetSecret(secretName string) string {

	fmt.Println("Fetching secret from Key Vault:", secretName)

	return "dummy-secret"
}
