This is a simple CLI program for managing your tasks!

Is is completely made in Go, with no frameworks or external tools (besides built-in libraries, of course)

I am a begginer in writing Go code, so you can make fun of me as you wish

# How to use

There are many arguments that can be passed to the program's parameters, and that's how the program is based on.

Tasklist-CLI:
    
    add [taskName] [taskDescription] [taskImportance] -- creates a new task
    delete [taskName] -- deletes the task with the given identifier (name)
    list -- show all current tasks
    changestat [taskname] [status] -- changes the status of the task with the given identifier