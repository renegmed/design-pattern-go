package main

func main() {
	r := NewRunway()

	h := NewHelicopter("Jack")
	p := NewPlane("Tom")
	u := NewUFO("ET")

	r.Join(h)
	r.Join(p)
	r.Join(u)

	h.Say("Trying to land on Runway")
	p.Say("Acknowledged")
	u.Say("Phone home")

	h.PrivateMessage(p, "Is there an UFO in the Runway")
	p.PrivateMessage(h, "Dude, there's an UFO in the Runway")

	u.Say("I shouldn't be able to read your PM's but I can... ")

}
