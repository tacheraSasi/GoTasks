```markdown
# GoTasks - A Simple To-Do List Application

GoTasks is a simple to-do list application built with Go and the Fyne UI toolkit. This application allows users to add, delete, and manage tasks with priority levels, all while storing data in a SQLite database.

## Features

- Add tasks with descriptions and priority levels (Low, Medium, High)
- Mark tasks as completed
- Delete tasks or clear all tasks
- Progress bar showing completion status
- Light/Dark theme toggle

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [Building for Linux](#building-for-linux)
- [Building for Android](#building-for-android)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go installed (version 1.17 or later)
- SQLite installed
- Android SDK and NDK (if building for Android)
- Fyne toolkit

### Installing Go

You can install Go by following the instructions on the [official Go website](https://golang.org/doc/install).

### Installing Fyne

To install the Fyne toolkit, run the following command:

```bash
go get fyne.io/fyne/v2
```

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/tacheraSasi/GoTasks.git
   cd GoTasks
   ```

2. **Install dependencies**:
   Make sure you have all the required dependencies installed:
   ```bash
   go mod tidy
   ```

## Running the Application

To run the application locally, execute the following command:

```bash
go run main.go
```

This will start the application, and a window will open for you to manage your tasks.

## Building for Linux

To build the application for Linux, use the following command:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o GoTasks
```

This will create a binary file named `GoTasks` in the current directory.

## Building for Android

To build the application for Android, follow these steps:

1. Ensure that you have the Android SDK and NDK installed and that you have set up `gomobile`.

2. Run the following command to create an APK:

   ```bash
   fyne package -os android -appID com.yourdomain.gotasks -icon path/to/icon.png
   ```

3. Install the APK on your Android device using ADB:

   ```bash
   adb install GoTasks.apk
   ```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any changes you make.


