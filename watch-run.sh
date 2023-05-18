#!/bin/bash

# Function to execute the command
run_command() {
    echo "Running command: $@"
    "$@" &

    # Sleep for a short time to allow the program to start
    sleep 2

    # Get the PID of the running Go server process
    PID=$(ps aux | grep "/tmp/go" | grep -v grep | awk '{print $2}')
    echo "Command PID: $PID"
}

# Check if an argument is provided
if [ $# -eq 1 ]; then
    # Use the argument for the command
    COMMAND=$1
else
    # Use the default command
    COMMAND="go run ."
fi

# Directory to monitor for file changes
DIRECTORY="./"

# Run the initial command
run_command $COMMAND

# Monitor the directory for file changes
while true; do
    # Wait for a file change event
    inotifywait -r -e modify,create,delete $DIRECTORY

    # File change detected, kill previous command and run a new one
    kill -9 "$PID"
    run_command $COMMAND
done
