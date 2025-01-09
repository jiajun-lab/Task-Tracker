package main

import (
	"fmt"
	"os"
	"encoding/json"
	"time"
	"strconv"
)

type Task struct {
	ID 			int 		`json:"id"`
	Description string 		`json:"description"`
	Status		string 		`json:"status"`
	CreatedAt 	string 		`json:"created_at"`
	UpdatedAt 	string 		`json:"updated_at"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

const taskFileName = "./task_data/task.json"

// ReadTasks 从指定文件中读取任务列表
// fileName: 任务数据文件的路径
// 返回值: 任务列表切片和可能的错误
func ReadTasks(fileName string) ([]Task, error) {
    // 读取文件内容
    data, err := os.ReadFile(fileName)
    if err != nil {
        return nil, err // 如果读取失败,返回错误
    }
    
    var tasks []Task
    // 将JSON数据解析为任务切片
    err = json.Unmarshal(data, &tasks)
    if err != nil {
        return nil, err // 如果解析失败,返回错误
    }
    
    return tasks, nil // 返回成功读取的任务列表
}

// WriteTasks 将任务列表写入指定文件
// fileName: 要写入的文件路径
// tasks: 要保存的任务列表
// 返回值: 可能的错误
func WriteTasks(fileName string, tasks []Task) error {
    // 将任务列表转换为格式化的JSON数据
    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err // 如果JSON编码失败,返回错误
    }
    
    // 将JSON数据写入文件
    return os.WriteFile(fileName, data, 0644)
}

func AddTask(tasks []Task, description string) ([]Task, error){
	newTask := Task{
		ID: len(tasks) + 1,
		Description: description,
		Status: "todo",
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	tasks = append(tasks, newTask)
	return tasks, nil
}

func UpdateTask(tasks []Task, id int, description string) ([]Task, error){
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			return tasks, nil
		}
	}
	fmt.Println("Task not found")
	return tasks, nil
}

func DeleteTask(tasks []Task, id int) ([]Task, error){
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks, nil
		}
	}
	fmt.Println("Task not found")
	return tasks, nil
}

func MarkInProgress(tasks []Task, id int) ([]Task, error){
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = "in-progress"
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			return tasks, nil
		}
	}
	fmt.Println("Task not found")
	return tasks, nil
}

func MarkDone(tasks []Task, id int) ([]Task, error){
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = "done"
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			return tasks, nil
		}
	}
	fmt.Println("Task not found")
	return tasks, nil
}

func ListTasks(tasks []Task, status string){
	if status == "" {
		for _, task := range tasks {
			fmt.Printf("%d. %s - %s - %s\n", task.ID, task.Description, task.Status, task.UpdatedAt)
		}
		return
	}else if status == "todo"{
		for _, task := range tasks {
			if task.Status == status {
				fmt.Printf("%d. %s - %s - %s\n", task.ID, task.Description, task.Status, task.UpdatedAt)
			}
		}
		return	
	}else if status == "in-progress" {
		for _, task := range tasks {
			if task.Status == "in-progress" {
				fmt.Printf("%d. %s - %s - %s\n", task.ID, task.Description, task.Status, task.UpdatedAt)
			}
		}
		return
	}else if status == "done" {
		for _, task := range tasks {
			if task.Status == "done" {
				fmt.Printf("%d. %s - %s - %s\n", task.ID, task.Description, task.Status, task.UpdatedAt)
			}
		}
		return
	}
}


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: taskcli <command> [args], type taskcli help for more information")
		os.Exit(1)
	}
	
	tasks, err := ReadTasks(taskFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
		case "add":
			if len(os.Args) < 3 {
				fmt.Println("Usage: taskcli add <description>")
				os.Exit(1)
			}
			description := os.Args[2]
			tasks, err = AddTask(tasks, description)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}else{
				fmt.Println("Task added successfully")
			}

		case "update":
			if len(os.Args) < 4 {
				fmt.Println("Usage: taskcli update <id> <description>")
				os.Exit(1)
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			description := os.Args[3]
			tasks, err = UpdateTask(tasks, id, description)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}else{
				fmt.Println("Task updated successfully")
			}

		case "delete":
			if len(os.Args) < 3 {
				fmt.Println("Usage: taskcli delete <id>")
				os.Exit(1)
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			tasks, err = DeleteTask(tasks, id)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}else{
				fmt.Println("Task deleted successfully")
			}

		case "list":
			if len(os.Args) < 2 {
				fmt.Println("Usage: taskcli list <status>")
				os.Exit(1)
			}
			if len(os.Args) == 2 {
				ListTasks(tasks, "")
				os.Exit(0)
			}
			status := os.Args[2]
			ListTasks(tasks, status)
		case "mark-in-progress":
			if len(os.Args) < 3 {
				fmt.Println("Usage: taskcli mark-in-progress <id>")
				os.Exit(1)
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			tasks, err = MarkInProgress(tasks, id)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case "mark-done":
			if len(os.Args) < 3 {
				fmt.Println("Usage: taskcli mark-done <id>")
				os.Exit(1)
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			tasks, err = MarkDone(tasks, id)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		case "help":
			fmt.Println("Usage: taskcli <command> [args]")
			fmt.Println("Commands:")
			fmt.Println("  add <description> - Add a new task")
			fmt.Println("  update <id> <description> - Update the description of a task")
			fmt.Println("  delete <id> - Delete a task")
			fmt.Println("  list - List all tasks")
			fmt.Println("  mark-in-progress <id> - Mark a task as in progress")
			fmt.Println("  mark-done <id> - Mark a task as done")
			fmt.Println("  help - Show this help message")
			os.Exit(0)
		default:
			fmt.Println("Unknown command, type taskcli help for more information")
			os.Exit(1)
	}

	err = WriteTasks(taskFileName, tasks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}