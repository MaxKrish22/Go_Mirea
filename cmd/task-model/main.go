package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func (t Task) Summary() string {
	// TODO: верните "#ID Название [статус], приоритет N".
	return fmt.Sprintf("#%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

func (t *Task) Complete() {
	// TODO: присвойте статус done исходной задаче.
	t.Status = "done"
}

func (t *Task) SetPriority(priority int) bool {
	// TODO: проверьте диапазон 1–3, при ошибке объект не менять.
	if priority < 1 || priority > 3 {
		return false
	}
	t.Priority = priority
	return true
}

func findTaskByID(tasks []Task, id int) *Task {
	// TODO: верните &tasks[i] при совпадении, иначе nil.
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}
	// TODO: найдите ID=2, завершите и установите приоритет 3.
	task2 := findTaskByID(tasks, 2)
	if task2 != nil {
		task2.Complete()
		task2.SetPriority(3)
	}
	// TODO: посчитайте завершённые задачи.
	var countdone int = 0
	for i := range tasks {
		if tasks[i].Status == "done" {
			countdone++
		}
	}
	for _, task := range tasks {
		fmt.Println(task.Summary())
	}
	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача не найдена")
	}
}
