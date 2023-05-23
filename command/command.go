package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin/v2"
)

var app = kingpin.New("App", "Simple app")

var commandSomething = app.Command("something", "do something")
var commandSomethingArgX = commandSomething.Flag("x", "arg x").String()
var commandSomethingFlagY = commandSomething.Flag("y", "flag y").String()

var commandAdd = app.Command("add", "add new user")
var commandAddFlagOverride = commandAdd.Flag("override", "override existing user").
	Short('o').Bool()
var commandAddArgUser = commandAdd.Arg("user", "username").Required().String()

var commandUpdate = app.Command("update", "update user")
var commandUpdateArgOldUser = commandUpdate.Arg("old", "old username").Required().
	String()
var commandUpdateArgNewUser = commandUpdate.Arg("new", "new username").Required().
	String()

var commandDelete = app.Command("delete", "delete user")
var commandDeleteFlagForce = commandDelete.Flag("force", "force deletion").
	Short('f').Bool()
var commandDeleteArgUser = commandDelete.Arg("user", "username").Required().
	String()

func main() {
	commandAdd.Action(func(ctx *kingpin.ParseContext) error {
		user := *commandAddArgUser
		override := *commandAddFlagOverride
		fmt.Printf("adding user %s, override %t \n", user, override)
		return nil
	})

	commandUpdate.Action(func(ctx *kingpin.ParseContext) error {
		oldUser := *commandUpdateArgOldUser
		newUser := *commandUpdateArgNewUser
		fmt.Printf("updateing user from %s %s \n", oldUser, newUser)

		return nil
	})

	commandDelete.Action(func(ctx *kingpin.ParseContext) error {
		user := *commandDeleteArgUser
		force := *commandDeleteFlagForce
		fmt.Printf("deleting user %s, force %t \n", user, force)

		return nil
	})
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
