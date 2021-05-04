package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/oliveagle/jsonpath"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("not given expensive")
		return
	}

	exp := os.Args[1]

	var err error
	var content []byte
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalf("stdin stat err %s", err)
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		log.Fatal("no file and no pipe")
	} else {
		content, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("read content error %s", err)
		}
	}

	var json_data interface{}
	json.Unmarshal(content, &json_data)
	res, err := jsonpath.JsonPathLookup(json_data, exp)
	if err != nil {
		log.Fatalf("no file and no pipe %s", err)
		return
	}

	m, err := json.MarshalIndent(res, "", "   ")
	if err != nil {
		log.Fatalf("marshal response error %s", err)
	}

	fmt.Println(string(m))

}
