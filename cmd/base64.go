package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	encodeBase64Command = &cobra.Command{
		Use:   "encode64",
		Short: "Command to encode string to base 64",
		Long:  "Command to encode string to base 64",
		RunE:  encodeBase64Execute,
	}
)

func init() {
	rootCmd.AddCommand(encodeBase64Command)
}

func encodeBase64Execute(cmd *cobra.Command, args []string) error {
	if len(os.Args[2:]) == 0 {
		return errors.New("unexpected args size")
	}

	s := ""

	for i, arg := range os.Args[2:] {
		if i == 0 {
			s = arg
			continue
		}

		s += fmt.Sprintf(" %s", arg)
	}

	println(encodeBase64(s))

	return nil
}

func encodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

var (
	decodeBase64Command = &cobra.Command{
		Use:   "decode64",
		Short: "Command to decode string to base 64",
		Long:  "Command to decode string to base 64",
		RunE:  decodeBase64Execute,
	}
)

func init() {
	rootCmd.AddCommand(decodeBase64Command)
}

func decodeBase64Execute(cmd *cobra.Command, args []string) error {
	if len(os.Args[2:]) == 0 {
		return errors.New("unexpected args size")
	}

	for _, arg := range os.Args[2:] {
		if len(os.Args[2:]) > 1 {
			return errors.New("unexpected args size")
		}

		println(decodeBase64(arg))
	}

	return nil
}

func decodeBase64(s string) string {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println("decode error:", err)
		return ""
	}

	return string(decoded)
}
