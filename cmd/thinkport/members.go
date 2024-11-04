package thinkport

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

type Member struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Position string `json:"position"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Linkedin string `json:"linkedin"`
}

type Members []Member

// membersCmd represents the members command
var membersCmd = &cobra.Command{
	Use:     "members",
	Example: "thinkport members [--search | -s <name>]",
	Aliases: []string{"member", "m", "mem", "mebr"},
	Version: version,
	Short:   "List all members of the thinkport team",
	Long: `Get informations about the thinkport team. For example:
thinkport members - Get all members of the thinkport team`,
	Run: func(cmd *cobra.Command, args []string) {
		// If there is no flag, get all members
		if !cmd.Flags().Changed("search") {
			getMembers()
		} else {
			getMember(cmd.Flag("search").Value.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(membersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	membersCmd.PersistentFlags().String("search", "s", "Search for a member")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// membersCmd.Flags().BoolP("search", "s", false, "Search for a member")
}

// Get members from the thinkport team from a REST API apiURL + /members
type MemberWrapper struct {
	Members []Member `json:"members"`
}

// Get a member from the thinkport team from a REST API apiURL + /members/{name}
func getMember(name string) {
	var member Member
	// Get member from REST API
	err := GetJSON(apiURL+"/member/"+name, &member)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Check if member is empty
	if member.Name == "" {
		color.Red("Member not found")
		return
	}

	// Print member in a table
	printMembers(Members{member})

}

func getMembers() {
	var members MemberWrapper
	// Get members from REST API
	err := GetJSON(apiURL+"/members", &members)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Print members
	printMembers(members.Members)
}

// Print members in a table
func printMembers(members Members) {
	// Print members
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Name", "Position", "Email")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	//// Add rows
	for _, member := range members {
		tbl.AddRow(member.Name+" "+member.Surname, member.Position, member.Email)
	}

	//// Print table
	tbl.Print()
}

func generateDocumentation()  {
	err := doc.GenMarkdownTree(membersCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
	docHeader := &doc.GenManHeader{
		Title: "Thinkport CLI",
	}
	err = doc.GenManTree(membersCmd, docHeader, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
