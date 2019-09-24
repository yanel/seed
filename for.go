package seed

import "github.com/qlova/seed/user"
import qlova "github.com/qlova/script"

//Go runs a Go function instead of a script Function.
func Go(function interface{}, args ...qlova.Type) func(q Script) {
	return func(q Script) {
		q.Go(function)
	}
}

//Update is a remote update to a seed.
type Update struct {
	id string
	user.User
}

//For returns a new remote update to a seed that can be used to remotely modify the seed.
func (seed Seed) For(u User) Update {
	return Update{seed.id, u}
}

//SetText sets the seed's text.
func (update Update) SetText(text string) {
	update.Document["#"+update.id+".innerText"] = text
}
