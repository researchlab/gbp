package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

var Data *Config

func init() {
	data, err := os.ReadFile("./env")
	if err != nil {
		panic(fmt.Sprintf("config read failed", err))
	}

	err = json.Unmarshal(data, &Data)
	if err != nil {
		panic(fmt.Sprintf("config decode failed", err))
	}
	validate()
}

type Config struct {
	Addr string
	Usr  string
	Pwd  string
}

func validate(){
	if Data.Addr == "" {
			panic("config item addr invalid")
	}
	if Data.Usr == "" {
		panic("config item usr invalid")
	}
	if Data.Pwd == "" {
		panic("config item pwd invalid")
	}
}
