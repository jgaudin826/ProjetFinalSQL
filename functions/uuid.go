package functions

import (
	"log"

	"github.com/gofrs/uuid"
)

//UUID
/*
! GetNewUuid uses the gofrs/uuid package to generate a random uuid to be assigned to the new users and posts created and stored in the database
*/
func GetNewUuid() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	return uuid.String()
}
