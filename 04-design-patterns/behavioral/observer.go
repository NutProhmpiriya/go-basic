// Observer Pattern defines a one-to-many dependency between objects so that when one
// object changes state, all its dependents are notified and updated automatically.
//
// Use cases:
// - When a change to one object requires changing others, and you don't know how many objects need to be changed
// - When an object should be able to notify other objects without making assumptions about who these objects are
// - When you need to maintain consistency between related objects without making them tightly coupled

package behavioral

// Observer interface defines the method that should be implemented by observers
type Observer interface {
	Update(temperature float64)
}

// WeatherStation is the subject that observers are watching
type WeatherStation struct {
	observers   []Observer
	temperature float64
}

// NewWeatherStation creates a new weather station
func NewWeatherStation() *WeatherStation {
	return &WeatherStation{
		observers: make([]Observer, 0),
	}
}

// RegisterObserver adds an observer to the list
func (w *WeatherStation) RegisterObserver(o Observer) {
	w.observers = append(w.observers, o)
}

// RemoveObserver removes an observer from the list
func (w *WeatherStation) RemoveObserver(o Observer) {
	for i, observer := range w.observers {
		if observer == o {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers notifies all observers of the temperature change
func (w *WeatherStation) NotifyObservers() {
	for _, observer := range w.observers {
		observer.Update(w.temperature)
	}
}

// SetTemperature changes the temperature and notifies observers
func (w *WeatherStation) SetTemperature(temp float64) {
	w.temperature = temp
	w.NotifyObservers()
}

// TemperatureDisplay is a concrete observer
type TemperatureDisplay struct {
	name string
}

// NewTemperatureDisplay creates a new temperature display
func NewTemperatureDisplay(name string) *TemperatureDisplay {
	return &TemperatureDisplay{name: name}
}

// Update implements the Observer interface
func (d *TemperatureDisplay) Update(temperature float64) {
	// In a real application, this would update a display
	// For this example, we'll just store the temperature
	d.display(temperature)
}

func (d *TemperatureDisplay) display(temperature float64) string {
	return "Display " + d.name + " shows temperature: " + string(rune(temperature))
}
