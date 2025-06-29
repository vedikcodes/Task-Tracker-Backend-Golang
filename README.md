# ✅ Task Tracker CLI

A simple command-line application to **track and manage your tasks** using a local JSON file. This tool helps you keep track of what you need to do, what you're currently working on, and what you’ve completed — all from your terminal.

---

## 📝 Description

**Task Tracker CLI** is a lightweight and dependency-free task manager built for the terminal. It allows users to add, update, delete, and mark tasks with statuses like `todo`, `in-progress`, and `done`. All task data is saved locally in a JSON file using the native file system — no external libraries or databases required.

---

## 🚀 Features

- 📌 Add, update, and delete tasks
- 📈 Mark tasks as `in-progress` or `done`
- 📋 List tasks by status (`all`, `done`, `todo`, `in-progress`)
- 🧠 Automatically assigns unique IDs to each task
- 🕒 Tracks creation and last updated timestamps
- 💾 Saves tasks in `tasks.json` (auto-created if not present)

---

## 📂 Example Usage

##Project Reference
https://roadmap.sh/projects/task-tracker

### ➕ Add a new task
```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
