package cmd

import (
	"edx_go_2/src/config"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var cmdRoot = &cobra.Command{
	Use: "edx_go",
}

func migrateCreate(dbUrl string) *cobra.Command {
	return &cobra.Command{
		Use: "migrate-create",
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]

			out, err := exec.Command("/bin/sh", "-c", "cd src/migrations; goose create "+name+" sql;").Output()

			if err != nil {
				log.Fatalln("migrate error", err)
			}

			log.Println("migrate successfully", string(out))
		},
	}
}

func migrateDB(dbUrl string) *cobra.Command {
	return &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {

			query := fmt.Sprintf("cd src/migrations; goose mysql %q up;", dbUrl)

			out, err := exec.Command("/bin/sh", "-c", query).Output()

			if err != nil {
				log.Fatalln("migrate up error", err)
			}

			log.Println("migrate up successfully", string(out))

		},
	}
}

func migrateDown(dbUrl string) *cobra.Command {
	return &cobra.Command{
		Use: "migrate-down",
		Run: func(cmd *cobra.Command, args []string) {
			query := fmt.Sprintf("cd src/migrations; goose mysql %q down;", dbUrl)

			out, err := exec.Command("/bin/sh", "-c", query).Output()

			if err != nil {
				log.Fatalln("migrate down error", err)
			}

			log.Println("migrate down successfully", string(out))

		},
	}
}

func GetRoot(cnf config.Config) *cobra.Command {
	cmdRoot.AddCommand(migrateCreate(cnf.DbUrl))
	cmdRoot.AddCommand(migrateDB(cnf.DbUrl))
	cmdRoot.AddCommand(migrateDown(cnf.DbUrl))
	return cmdRoot
}
