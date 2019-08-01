package tenv

import "os"

// Context handles setting test environment variables and restoring
// it's state.
type Context struct {
	backup map[string]*string
}

// backupVar backs up an environment variable to be later restored.
func (c *Context) backupVar(key string) {
	if _, ok := c.backup[key]; ok {
		// exit if already backed up
		return
	}

	if value := os.Getenv(key); value != "" {
		c.backup[key] = &value
		return
	}

	c.backup[key] = nil
}

// Unset unsets a single environment variable.
func (c *Context) Unset(key string) {
	c.backupVar(key)
	os.Unsetenv(key)
}

// Set sets a single environment variable.
func (c *Context) Set(key, value string) {
	c.backupVar(key)
	os.Setenv(key, value)
}

// Restore restores the environment to it's original state.
func (c *Context) Restore() {
	for key, value := range c.backup {
		if value == nil {
			os.Unsetenv(key)
			return
		}

		os.Setenv(key, *value)
	}
}

// New creates a new environment context for managing temporary
// environment variables.
func New() *Context {
	return &Context{
		backup: make(map[string]*string),
	}
}
