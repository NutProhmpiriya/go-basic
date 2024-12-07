package main

import (
	"fmt"
	"github.com/your-username/golang-basic/04-design-patterns/behavioral"
	"github.com/your-username/golang-basic/04-design-patterns/creational"
	"github.com/your-username/golang-basic/04-design-patterns/structural"
)

func main() {
	// Creational Patterns

	// 1. Singleton
	fmt.Println("=== Singleton Pattern ===")
	singleton1 := creational.GetInstance()
	singleton2 := creational.GetInstance()
	
	singleton1.IncrementCount()
	fmt.Printf("Singleton1 count: %d\n", singleton1.GetCount())
	fmt.Printf("Singleton2 count: %d\n", singleton2.GetCount())
	fmt.Println()

	// 2. Factory
	fmt.Println("=== Factory Pattern ===")
	creditCard := creational.PaymentFactory(creational.CreditCardType)
	paypal := creational.PaymentFactory(creational.PayPalType)
	
	fmt.Println(creditCard.Pay(100.0))
	fmt.Println(paypal.Pay(50.0))
	fmt.Println()

	// 3. Builder
	fmt.Println("=== Builder Pattern ===")
	builder := creational.NewComputerBuilder()
	director := creational.NewDirector(builder)
	
	gamingPC := director.BuildGamingPC()
	officePC := director.BuildOfficePC()
	
	fmt.Printf("Gaming PC: %+v\n", gamingPC)
	fmt.Printf("Office PC: %+v\n", officePC)
	fmt.Println()

	// Structural Patterns

	// 4. Adapter
	fmt.Println("=== Adapter Pattern ===")
	adapter := structural.NewAdapter()
	fmt.Println(adapter.Request())
	fmt.Println()

	// 5. Decorator
	fmt.Println("=== Decorator Pattern ===")
	coffee := &structural.SimpleCoffee{}
	coffeeWithMilk := structural.NewMilkDecorator(coffee)
	coffeeWithMilkAndSugar := structural.NewSugarDecorator(coffeeWithMilk)
	
	fmt.Printf("Cost: %.2f, Description: %s\n", 
		coffeeWithMilkAndSugar.GetCost(), 
		coffeeWithMilkAndSugar.GetDescription())
	fmt.Println()

	// 6. Facade
	fmt.Println("=== Facade Pattern ===")
	computer := structural.NewComputerFacade()
	startupSteps := computer.Start()
	for _, step := range startupSteps {
		fmt.Println(step)
	}
	fmt.Println()

	// Behavioral Patterns

	// 7. Observer
	fmt.Println("=== Observer Pattern ===")
	weatherStation := behavioral.NewWeatherStation()
	display1 := behavioral.NewTemperatureDisplay("Display 1")
	display2 := behavioral.NewTemperatureDisplay("Display 2")
	
	weatherStation.RegisterObserver(display1)
	weatherStation.RegisterObserver(display2)
	weatherStation.SetTemperature(25.0)
	fmt.Println()

	// 8. Strategy
	fmt.Println("=== Strategy Pattern ===")
	cart := behavioral.NewShoppingCart(behavioral.NewCreditCardStrategy("1234", "123"))
	fmt.Println(cart.Checkout(100.0))
	
	cart.SetPaymentStrategy(behavioral.NewPayPalStrategy("test@test.com", "password"))
	fmt.Println(cart.Checkout(50.0))
	fmt.Println()

	// 9. Chain of Responsibility
	fmt.Println("=== Chain of Responsibility Pattern ===")
	loggerChain := behavioral.NewLoggerChain()
	
	fmt.Println(loggerChain.Log(behavioral.LogEntry{
		Message: "This is an information.",
		Level:   behavioral.INFO,
	}))
	
	fmt.Println(loggerChain.Log(behavioral.LogEntry{
		Message: "This is a debug information.",
		Level:   behavioral.DEBUG,
	}))
	
	fmt.Println(loggerChain.Log(behavioral.LogEntry{
		Message: "This is an error information.",
		Level:   behavioral.ERROR,
	}))
}
