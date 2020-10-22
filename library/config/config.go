package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"
)

type Common struct {
	Port struct {
		HTTP  string
		DEBUG string
		GRPC  string
	}

	Mongodb struct {
		URL string
	}
}

func Load(app string, conf interface{}) (err error) {
	commonPath := "./src/common.toml"

	strs := strings.SplitN(app, "-", 2)
	i := 2
	strs = append(strs[:i], append([]string{"cmd"}, strs[i:]...)...)

	appPath := fmt.Sprintf("./src/%s/%s.toml", strings.Join(strs, "/"), app)

	if _, err = toml.DecodeFile(commonPath, conf); err != nil {
		return
	}

	if _, err = toml.DecodeFile(appPath, conf); err != nil {
		return
	}

	if httpPort := os.Getenv("HTTP_PORT"); httpPort != "" {
		reflect.ValueOf(conf).Elem().FieldByName("Port").FieldByName("HTTP").SetString(httpPort)
	}
	if debugPort := os.Getenv("DEBUG_PORT"); debugPort != "" {
		fmt.Println(debugPort)
		reflect.ValueOf(conf).Elem().FieldByName("Port").FieldByName("DEBUG").SetString(debugPort)
	}
	if grpcPort := os.Getenv("GRPC_PORT"); grpcPort != "" {
		reflect.ValueOf(conf).Elem().FieldByName("Port").FieldByName("GRPC").SetString(grpcPort)
	}

	return
}
