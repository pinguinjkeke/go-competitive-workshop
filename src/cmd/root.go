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
			client := NewPokemonApiClient()

			pokemon, err := client.FindPokemon("bulbasaur")
			if err != nil {
				cmd.Println("Something went wrong during pokemon retrieval")

				return
			}
			cmd.Println(pokemon)

			cmd.Println("Welcome to FlyPoke!!")
		},
	}

	// Register your commands here
	// cmd.AddCommand(ExampleCommand)

	return cmd
}
