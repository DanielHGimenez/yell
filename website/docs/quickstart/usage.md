---
sidebar_position: 2
---

# Usage

Yell can be used in various ways to alert you about the completion of tasks. Below are some common usage examples.

## Type of Alerts

Yell supports different types of alerts to notify you about task completion:

### Beep (default)

Plays a generated terminal sound. Examples:

```sh
yell
```

Or:

```sh
yell --beep
```

### Sound

Plays a MP3 sound. Example:

```sh
yell --sound
```
Go to [configuration page](configuration) to learn more.

### WhatsApp

Sends an alert message to a WhatsApp group. Example:

```sh
yell --wpp --group "Group Name" --message "Task completed"
```

:::info
To use WhatsApp alerts, you need to configure your WhatsApp API credentials executing the following command.
```sh
yell config whatsapp
```
It will show a QR code that you need to scan with your WhatsApp to authenticate the API access.<br/>
It only need to be done once.
:::

## Command Monitoring

You can use Yell to execute a command and alert you when it finishes. For example:

```sh
yell --sound on complete command -- sleep 5
```
In the example above, Yell will play a sound notification when the <code>sleep</code> command completes after 5 seconds.

Another example:
```sh
yell --wpp --group "Group Name" --message "Task completed" on complete command -- sleep 5
```
In the example above, Yell will send a WhatsApp message to the specified group when the <code>sleep</code> command completes after 5 seconds.

### Command Flags

Yell supports the following flags for command alerts:

- `--timeout` or `-t`: Define the maximum wait time for the command to complete. If not specified, will wait indefinitely.
    
    Example:
    ```sh
    yell --sound on complete command --timeout 2 -- sleep 5
    ```
    In the example above, Yell will wait 2 seconds before killing the command execution, and still play a sound notification.

- `--output` or `-o`: Attach the command output to yell output, this way you can see the command output. If not set, the output will be hidden.

    Example:
    ```sh
    yell --sound on complete command --output -- echo "Hello, World!"
    ```
    In the example above, Yell will show `Hello, World!` in the output.

## Process Monitoring

Yell can also monitor processes and alert you when they exit. For example:

```sh
sleep 10 &   
    [1] 3835
yell --sound on complete process --pid 3835
```
In the example above, sleep command was called in the background printing the process id (PID), and Yell was called after to play a sound notification when the process with PID 3835 exits.

### Process Flags

Yell supports the following flags for process alerts:

:::info
Only one of the flags can be used at a time. If all three flags are specified, Yell will stop after finding the first match in the following sequence: `--pid`, `--name`, `--contains`.
:::

- `--pid` or `-p`: Define the process ID to monitor.

    Example:
    ```sh
    yell --sound on complete process --pid 3835
    ```
    In the example above, Yell will wait for the process to complete/exits, and then play a sound notification.

- `--name` or `-n`: Define the process name to monitor.

    Example:
    ```sh
    yell --sound on complete process --name "sleep"
    ```
    In the example above, Yell will wait for the process with the name "sleep" to complete/exits, and then play a sound notification.

- `--contains` or `-c`: Define a substring to match against the process name.

    Example:
    ```sh
    yell --sound on complete process --contains "sl"
    ```
    In the example above, Yell will wait for any process with a name containing "sl" to complete/exits, and then play a sound notification.

## No Monitoring

It's possible to use alerts without using monitoring.

Examples:
```sh
sleep 5 && yell --sound
```
```sh
sleep 5 || yell --sound
```
