# Taskmanatee

Taskmanatee is currently an early stage interactive shell script that allows users to manage their tasks and notes on those tasks via the CLI. The tasks are stored in a json file at your home directory within the hidden file .taskmanTasks.json

# Install Instructions

For now I'm solely maintaing this tool for myself, and don't plan on building out the usability too much past where it is now. That's not to say this can't be changed in the future...

If you find yourself wanting to play with my toy all you need to do is clone this repo into your GOPATH and compile:

git clone https://github.com/FlippinBerger/taskmanatee 
cd taskmanatee
go build

# Usage
Once compiled you simple launch the executable
./task

After you have created tasks, lauching the executable will show you your tasks automatically numbered. The numbers are used to target tasks and notes with future operations

# Add a task
c t "Task to create"

# Delete a task
d t <task_num>

# Mark a task complete
x <task_num>

# Mark a task uncomplete
u <task_num>

# Add a note to task <task_num>
c n <task_num> "Note text here"

# Delete a note <note_num> from task <task_num>
d n <task_num> <note_num>
