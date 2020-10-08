package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/grahamgreen/arthur/auth"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: list,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func list(cmd *cobra.Command, args []string) {
	//b, err := ioutil.ReadFile("creds.json")
	//if err != nil {
	//	log.Fatalf("Unable to read client secret file: %v", err)
	//}
	//// If modifying these scopes, delete your previously saved token.json.
	//config, err := google.ConfigFromJSON(b, tasks.TasksReadonlyScope)
	//if err != nil {
	//	log.Fatalf("Unable to parse client secret file to config: %v", err)
	//}

	client := auth.GetClient()
	ctx := context.Background()
	//srv, err := tasks.New(client)
	srv, err := tasks.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve tasks Client %v", err)
	}
	r, err := srv.Tasklists.List().MaxResults(10).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve task lists. %v", err)
	}
	fmt.Println("Task Lists:")
	if len(r.Items) > 0 {
		for _, i := range r.Items {
			fmt.Printf("%s (%s)\n", i.Title, i.Id)

		}
	} else {
		fmt.Print("No task lists found.")
	}
	t, err := srv.Tasks.List("MDg2MzEzMTEwMzU0NTc5NjM1Mjk6MDow").Do()
	if len(t.Items) > 0 {
		for _, i := range t.Items {
			fmt.Printf("%s (%s)\n", i.Title, i.Notes)

		}
	} else {
		fmt.Print("No task lists found.")
	}

}
