package main

import ihttp "message-center/pkg/http"

func main() {

	if err := run(); err != nil {
		panic(err)
	}

}

func run() error {

	server := ihttp.NewServer()
	if err := server.Run(); err != nil {
		return err
	}
	return nil

}
