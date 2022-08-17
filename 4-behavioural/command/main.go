package main

import "fmt"

// Command and Strategy may look similar because you can use both to parameterize an object with some action. However, they have very different intents.
// You can use Command to convert any operation into an object. The operation’s parameters become fields of that object.
// The conversion lets you defer execution of the operation, queue it, store the history of commands, send commands to remote services, etc.
// On the other hand, Strategy usually describes different ways of doing the same thing, letting you swap these algorithms within a single context class.

// Client
type Body struct {
	HasHeadache    bool
	HasStomachAche bool
}

// Commands
// ========
// a Command must provide a parameterless api for the CommandContainer to call
// a Command that requires data should be pre-constructed outside the CommandContainer or should be able to acquire data without
// requiring the CommandContainer to provide data to it (e.g. env vars, globally accessible vars, restapi)

type IMedicate interface {
	InitiateEffect()
}

type MigraineMedicine struct {
	Body *Body
}

func (f *MigraineMedicine) InitiateEffect() {
	fmt.Printf("Has headache: %s\n", fmt.Sprint(f.Body.HasHeadache))
	fmt.Println("MigraineMedicine taking effect!")
	f.Body.HasHeadache = false
	fmt.Printf("Has headache: %s\n", fmt.Sprint(f.Body.HasHeadache))
}

type StomachAcheMedicine struct {
	Body *Body
}

func (f *StomachAcheMedicine) InitiateEffect() {
	fmt.Printf("Has stomachAche: %s\n", fmt.Sprint(f.Body.HasStomachAche))
	fmt.Println("StomachAcheMedicine taking effect!")
	f.Body.HasStomachAche = false
	fmt.Printf("Has stomachAche: %s\n", fmt.Sprint(f.Body.HasStomachAche))
}

// Command Container
type Capsule struct {
	MedicineSubstance IMedicate
}

func (c *Capsule) Consumed() {
	c.MedicineSubstance.InitiateEffect()
}

func main() {
	john := &Body{HasHeadache: true}
	migraineCure := &MigraineMedicine{Body: john}
	migraineCapsule := &Capsule{MedicineSubstance: migraineCure}
	migraineCapsule.Consumed()

	sam := &Body{HasStomachAche: true}
	stomachAcheCure := &StomachAcheMedicine{Body: sam}
	stomachAcheCapsule := &Capsule{MedicineSubstance: stomachAcheCure}
	stomachAcheCapsule.Consumed()
}
