package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var (
	zipCodeCommand = &cobra.Command{
		Use:   "zipcode",
		Short: "Command to get given zipcode",
		Long:  "Command to get given zipcode (service used: viacep.com.br)",
		RunE:  zipCodeExecute,
	}
)

func init() {
	rootCmd.AddCommand(zipCodeCommand)
}

type zipCodeResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func zipCodeExecute(cmd *cobra.Command, args []string) error {
	for _, arg := range os.Args[2:] {
		if len(os.Args[2:]) > 1 {
			return errors.New("unexpected args size")
		}

		req, err := http.Get("https://viacep.com.br/ws/" + arg + "/json/")
		if err != nil {
			log.Println(err)
			return err
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
			return err
		}

		var z zipCodeResponse
		err = json.Unmarshal(res, &z)
		if err != nil {
			log.Println(err)
			return err
		}

		println(fmt.Sprintf("CEP: %s", z.Cep))
		println(fmt.Sprintf("Logradouro: %s", z.Logradouro))
		println(fmt.Sprintf("Complemento: %s", z.Complemento))
		println(fmt.Sprintf("Bairro: %s", z.Bairro))
		println(fmt.Sprintf("Localidade: %s", z.Localidade))
		println(fmt.Sprintf("Uf: %s", z.Uf))
		println(fmt.Sprintf("Ddd: %s", z.Ddd))
	}

	return nil
}
