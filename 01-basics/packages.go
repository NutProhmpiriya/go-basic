package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// This file demonstrates the use of various Go standard library packages
func main() {
	fmt.Println("=== String Package Examples ===")
	text := "  Hello, Go Programming!  "
	fmt.Printf("Original: %q\n", text)
	fmt.Printf("Trimmed: %q\n", strings.TrimSpace(text))
	fmt.Printf("Upper: %s\n", strings.ToUpper(text))
	fmt.Printf("Lower: %s\n", strings.ToLower(text))
	fmt.Printf("Contains 'Go': %t\n", strings.Contains(text, "Go"))
	fmt.Printf("Replace: %s\n", strings.Replace(text, "Go", "Golang", 1))

	fmt.Println("\n=== Math Package Examples ===")
	fmt.Printf("Pi: %.5f\n", math.Pi)
	fmt.Printf("Square root of 16: %.2f\n", math.Sqrt(16))
	fmt.Printf("Power 2^3: %.2f\n", math.Pow(2, 3))
	fmt.Printf("Absolute of -42: %d\n", int(math.Abs(-42)))
	fmt.Printf("Max of 10 and 5: %d\n", int(math.Max(10, 5)))

	fmt.Println("\n=== Time Package Examples ===")
	now := time.Now()
	fmt.Printf("Current time: %v\n", now)
	fmt.Printf("Year: %d\n", now.Year())
	fmt.Printf("Month: %s\n", now.Month())
	fmt.Printf("Day: %d\n", now.Day())
	
	// Time formatting
	fmt.Printf("Formatted time: %s\n", now.Format("2006-01-02 15:04:05"))
	
	// Time operations
	tomorrow := now.Add(24 * time.Hour)
	fmt.Printf("Tomorrow: %v\n", tomorrow)
	
	duration := time.Hour * 2 + time.Minute * 30
	fmt.Printf("Duration: %v\n", duration)

	// Note about custom packages:
	fmt.Println("\n=== Custom Package Note ===")
	fmt.Println("To create custom packages:")
	fmt.Println("1. Create a new directory for your package")
	fmt.Println("2. Create .go files with 'package packagename'")
	fmt.Println("3. Initialize module: 'go mod init module-name'")
	fmt.Println("4. Import your package: 'import \"module-name/packagename\"'")
}
