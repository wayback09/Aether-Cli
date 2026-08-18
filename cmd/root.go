package cmd

import (
	"fmt"
)

func PrintHelp() {
	fmt.Println("Aether CLI - Developer toolkit for Aether extensions")
	fmt.Println("\nUsage:")
	fmt.Println("  aether-cli [command]")
	fmt.Println("\nAvailable Commands:")
	fmt.Println("  init        Scaffolds a new Aether extension project")
	fmt.Println("  validate    Validates the manifest.json of an extension")
	fmt.Println("  build       Builds and packages the extension into an .aex file")
}

func Execute(command string, args []string) error {
	switch command {
	case "init":
		return RunInit(args)
	case "validate":
		return RunValidate(args)
	case "build":
		return RunBuild(args)
	case "help", "--help", "-h":
		PrintHelp()
		return nil
	default:
		return fmt.Errorf("unknown command %q", command)
	}
}
