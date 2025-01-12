# Tasks

## Understanding Tasks in Hide

Tasks in Hide are predefined commands or scripts that can be executed within a project's devcontainer environment. They allow coding agents to perform various operations such as running tests, linting code, or executing custom scripts.

Key features of Hide tasks:

1. **Environment-Specific**: Tasks run within the project's devcontainer, ensuring consistency with the project's environment.
2. **Predefined or Ad-hoc**: Tasks can be predefined in the project configuration or executed ad-hoc. The ad-hoc tasks are essentially the shell commands.
3. **Aliased Commands**: Predefined tasks can be given aliases for easy reference.

!!! note

    For all code examples, the server is assumed to be running on `localhost:8080`. Adjust the URL if your Hide server is running on a different host or port.

## Defining Tasks

Tasks are defined in the `devcontainer.json` file or in the project creation request under the `customizations.hide.tasks` section. Here's an example:

```json
{
    "customizations": {
        "hide": {
            "tasks": [
                {
                    "alias": "test",
                    "command": "pytest"
                },
                {
                    "alias": "lint",
                    "command": "flake8 ."
                }
            ]
        }
    }
}
```

## Running Predefined Tasks

To run a predefined task, you need to specify the project ID and the task alias.

=== "curl"

    ```bash
    curl -X POST http://localhost:8080/projects/123/tasks \
      -H "Content-Type: application/json" \
      -d '{
        "alias": "test"
      }'
    ```

=== "python"

    ```python
    # Coming soon
    ```

## Running Ad-hoc Tasks

To run an ad-hoc task, you need to specify the project ID and the task command.

=== "curl"

    ```bash
    curl -X POST http://localhost:8080/projects/123/tasks \
      -H "Content-Type: application/json" \
      -d '{
        "command": "pytest path/to/test_file.py::test_function_name"
      }'
    ```

=== "python"

    ```python
    # Coming soon
    ```

## Running Tasks with a timeout

It's a good practice to set timeout for tasks, especially those generated by the agent. In order to do that add following header to the request

```
X-Timeout-Seconds: <value>
```

## Task Results

When a task is executed, Hide returns a TaskResult object containing:

-   `stdOut`: The standard output of the command
-   `stdErr`: The standard error output of the command
-   `exitCode`: The exit code of the command

Example response:

```json
{
    "stdOut": "All tests passed successfully.\n",
    "stdErr": "",
    "exitCode": 0
}
```

Example response on timeout:

```json
{
    "stdOut": "Test 1 passed successfully.\n",
    "stdErr": "",
    "exitCode": 124
}
```

## Listing Available Tasks

To get a list of all predefined tasks for a project:

=== "curl"

    ```bash
    curl -X GET http://localhost:8080/projects/123/tasks
    ```

=== "python"

    ```python
    # Coming soon
    ```

This will return a list of all predefined tasks for the project with id `123`.

## Notes

-   Tasks are executed synchronously, meaning that the request will be blocked until the task completes. The support for asynchronous tasks is coming soon!
