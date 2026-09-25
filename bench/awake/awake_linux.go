package awake

// hold runs systemd-inhibit with a command that waits forever, so the lock
// lasts until release kills it.
func hold(why string) (func(), error) {
	return holdCommand("systemd-inhibit", "--what=idle:sleep", "--who=multimap bench", "--why="+why,
		"--mode=block", "sleep", "infinity")
}
