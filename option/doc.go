/*
Package option provides an implementation of the Option monad, which is used to represent the presence (Some) or absence (None) of a value.

The Option type is useful for computations or operations where a value may or may not exist, eliminating the need for null values or panics. It helps ensure safe handling of optional values without having to check explicitly for null or use panics.

Option consists of:

	Some: Represents the presence of a value.
	None: Represents the absence of a value.

Usage Example:

	func findUser(id int) Option[User] {
	  user := db.GetUserByID(id)
		return option.Wrap(user)
	}

	func doSomething() {
	  user := findUser(1)
		user.Match(
	     func(user User) {},
	     func() { // No user found }
		)
	}

In this example, Some wraps the found user, and None represents the case where no user was found.
*/
package option
