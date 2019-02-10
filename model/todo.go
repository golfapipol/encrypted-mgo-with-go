package model

type Todo struct {
	Name          string          `bson:"name"`
	EncryptedName EncryptedString `bson:"encrypted_name"`
}
