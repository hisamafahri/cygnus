package add

import (
	"github.com/hisamafahri/cygnus/utils"
	"github.com/spf13/cobra"
)


var AddCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a file or user to an existing group",
    ValidArgs: []string{"user", "file"},
    Args: utils.MatchMultipleArgs(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
    Run: func(cmd *cobra.Command, args []string) {
        if args[0] == "user" {
            addUser()
        } else if args[0] == "file" {
            addFile()
        } 
    },
}
