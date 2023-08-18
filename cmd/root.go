/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "simpleGoServer",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		port, _ := cmd.Flags().GetString("portNum")
		directory, _ := cmd.Flags().GetString("directory")
		// curWD, err := os.Getwd()
		// log.Printf("type of wd: %T", curWD)
		// port := flag.String("p", "8080", "port to serve on")
		// directory := flag.String("d", curWD, "the directory of the static files to serve")
		// flag.Parse()

		http.Handle("/", loggingHandler(http.FileServer(http.Dir(directory))))

		log.Printf("Serving %s on HTTP port: %s\n", directory, port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	},
}

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.simpleGoServer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	var portNum string
	var directory string

	curWD, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&portNum, "portNum", "p", "8080", "Port Number to use")
	rootCmd.Flags().StringVarP(&directory, "directory", "d", curWD, "Directory to serve")
}
