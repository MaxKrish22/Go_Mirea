package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

// Выводит "#ID Название [статус], приоритет N
func (t Task) Summary() string {
	return fmt.Sprintf("#%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

// Присваивает статус done исходной задаче.
func (t *Task) Complete() {
	t.Status = "done"
}

// Проверяет диапазон от 1 до 3, при ошибке объект не меняется
func (t *Task) SetPriority(priority int) bool {
	if priority < 1 || priority > 3 {
		return false
	}
	t.Priority = priority
	return true
}

// Возвращает &tasks[i] при наличии совпадения, иначе nil.
func findTaskByID(tasks []Task, id int) *Task {
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
	// Находит ID=2, 'завершает' её и устанавливает приоритет 3.
	task2 := findTaskByID(tasks, 2)
	if task2 != nil {
		task2.Complete()
		task2.SetPriority(3)
	}
	// Считает завершенные задачи
	var countdone int = 0
	for i := range tasks {
		if tasks[i].Status == "done" {
			countdone++
		}
	}

	for _, task := range tasks {
		fmt.Println(task.Summary())
	}

	fmt.Printf("Число завершенных задач: %d\n", countdone)

	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача не найдена")
	}
}
