package invoice

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type seller struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Nip     string `yaml:"nip"`
	Account string `yaml:"account"`
	Bank    string `yaml:"bank,omitempty"`
	Phone   string `yaml:"phone,omitempty"`
	Team    string `yaml:"team,omitempty"`
}

type buyer struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Nip     string `yaml:"nip"`
	// email   string `yaml:"email"`
}

// type vat struct{}
// type email struct {}

type item struct {
	Name     string  `yaml:"name"`
	Quantity int     `yaml:"quantity"`
	Netto    float64 `yaml:"netto"`
	Vat      string  `yaml:"vat"`
	Email    string  `yaml:"email"`
}

type invoice struct {
	Buyer  buyer  `yaml:"buyer,flow"`
	Seller seller `yaml:"seller,flow"`
	Items  []item `yaml:"items,flow"`
}

// Config used for keeping information about all the invoices
type Config struct {
	Invoices []invoice `yaml:"invoices,flow"`
}

func (c *Config) readFromFile(path string) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Error while reading the file %v: %v", path, err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Error while unmarshaling: %v", err)
	}
}

// NewConfigFromFile create a new Config object from yaml file.
func NewConfigFromFile(path string) Config {
	c := Config{}
	c.readFromFile(path)
	return c
}
