<h1 align="center">
  <a href="https://github.com/zxctatar/CLI-task-tracker-GO" target="_blank">CLI-task-tracker</a>
</h1>

<h3 align="center">CLI application for task tracking</h3>

## Features

- Add a new task
- Update a task
- Delete a task
- Mark a task as not started, started, or done
- List all tasks

## Installation

1. Clone the repository

```bash
git clone https://github.com/zxctatar/CLI-task-tracker-GO.git
```

2. Go to src and build the project<br>

```bash
go build
```


3. Start task tracker<br>
    
```bash
./CLI-task-tracker
```

## Usage
```bash
add <task description>                      - Add a new task.
update <task id> <new task description>     - Update the description of a task.
list / list<argument>                       - Show tasks. Arguments: not-started, started, done
delete <task id>                            - Delete a task.
start <task id>                             - Start a task.
done <task id>                              - Mark a task as done.
cancel <task id>                            - Cancel a task and mark it as not started.
exit                                        - Exit the application.
help                                        - Show available commands.
```

## License

Distributed under the MIT License. See `LICENSE` for more information.