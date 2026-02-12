# CLI Task Tracker

A minimal task tracker built with **Go**.  
Fast, lightweight, and designed for terminal freaks.

## Features

- Add tasks
- List tasks
- Mark tasks as done / in-progress
- Clear completed tasks
- Persistent JSON storage

## Installation

```bash
git clone https://github.com/z66x/learn-backend.git
cd learn-backend/tracker
go build -o tracker
```
## Usage

```bash
tracker add "build the next big thing"
tracker list
tracker done 1
tracker progress 1
tracker clear
```
