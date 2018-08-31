package pcidb

import (
	"net/http"
	"io/ioutil"
	"os"
	"fmt"
)

var path_to_pcidb = os.Getenv("LOCALAPPDATA") + `\pcidb\`
var filename = `\pciids.txt`

func getAndSave(){
	response, err := http.Get("http://pci-ids.ucw.cz/v2.2/pci.ids")
	bytes, _ := ioutil.ReadAll(response.Body)

	if err != nil {
		panic("cannot get device id list")
	}
	
	defer response.Body.Close()

	createPCIDBDirectory()
	createPCIIDTxtFile(bytes)

}

func createPCIDBDirectory() {
	if _, err := os.Stat(path_to_pcidb); os.IsNotExist(err){
		os.Mkdir(path_to_pcidb, os.ModePerm)
	}
}

func createPCIIDTxtFile(bytes []byte) {
	file, fileErr := os.Create(path_to_pcidb + filename)
	fmt.Fprintf(file, "%v\n", string(bytes))
}