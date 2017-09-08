package archivebuffer

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/microwaves/go-utils/randomizer"
)

func TestArchive(t *testing.T) {
	testfilePath, err := createTestFile()
	if err != nil {
		t.Errorf("Some errors occurred when creating the test file. %v", err)
	}
	tarBuf, err := NewTarballBuffer(testfilePath)
	if err != nil {
		t.Errorf("Some errors occurred when archiving. %v", err)
	}
	unarchivePath := "/tmp/unarchiving-test"
	os.Mkdir("/tmp/unarchiving-test", 0777)
	err = UntarToFile(tarBuf, unarchivePath)
	if err != nil {
		t.Errorf("Some errors occurred when unarchiving. %v", err)
	}
	want := "hello, foobar!\n"
	f := strings.Split(testfilePath, "/")
	got, err := ioutil.ReadFile(unarchivePath + "/" + f[len(f)-1])
	if err != nil {
		t.Errorf("Some errors occurred when reading the file. %v", err)
	}
	if string(got) != want {
		t.Errorf("Expected '%v' got '%v'", want, got)
	}
}

func createTestFile() (string, error) {
	path := fmt.Sprintf("/tmp/foobar-%v", randomizer.GenerateRandomString(8))
	d := []byte("hello, foobar!\n")
	err := ioutil.WriteFile(path, d, 0644)
	if err != nil {
		return "", err
	}
	return path, nil
}
