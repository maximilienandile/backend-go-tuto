package uniqueid

import uuid "github.com/satori/go.uuid"

type UUIDV4 struct {
}

func (U UUIDV4) Generate() string {
	return uuid.NewV4().String()
}
