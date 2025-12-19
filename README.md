Project Overview
This project is part of my Moringa School capstone, where the goal was to use Generative AI prompts to explore and learn a new software development technology and document the learning process in a way that others can easily follow.
For this project, I chose to learn Golang (Go) a compiled, statically typed programming language developed by Google and built a simple command-line calculator as a runnable beginner project.

Throughout the project, I used AI prompts via ai.moringaschool.com to:
	•	Understand Golang syntax and concepts
	•	Debug errors and improve code structure
	•	Learn Go best practices such as switch statements
	•	Validate logic and handle edge cases
Project Goals
The main objectives of this project were to:
	•	Learn a new programming language (Go) from a beginner perspective
	•	Use Generative AI prompts to guide learning and problem-solving
	•	Build a simple, runnable application
	•	Clearly document each step so others can replicate the process
	•	Test and iterate the solution through debugging and improvements
Technology Used
	•	Programming Language: Go (Golang)
	•	Development Environment: Visual Studio Code (VS Code)
	•	Operating System: Windows
	•	AI Tool Used: ai.moringaschool.com
	•	Application Type: Command-Line Interface (CLI)
What the Calculator Does
The calculator allows a user to:
	•	Enter two numbers
	•	Choose an arithmetic operator (+, -, *, /)
	•	View the calculated result
	•	Receive a friendly error message if division by zero is attempted
	•	Receive an error message for invalid operators
Project Structure
simple-go-calculator/
│
├── main.go
└── README.md
Setup Instructions (Windows + VS Code)
1.	Install Go (Golang)
2.	Download Go from the official website: https://go.dev/dl/
3.	Run the installer and follow the setup steps
4.	Confirm installation by running in the terminal: go version
Install Visual Studio Code
	1.	Download VS Code from: https://code.visualstudio.com/
	2.	Install the Go extension from the VS Code Extensions Marketplace

Clone or Download the Project
You can either:
	•	Clone the repository, or
	•	Download the files and open the project folder in VS Code

 How to Run the Project
Option 1: Run directly (for testing)
go run main.go
Follow the prompts:
Enter first number:
Enter operator (+, -, *, /):
Enter second number:
Option 2: Build once and run multiple times (recommended)
Build the executable:
go build -o calculator.exe main.go
Run the calculator:
	.\calculator
This allows the program to run without retyping go run main.go every time.
Example Usage:
Enter first number: 10
Enter operator (+, -, *, /): /
Enter second number: 0
Error: Cannot divide by zero!
How Generative AI Was Used

Generative AI played a key role in this project by helping me:
	•	Understand Go syntax and data types
	•	Learn when to use if-else vs switch statements
	•	Debug divide-by-zero errors
	•	Improve code readability and structure
	•	Write beginner-friendly documentation

Example AI prompt used:
“Explain why a switch statement is preferred over multiple if-else conditions in Golang.”
Key Learnings
	•	Golang uses switch statements efficiently for multiple condition checks
	•	Go does not require break in switch cases
	•	Go programs are compiled before execution
	•	Defensive programming (e.g., checking for division by zero) improves reliability
	•	Shell errors (PowerShell) are different from Go runtime errors
Common Beginner Errors & Quick Solutions (Golang)
While building this project, I encountered (and learned to fix) several common errors that beginners using Golang are likely to face. Below are some of these issues and their quick solutions to help others avoid confusion.

1.	Forgetting to Install Go or Set Up PATH
Error: 'go' is not recognized as an internal or external command
Cause:
	•	Go is not installed
	•	OR Go is installed but not added to the system PATH
Quick Solution:
	•	Reinstall Go from https://go.dev/dl/
	•	Restart the computer after installation
	•	Confirm with: go version
Using go run main.go Every Time
Issue:
	•	Repeating the same command becomes tedious
Quick Solution:
Build the program once: go build -o calculator.exe main.go
Then run it anytime using: .\calculator
Division by Zero Errors from the Terminal (PowerShell)

Error Message: Attempted to divide by zero.
Cause:
	•	PowerShell evaluates expressions like 60/0 before Go runs
	•	The Go program never executes
Quick Solution:
	•	Do not pass arithmetic expressions in the terminal
	•	Run the program normally and enter inputs when prompted
Correct way: .\calculator
Expecting break in a Go switch Statement

Mistake: case "+":
    fmt.Println("Addition")
    break
Cause:
	•	Beginners often come from languages like C or Java
Quick Solution:
	•	Go automatically breaks after each case
	•	Remove break
Correct:  case "+":
    fmt.Println("Addition")

Future Improvements
	•	Add a loop to allow continuous calculations in one session
	•	Handle invalid numeric inputs more gracefully
	•	Extend the calculator with additional operations
	•	Add unit tests for arithmetic operations
Conclusion
This project successfully demonstrates how Generative AI can accelerate learning when exploring a new technology. By combining AI guidance with hands-on coding, I was able to learn Golang basics, build a functional application, and document the process for others to follow.
This calculator serves as a beginner-friendly Go toolkit and fulfills the requirements of the Moringa School Generative AI capstone project.
Tips for Quickly Learning Golang (Beginner-Friendly)
As a beginner learning Golang for the first time, I found the following strategies extremely helpful in speeding up my understanding and confidence with the language.

1.	Focus on Core Concepts First

Instead of trying to learn everything at once, I focused on:
	•	Variables and data types
	•	fmt.Print, fmt.Println, and fmt.Printf
	•	Conditional statements (if, switch)
	•	Loops (for)
	•	Functions

These concepts cover most beginner-level Go programs.
2.	Use Go’s Simplicity to Your Advantage
Golang is intentionally simple. I learned faster by:
	•	Avoiding complex libraries early
	•	Writing small, focused programs (like this calculator)
	•	Reading the code out loud to understand the flow

If the code looks simple, it usually is, Go prefers clarity over cleverness.
3.	Learn by Writing, Not Watching
Instead of only watching tutorials:
	•	I typed every line of code myself
	•	I intentionally broke the code to see error messages
	•	I fixed errors using AI prompts and documentation
This hands-on approach helped me understand why things work, not just how.

4.	Use Generative AI as a Learning Assistant
Generative AI significantly accelerated my learning when used correctly.
Examples of effective prompts I used:
	•	“Explain this Go error in simple terms.”
	•	“Why is a switch statement preferred here in Go?”
	•	“Rewrite this Go code in a more beginner-friendly way.”
I treated AI as a mentor, not a code-copying tool.
5.	Read Go Error Messages Carefully
Go error messages are usually clear and helpful.
Tips:
	•	Read the full message
	•	Note the line number
	•	Fix errors one at a time
Understanding errors is one of the fastest ways to learn any language.

6.	Use VS Code Go Tools

.
