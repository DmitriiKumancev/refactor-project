package storage

import (
	"encoding/json"
	"io/fs"
	"os"

	"github.com/DmitriiKumancev/refactor-project/models"
)

const Store = "users.json"

func InitStore() {
	_, err := os.Stat(Store)
	if os.IsNotExist(err) {
		initialStore := models.UserStore{
			Increment: 0,
			List:      make(models.UserList),
		}
		SaveStore(initialStore)
	}
}

func SaveStore(data models.UserStore) error {
	b, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	err = os.WriteFile(Store, b, fs.ModePerm)
	return err
}

func LoadStore() (models.UserStore, error) {
	f, err := os.ReadFile(Store)
	if err != nil {
		return models.UserStore{}, err
	}
	var data models.UserStore
	err = json.Unmarshal(f, &data)
	return data, err
}
