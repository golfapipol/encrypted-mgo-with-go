package model

import (
	"encryptedmgo/util"

	"gopkg.in/mgo.v2/bson"
)

var (
	key = "32 bytes secret key for aes CFB!"
)

type EncryptedString string

func (e EncryptedString) GetBSON() (interface{}, error) {
	return util.Encrypt(key, string(e))
}

func (e *EncryptedString) SetBSON(raw bson.Raw) error {
	var str string
	raw.Unmarshal(&str)
	decrypted, err := util.Decrypt(key, str)
	if err != nil {
		return err
	}

	*e = EncryptedString(decrypted)
	return nil
}

// try mark
// func (e EncryptedString) String() string {
// 	return fmt.Sprintf("%s", e)
// }
