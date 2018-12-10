package main

import "io/ioutil"

type fileSystem interface {
	readFile(filename string) ([]byte, error)
}

type osFileSystem struct {
}

func (osFileSystem) readFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
