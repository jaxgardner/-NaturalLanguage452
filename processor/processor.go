package processor

import (
	"bufio"
	"fmt"
	"natural-language-app/chatgpt"
	"natural-language-app/database"
	"os"
	"strings"
)

var setupSqlScript = database.CreateDepartmentTableSQL + "\n" +
	database.CreateInstructorTableSQL + "\n" +
	database.CreateCoursesTableSQL + "\n" +
	database.CreateInstructorTableSQL + "\n" +
	database.CreateStudentsTableSQL + "\n"

var userInput string
var commonStrategy = "Give me a postgres select statement that answers the question. Only respond with postgres sql syntax. If there is an error do not expalin it!\n"
var strategies = map[string]string{
	"zero_shot": setupSqlScript + commonStrategy,
	"single_domain_double_shot": (setupSqlScript +
		" Who are all the students enrolled in CS101?\n" +
		`SELECT s.first_name, s.last_name, c.name AS course_name
		FROM students s
		JOIN enrollments e ON s.id = e.student_id
		JOIN courses c ON e.course_id = c.id
		WHERE c.code = CS101;\n` +
		commonStrategy),
}

func Process() {
	for userInput != "quit" {
		fmt.Print("Enter input (type 'quit' to exit)\n")
		if userInput != "quit" {
			AskAboutDatabase()
		}
	}
}

func AskAboutDatabase() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please ask about the database: ")
	userInput, _ = reader.ReadString('\n')
	fmt.Println("Would you like to use strategy 1 or 2?")
	strategy, _ := reader.ReadString('\n')

	strategy = strings.TrimSpace(strategy)

	question := "Question: " + userInput

	sqlStatement := ""

	if strategy == "1" {
		sqlStatement = chatgpt.GetChatResponse(strategies["zero_shot"] + question)
	} else if strategy == "2" {
		sqlStatement = chatgpt.GetChatResponse(strategies["single_domain_double_shot"] + question)
	}

	sqlResult := ExecuteSQL(sqlStatement)

	beforeStatement := `
	The original question was:` + userInput +
		`You gave me the following SQL statement: ` + sqlStatement +
		`This is the result from the sql statement executed, please turn it into a natural response to the question: \n`

	naturalLanguageResponse := chatgpt.GetChatResponse(beforeStatement + sqlResult)

	fmt.Println(naturalLanguageResponse)
}

func ExecuteSQL(sqlStatement string) string {
	db := database.GetDB()
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("Error executing SQL statement: ", err)
	}
	defer rows.Close()

	colNames, err := rows.Columns()
	if err != nil {
		fmt.Println("Error getting columns: ", err)
	}

	results := make([]map[string]interface{}, 0)
	for rows.Next() {
		columnValues := make([]interface{}, len(colNames))
		columnPointers := make([]interface{}, len(colNames))
		for i := range columnValues {
			columnPointers[i] = &columnValues[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			fmt.Println("Error scanning row: ", err)
		}

		rowMap := make(map[string]interface{})
		for i, colName := range colNames {
			rowMap[colName] = columnValues[i]
		}

		results = append(results, rowMap)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error with rows: ", err)
	}

	finalResult := ""
	for _, row := range results {
		finalResult += fmt.Sprintf("%v\n", row)
	}

	return finalResult
}
