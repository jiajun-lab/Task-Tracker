# Task-Tracker
https://roadmap.sh/projects/task-tracker

Task-Tracker 是一个简单的任务管理应用程序。可以通过简单的命令行程序来创建、编辑和删除任务。

## 功能特点

- 创建、编辑和删除任务
- 设置任务状态（例如：todo、in-progress、done）
- 查看任务列表和进度报告
- 用户友好的界面，易于使用

## 打包为可执行文件

要将Task-Tracker打包成可执行文件，请按照以下步骤操作：

1. 确保您已经安装了Go语言环境（推荐Go 1.16或更高版本）。

2. 在项目根目录下，打开终端或命令提示符。

3. 对于不同的操作系统，使用以下命令进行打包：

   - 对于Windows：
     ```
     go build -o taskcli.exe main.go
     ```

   - 对于macOS或Linux：
     ```
     go build -o taskcli main.go
     ```

4. 打包完成后，您将在当前目录下看到名为`taskcli`（或在Windows上为`taskcli.exe`）的可执行文件。

5. 现在您可以直接运行这个可执行文件：

   - 在Windows上：
     ```
     .\taskcli.exe
     ```

   - 在macOS或Linux上：
     ```
     ./taskcli
     ```

注意：确保将可执行文件移动到系统的PATH中，或者在运行时提供完整的文件路径，以便从任何位置访问taskcli。

如果您的项目依赖于配置文件或其他资源，请确保这些文件与可执行文件放在同一目录下，或在运行时指定正确的路径。

## 使用

taskc是一个简单的命令行程序，您可以通过`./taskcli help`来查看文档：
```
Usage: taskcli <command> [args]
Commands:
  add <description> - Add a new task
  update <id> <description> - Update the description of a task
  delete <id> - Delete a task
  list - List all tasks
  mark-in-progress <id> - Mark a task as in progress
  mark-done <id> - Mark a task as done
  help - Show this help message
```
