/*
   =====================================================
   ==================== GoTask App =====================
   =====================================================

   GoTask is a simple and fun todo application built using Go and the Fyne GUI framework.
   Created by **Tachera Sasi**, this app helps users manage their tasks effortlessly.

   Key Features:
   -----------------------------------------------------
   - Create tasks with priorities.
   - Mark tasks as completed.
   - System and in-app notifications.

   Technologies Used:
   -----------------------------------------------------
   - **Go**: The Go programming language, known for its simplicity and speed.
   - **Fyne**: A modern UI toolkit for building graphical apps in Go.

   GoTask combines Go’s efficiency with a clean, user-friendly design to make task management easy and enjoyable.

   =====================================================
   =================== Tachera Sasi ====================
   =====================================================
*/

package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

// Task structure to store task data
type Task struct {
	ID          int
	Description string
	Priority    string
	Completed   bool
}

func main() {
	myApp := app.New()

	
	// Load your icon as a byte slice
	iconData, err := os.ReadFile("assets/icon.png") // Path to your icon file
	if err != nil {
		log.Fatal("Failed to load icon:", err) // Handle error appropriately
	}

	// Create a static resource for the icon
	iconResource := fyne.NewStaticResource("icon.png", iconData)

	// Set the icon for the application
	myApp.SetIcon(iconResource)

	myWindow := myApp.NewWindow("GoTask")

	// db connection
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		ShowNotification(myApp, myWindow,"Error","Failed to connect to the db",true)
	}
	defer db.Close()

	// Create tasks table if it doesn't exist
	createTable(db)

	// Fetch existing tasks from the database
	taskList := fetchTasksFromDB(db)

	// Progress bar for completed tasks
	progressBar := widget.NewProgressBar()
	updateProgress(progressBar, db)
	

	// Task List Display
	taskListView := widget.NewList(
		func() int { return len(taskList) },
		func() fyne.CanvasObject { return container.NewHBox(widget.NewCheck("", nil), widget.NewLabel("")) },
		func(i int, item fyne.CanvasObject) {
			box := item.(*fyne.Container)
			check := box.Objects[0].(*widget.Check)
			label := box.Objects[1].(*widget.Label)
			label.SetText(taskList[i].Description + " (" + taskList[i].Priority + ")")
			check.SetChecked(taskList[i].Completed)
			check.OnChanged = func(checked bool) {
				taskList[i].Completed = checked
				updateTaskCompletion(db, taskList[i].ID, checked)
				updateProgress(progressBar, db)
			}
		},
	)

	taskEntry := widget.NewEntry()
	taskEntry.SetPlaceHolder("Enter new task")

	prioritySelect := widget.NewSelect([]string{"Low", "Medium", "High"}, func(value string) {})
	prioritySelect.PlaceHolder = "Select Priority"

	addButton := widget.NewButtonWithIcon("Add Task", theme.ContentAddIcon(), func() {
		if taskEntry.Text != "" && prioritySelect.Selected != "" {
			taskID := addTask(db, taskEntry.Text, prioritySelect.Selected)
			if taskID > 0 {
				taskList = fetchTasksFromDB(db)
				taskEntry.SetText("")
				prioritySelect.ClearSelected()
				taskListView.Refresh()
				updateProgress(progressBar, db)
			}
		}
	})

	deleteButton := widget.NewButtonWithIcon("Delete Last Task", theme.ContentRemoveIcon(), func() {
		if len(taskList) > 0 {
			lastTask := taskList[len(taskList)-1]
			delTask(db, lastTask.ID)
			taskList = fetchTasksFromDB(db)
			taskListView.Refresh()
			updateProgress(progressBar, db)
		}
	})

	clearButton := widget.NewButtonWithIcon("Clear All Tasks", theme.DeleteIcon(), func() {
		clearTasks(db)
		taskList = fetchTasksFromDB(db)
		taskListView.Refresh()
		updateProgress(progressBar, db)
	})

	topBar := container.NewVBox(
		taskEntry,
		prioritySelect,
		container.NewHBox(addButton, deleteButton, clearButton),
		progressBar,
	)

	content := container.NewBorder(topBar, nil, nil, nil, taskListView)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(700, 600))
	myWindow.CenterOnScreen()

	myWindow.ShowAndRun()
}

func updateProgress(progressBar *widget.ProgressBar, db *sql.DB) {
	taskList := fetchTasksFromDB(db)
	if len(taskList) == 0 {
		progressBar.SetValue(0)
	} else {
		completedCount := 0
		for _, task := range taskList {
			if task.Completed {
				completedCount++
			}
		}
		progressBar.SetValue(float64(completedCount) / float64(len(taskList)))
	}
}

func fetchTasksFromDB(db *sql.DB) []Task {
	rows := getTasks(db)
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Description, &task.Priority, &task.Completed)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func updateTaskCompletion(db *sql.DB, taskID int, completed bool) {
	completeTask(db, strconv.Itoa(taskID), completed)
}


