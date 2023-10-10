package cmd

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	cnpjCommand = &cobra.Command{
		Use:   "cnpj",
		Short: "Command to generate new random cnpj",
		Long:  "Command to generate new random cnpj (Brazilian Companies Document)",
		RunE:  cnpjExecute,
	}
)

func init() {
	rootCmd.AddCommand(cnpjCommand)
}

func cnpjExecute(cmd *cobra.Command, args []string) error {
	if len(os.Args[2:]) == 0 {
		println(cnpj())
	}

	for _, arg := range os.Args[2:] {
		if len(os.Args[2:]) > 1 {
			return errors.New("unexpected args size")
		}

		q, err := strconv.Atoi(arg)
		if err != nil {
			return err
		}

		for i := 0; i < q; i++ {
			println(cnpj())
		}
	}

	return nil
}

const maxDigit = 9

func cnpj() string {
	n1 := rand.Intn(maxDigit)
	n2 := rand.Intn(maxDigit)
	n3 := rand.Intn(maxDigit)
	n4 := rand.Intn(maxDigit)
	n5 := rand.Intn(maxDigit)
	n6 := rand.Intn(maxDigit)
	n7 := rand.Intn(maxDigit)
	n8 := rand.Intn(maxDigit)
	n9, n10, n11 := 0, 0, 0
	n12 := 1

	d1 := n12*2 + n11*3 + n10*4 + n9*5 + n8*6 + n7*7 + n6*8 + n5*9 + n4*2 + n3*3 + n2*4 + n1*5
	d1 = 11 - (d1 % 11)
	if d1 >= 10 {
		d1 = 0
	}

	d2 := d1*2 + n12*3 + n11*4 + n10*5 + n9*6 + n8*7 + n7*8 + n6*9 + n5*2 + n4*3 + n3*4 + n2*5 + n1*6
	d2 = 11 - (d2 % 11)
	if d2 >= 10 {
		d2 = 0
	}

	return fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d%d%d%d", n1, n2, n3, n4, n5, n6, n7, n8, n9, n10, n11, n12, d1, d2)
}
