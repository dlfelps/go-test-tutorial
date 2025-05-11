package cmd

import (
        "fmt"
        "strconv"
        "strings"

        "go-test-tutorial/internal/tutorial"
        "go-test-tutorial/internal/utils"

        "github.com/spf13/cobra"
)

var topic string
var interactive bool

var learnCmd = &cobra.Command{
        Use:   "learn [topic number]",
        Short: "Learn about Go testing concepts",
        Long: `Learn about Go testing concepts with explanations and examples.
You can specify a topic number or run in interactive mode to go through all topics.`,
        Run: func(cmd *cobra.Command, args []string) {
                if len(args) > 0 {
                        topicNum, err := strconv.Atoi(args[0])
                        if err != nil {
                                fmt.Printf("Invalid topic number: %s\n", args[0])
                                listTopics()
                                return
                        }
                        showTopic(topicNum)
                } else if interactive {
                        startInteractiveTutorial()
                } else {
                        listTopics()
                }
        },
}

func init() {
        learnCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Start an interactive tutorial")
}

func listTopics() {
        fmt.Println("Available topics to learn:")
        for i, topic := range tutorial.GetAllConcepts() {
                fmt.Printf("%d. %s\n", i+1, topic.Title)
        }
        fmt.Println("\nUse 'gotest-learn learn <topic number>' to learn about a specific topic")
        fmt.Println("Use 'gotest-learn learn --interactive' to start an interactive tutorial")
}

func showTopic(topicNum int) {
        concepts := tutorial.GetAllConcepts()
        if topicNum < 1 || topicNum > len(concepts) {
                fmt.Printf("Topic number %d is out of range\n", topicNum)
                listTopics()
                return
        }

        concept := concepts[topicNum-1]
        fmt.Printf("\n=== %s ===\n\n", concept.Title)
        fmt.Println(concept.Description)
        
        if concept.Example != "" {
                fmt.Println("\nExample:")
                fmt.Println("```go")
                fmt.Println(concept.Example)
                fmt.Println("```")
        }
        
        if concept.PracticalExercise != "" {
                fmt.Println("\nPractical Exercise:")
                fmt.Println(concept.PracticalExercise)
        }
}

func startInteractiveTutorial() {
        concepts := tutorial.GetAllConcepts()
        currentTopic := 0

        for {
                utils.ClearScreen()
                showTopic(currentTopic + 1)
                
                if currentTopic == len(concepts)-1 {
                        fmt.Println("\nYou've completed all topics!")
                        fmt.Println("Press Enter to exit, or 'p' to go back to the previous topic...")
                } else {
                        fmt.Println("\nPress Enter to continue to the next topic, 'p' for previous topic, or 'q' to quit...")
                }
                
                var input string
                fmt.Scanln(&input)
                input = strings.ToLower(input)
                
                if input == "q" {
                        break
                } else if input == "p" && currentTopic > 0 {
                        currentTopic--
                } else if input != "p" && currentTopic < len(concepts)-1 {
                        currentTopic++
                } else if input != "p" && currentTopic == len(concepts)-1 {
                        break
                }
        }
}
