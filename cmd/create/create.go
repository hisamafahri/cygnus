package create

import (
	"github.com/hisamafahri/cygnus/utils"
	"github.com/spf13/cobra"
)


var CreateCmd = &cobra.Command{
    Use:   "create",
    Short: "Create a new group or user",
    ValidArgs: []string{"user", "group"},
    Args: utils.MatchMultipleArgs(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
    Run: func(cmd *cobra.Command, args []string) {
        if args[0] == "user" {
            createUser()
        } else if args[0] == "group" {
            createGroup()
        } 
    },
}
