package cmd

import (
	"errors"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var (
	cpfCommand = &cobra.Command{
		Use:   "cpf",
		Short: "Command to generate new random cpf",
		Long:  "Command to generate new random cpf (Brazilian Document)",
		RunE:  cpfExecute,
	}
)

func init() {
	rootCmd.AddCommand(cpfCommand)
}

func cpfExecute(cmd *cobra.Command, args []string) error {
	if len(os.Args[2:]) == 0 {
		println(cpf())
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
			println(cpf())
		}
	}

	return nil
}

func randomize(limit int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(limit)
}

func cpf() string {
	n := 9
	n1 := randomize(n)
	n2 := randomize(n)
	n3 := randomize(n)
	n4 := randomize(n)
	n5 := randomize(n)
	n6 := randomize(n)
	n7 := randomize(n)
	n8 := randomize(n)
	n9 := randomize(n)
	d1 := n9*2 + n8*3 + n7*4 + n6*5 + n5*6 + n4*7 + n3*8 + n2*9 + n1*10
	d1 = 11 - (d1 % 11)
	if d1 >= 10 {
		d1 = 0
	}
	d2 := d1*2 + n9*3 + n8*4 + n7*5 + n6*6 + n5*7 + n4*8 + n3*9 + n2*10 + n1*11
	d2 = 11 - (d2 % 11)
	if d2 >= 10 {
		d2 = 0
	}
	_cpf := ""
	_cpf = strconv.Itoa(n1) + strconv.Itoa(n2) + strconv.Itoa(n3) + strconv.Itoa(n4) + strconv.Itoa(n5) + strconv.Itoa(n6) + strconv.Itoa(n7) + strconv.Itoa(n8) + strconv.Itoa(n9) + strconv.Itoa(d1) + strconv.Itoa(d2)
	return _cpf
}
