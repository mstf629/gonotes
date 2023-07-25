package cmd 

import (
	"os"
  "log"
	"github.com/spf13/cobra"
)

var home = os.Getenv("HOME")
var BasePath = home + "/.cache/gonotes" 
var DbPath = home + "/.cache/gonotes/gonotes.db"


func CheckDb() {

   _ , err := os.Stat(DbPath)
   if os.IsNotExist(err) {
      log.Fatalln("please run [gonotes initdb] to create new database")
   }
}

func CheckErr(errs ...error) {
   for _, err := range errs {
      if err != nil {
         log.Fatalln(err)
      }
   }
}


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gonotes",
	Short: "its a simple program to manage notes",
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
  err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}



func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


