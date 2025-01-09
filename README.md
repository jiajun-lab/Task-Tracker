# Task-Tracker
https://roadmap.sh/projects/task-tracker

Task-Tracker 是一个简单的任务管理应用程序。可以通过简单的命令行程序来创建、编辑和删除任务。

## 功能特点

- 创建、编辑和删除任务
- 设置任务状态（例如：todo、in-progress、done）
- 查看不同状态的任务列表

## 打包为可执行文件

要将Task-Tracker打包成可执行文件，请按照以下步骤操作：

1. 确保您已经安装了Go语言环境（推荐Go 1.16或更高版本）。

2. 在项目根目录下，打开终端或命令提示符。

3. 对于不同的操作系统，使用以下命令进行打包：

   - 对于Windows：
     ```
     go build -o task-cli.exe main.go
     ```

   - 对于macOS或Linux：
     ```
     go build -o task-cli main.go
     ```

4. 打包完成后，您将在当前目录下看到名为`task-cli`（或在Windows上为`task-cli.exe`）的可执行文件。

5. 现在您可以直接运行这个可执行文件：

   - 在Windows上：
     ```
     .\task-cli.exe
     ```

   - 在macOS或Linux上：
     ```
     ./task-cli
     ```

注意：确保将可执行文件移动到系统的PATH中，或者在运行时提供完整的文件路径，以便从任何位置访问taskcli。

如果您的项目依赖于配置文件或其他资源，请确保这些文件与可执行文件放在同一目录下，或在运行时指定正确的路径。

## 使用

taskc是一个简单的命令行程序，您可以通过`./task-cli help`来查看文档：
```shell
   Usage: task-cli <command> [args]
   Commands:
     add <description> - Add a new task
     update <id> <description> - Update the description of a task
     delete <id> - Delete a task
     list <state> - List all tasks (state: "todo"\"in-progress"\"done"\"") 
     mark-in-progress <id> - Mark a task as in progress
     mark-done <id> - Mark a task as done
     help - Show this help message
```

## 示例
```shell
   # Adding a new task
   task-cli add "Buy groceries"
   # Output: Task added successfully (ID: 1)
   
   # Updating and deleting tasks
   task-cli update 1 "Buy groceries and cook dinner"
   task-cli delete 1
   
   # Marking a task as in progress or done
   task-cli mark-in-progress 1
   task-cli mark-done 1
   
   # Listing all tasks
   task-cli list
   
   # Listing tasks by status
   task-cli list done
   task-cli list todo
   task-cli list in-progress
```
