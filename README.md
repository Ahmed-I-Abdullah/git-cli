# Git-CLI (CPSC 559 Final Project)

## Team Members

| Group \#:      |  2   |
| -------------- | --- |
| Student Names: | Ahmed Abdullah |
|                | Rayyan Khalil |
|                | Ammar Elzeftawy |
|                | Rahat Islam |
|                | Ahmed Waly |

## Overview

This project provides a Command Line Interface (CLI) wrapper for Git commands, integrating GRPC support for specific tasks. It intercepts standard Git commands (like `push`, `pull`, `init`) to either execute them directly or forward other commands to a GRPC service based on the provided configuration.

## Features

- Execute standard Git commands directly through the CLI.
- Use GRPC for non-Git commands, leveraging libp2p for peer-to-peer communication.
- Easy configuration of peer URL for GRPC services.

## Prerequisites

Before you begin, ensure you have met the following requirements:
- You have installed the latest version of [Go](https://golang.org/dl/).
- You have a basic understanding of how Git and GRPC work.
- (Optional) You have libp2p and a GRPC server running if you wish to use GRPC related commands.

## Installing Git-CLI

To install Git-CLI, follow these steps:

For Linux and macOS:

```bash
git clone https://github.com/yourusername/git-cli.git
cd git-cli
go build -o git-peer ./cmd/cli/main.go
sudo mv git-peer /usr/local/bin
```

For Windows:

```bash
git clone https://github.com/yourusername/git-cli.git
cd git-cli
go build -o git-peer.exe ./cmd/cli/main.go
```

### Move the Binary to a Directory in Your PATH

For the command to be globally accessible, the binary needs to be in a directory that's in your system's PATH.

#### Windows

You could move it to a directory like `C:\Windows\system32` or any other directory that's in the PATH environment variable. To find directories in your PATH, you can run `echo %PATH%` in your command prompt.

Alternatively, you can add the directory where your binary is to the PATH variable:

1. Find the path to your `git-cli` directory. For example, if it's in `C:\Users\amzef\Documents\git-cli`, that's the path you'll use.
2. Right-click on 'This PC' or 'Computer' on your desktop or in File Explorer.
3. Click 'Properties'.
4. Click 'Advanced system settings'.
5. Click 'Environment Variables'.
6. Under 'System Variables', scroll down and find the `Path` variable, then click 'Edit'.
7. Add the full path to your `git-cli` directory.
8. Click 'OK' to close all dialog boxes.

#### Linux and macOS

You may need to move the binary to a directory that is in your PATH if it is not already. A common directory to place system-wide executables is `/usr/local/bin`. This can typically be done with:

```bash
sudo mv git-cli /usr/local/bin
```

## Using Git-CLI

Here is how you can use the Git-CLI:

Replace `<git-command>` with any git command you wish to run, like `push`, `add`, `commit`, etc., and replace `<your-grpc-command>` with the command you wish to execute via GRPC.

The following are some of the main commands supported by the Git-CLI:

1. Run `git-peer init (repo name)` to initialize a bare repo with that repo name
2. Run `git-peer pull (repo name)` to pull the latest changes from that repo name
3. Run `git-peer add` to add any new local changes to the staging area
4. Run `git-peer commit -m "(message)"` to commit any changes in the staging area
5. Run `git-peer push` to push all commits to the leader branch
