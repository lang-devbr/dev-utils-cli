package cmd

import (
	"errors"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	uuidCommand = &cobra.Command{
		Use:   "uuid",
		Short: "Command to generate new uuid",
		Long:  "Command to generate new uuid",
		RunE:  uuidExecute,
	}
)

func init() {
	rootCmd.AddCommand(uuidCommand)
}

func uuidExecute(cmd *cobra.Command, args []string) error {
	if len(os.Args[2:]) == 0 {
		println(uuidNew())
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
			println(uuidNew())
		}
	}

	return nil
}

func uuidNew() string {
	return uuid.NewString()
}
