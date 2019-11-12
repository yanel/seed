package script

//Error is a Go call error.
type Error struct {
	String
}

//Promise represents a future action that can either succeed or fail.
type Promise struct {
	Native
	q Ctx
}

//Raw returns the raw JS promise.
func (promise Promise) Raw() string {
	return promise.LanguageType().Raw()
}

//Then executes the provided function when the promise succeeds.
func (promise Promise) Then(f func(value Dynamic)) Promise {
	promise.q.Javascript(promise.Raw() + ` = ` + promise.Raw() + ".then(function(promise_result) {")
	f(promise.q.Value("promise_result").Dynamic())
	promise.q.Javascript("});")
	return promise
}

//Catch executes the provided function when the promise fails.
func (promise Promise) Catch(f func(err Error)) Promise {
	promise.q.Javascript(promise.Raw() + ".catch(function(rpc_result) {")
	f(Error{promise.q.Value("rpc_result.response").String()})
	promise.q.Javascript("});")
	return promise
}