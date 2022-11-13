# Logging

ALWAYS log to `stdout`.

## Logging levels

Panic: The instance closes immediately for security.

Error: The instance is in an invalid state, waiting to be restarted.

Warning: Something is not right, but the instance is still healthy.

Info: Logs normal events without sensitive data.

## Logging sources

Which function is emitting this?

## Logging message

Some useful hints, e.g. `err.Error()`.
