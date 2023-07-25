//the decliration of BasePath DbPath vars in root.go
//the decliration of CheckErr function in root.go 
//the decliration of CheckDb in root.go 
package cmd 

import (
	"fmt"
  "os"
	"github.com/spf13/cobra"
   _ "github.com/mattn/go-sqlite3"
   "database/sql"
   "github.com/olekukonko/tablewriter"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	 Use:   "read",
	 Short: "read notes form database",
	 Long: "read notes from database with class, id and date",
	 Run: func(cmd *cobra.Command, args []string) {
      CheckDb()

      class, _ := cmd.Flags().GetString("class")
      id, _ := cmd.Flags().GetString("id")
      date, _ := cmd.Flags().GetString("date")

      query := GenReadQuery(class, id, date)
      notes := DbReader(query)
      TableWriter(notes)

	 },
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.PersistentFlags().StringP("class","c", "", "set class")
	readCmd.PersistentFlags().StringP("id","", "", "set id")
  readCmd.PersistentFlags().StringP("date","d", "", "set date in format YYYY:MM:DD")
}

func GenReadQuery(class string ,id string, date string) string {
   var query string
   switch {
   case class != "" && id != "" && date != "":
      query = fmt.Sprintf("select * from notes where class='%s' and id='%s' and date='%s';",class, id, date)
   case class != "" && id != "" && date == "":
      query = fmt.Sprintf("select * from notes where class='%s' and id='%s';", class, id)
   case class != "" && id == "" && date != "":
      query = fmt.Sprintf("select * from notes where class='%s' and date='%s';", class, date)
   case class == "" && id != "" && date != "":
      query = fmt.Sprintf("select * from notes where  id='%s' and date='%s';", id, date)
   case class == "" && id == "" && date == "":
      query = fmt.Sprintf("select * from notes") 
   case class == "" && id == "" && date != "":
      query = fmt.Sprintf("select * from notes where date='%s';", date)
   case class == "" && id != "" && date == "":
      query = fmt.Sprintf("select * from notes where id='%s';", id,)
   case class != "" && id == "" && date == "":
      query = fmt.Sprintf("select * from notes where class='%s';", class) 
      
   }
   return query 
}

func DbReader(query string) [][]string{
   db, err := sql.Open("sqlite3", DbPath)
   CheckErr(err)
   defer db.Close()

   rows, err := db.Query(query)
   CheckErr(err)
   defer rows.Close()

   var notes [][]string
   for rows.Next() {
      var note = make([]string, 5, 5)
      err = rows.Scan(&note[0], &note[1], &note[2], &note[3], &note[4])
      CheckErr(err)
      notes = append(notes, note)
   }
   
   return notes
}

func TableWriter(notes [][]string) {
   table := tablewriter.NewWriter(os.Stdout)
   table.SetHeader([]string{"non", "content", "class", "id", "date"})

   for _, note := range notes {
      table.Append(note)
   }

   table.Render()
}
