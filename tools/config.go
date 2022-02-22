package tools

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func Config() map[string]string {
	file, err := os.Open("./ini")
	config := map[string]string{}
	if err != nil {
		fmt.Println("cannot find ini file!")
		return config
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		conf := strings.Split(scanner.Text(), "=")
		config[conf[0]] = conf[1]
	}
	return config
}
