package utils

import (
	`fmt`
	`log`

	`github.com/google/uuid`
)

func GenerateUuid() (string, error) {
	generatedUuid, err := uuid.NewRandom()
	if err != nil {
		log.Printf( "error in generating the uuid")
		return "", err
	}
	id := fmt.Sprintf("%s", generatedUuid)

	return id, nil
}