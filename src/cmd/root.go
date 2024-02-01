package cmd

import (
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flypoke",
		Short: "A basic CLI example",
		Long:  "A basic CLI example using Cobra",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	pokemonCommand := &cobra.Command{
		Use:   "pokemon",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Println("Please enter a pokemon name")
			}

			client := NewPokemonApiClient()

			if args[0] == "pokemon" {
				client.FindPokemon(args[0])
			}
		},
	}

	attackCommand := &cobra.Command{
		Use:   "attack",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	statsCommand := &cobra.Command{
		Use:   "stats",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	// Register your commands here
	// cmd.AddCommand(ExampleCommand)

	cmd.AddCommand(pokemonCommand)
	cmd.AddCommand(attackCommand)
	cmd.AddCommand(statsCommand)

	return cmd
}
