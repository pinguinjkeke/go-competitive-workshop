package cmd

type Pokemon struct {
	Name      string
	Abilities []struct {
		Ability struct {
			Name string
		}
	}
	Types []struct {
		Type struct {
			Name string
		}
	}
}
