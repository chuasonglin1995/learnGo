package main
import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal      uint16 // min 0 max 65535
	brake_pedal    uint16
	steering_wheel int16 // -32k - +32k
	top_speed_kmh  float64
}

// Value receiver
// - operates on a copy of the original type's value. 
// - When the method is invoked, a copy of the value is passed to the method. 
// - Changes made in the method will not be reflected in the original value.

// Eg. Does not change the original value
/*func (m MyType) SetValue(val int) {
    m.Value = val
	}

func main() {
    m := MyType{Value: 0}
    m.SetValue(5)
    fmt.Println(m.Value) // prints "0", not "5"
}*/

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax/kmh_multiple)
}

// Pointer receiver
// - If struct is small, or you want to work woth a copy of the value, use a value receiver 
// - If struct is large, use a pointer receiver to avoid copying the entire struct
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main() {
	a_car := car{
		gas_pedal: 22341,
		brake_pedal:    0,
		steering_wheel: 12561,
		top_speed_kmh:  225.0,
	}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
}