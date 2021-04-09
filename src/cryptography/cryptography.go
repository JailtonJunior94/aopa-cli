package cryptography

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jailtonjunior94/aopa-cli/src/infrastructure"

	"github.com/urfave/cli/v2"
)

func Encrypt(c *cli.Context) error {
	password := c.String("password")
	text := c.String("text")

	url := fmt.Sprintf("%s/api/v1/criptografias/%s/criptografar", infrastructure.BaseURL, password)
	res, err := http.Post(url, "text/plain", bytes.NewBuffer([]byte(text)))
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return err
	}

	fmt.Printf("SUCCESS: %s", string(body))
	return nil
}

func Decrypt(c *cli.Context) error {
	password := c.String("password")
	text := c.String("text")

	url := fmt.Sprintf("%s/api/v1/criptografias/%s/descriptografar", infrastructure.BaseURL, password)
	res, err := http.Post(url, "text/plain", bytes.NewBuffer([]byte(text)))
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return err
	}

	fmt.Printf("SUCCESS: %s", string(body))
	return nil
}
