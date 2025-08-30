# 📌 GitHub Activity CLI

A simple command-line tool written in **Go** that fetches and displays recent public activity (commits, pull requests, issues, stars, forks, etc.) of a GitHub user.

## ⚡ Features
- Fetches recent public events of a GitHub user from the GitHub API.
- Displays activity such as:
  - Pushes
  - Pull Requests
  - Issues
  - Stars
  - Forks
  - Repository creations
- Lightweight and easy to use.

## 📦 Installation
Clone this repository:
```bash
git clone https://github.com/prateek-pradhan/github-activity.git
cd github-activity
```

Build the binary:
```bash
go build -o github-activity
```

(Optional) Move it to your `$PATH` for global access:
```bash
mv github-activity /usr/local/bin/
```

## 🚀 Usage
Run the CLI with a GitHub username:
```bash
./github-activity <username>
```

### Example:
```bash
./github-activity prateek-pradhan
```

📜 Output:
```
Created in prateek-pradhan/skyscanner-android-app at 2025-08-15T08:36:35Z
Created in prateek-pradhan/skyscanner-android-app at 2025-08-15T08:34:55Z
```

## ⚙️ Development
Run directly without building:
```bash
go run main.go <username>
```

## 🔗 Reference
This project is inspired by the [GitHub Activity Tracker project](https://roadmap.sh/projects/github-user-activity).

## 🛠️ Tech Stack
- [Go](https://golang.org/) (Cobra CLI)
- [GitHub REST API v3](https://docs.github.com/en/rest)

