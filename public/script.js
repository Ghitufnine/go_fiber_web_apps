// public/script.js (Frontend)

const taskInput = document.getElementById('task');
const taskList = document.getElementById('taskList');

// Function to fetch and display todos
async function fetchTodos() {
    taskList.innerHTML = '';
    const response = await fetch('/api/todos');
    const todos = await response.json();

    todos.forEach(todo => {
        const listItem = document.createElement('li');
        listItem.innerHTML = `
            <input type="checkbox" ${todo.completed ? 'checked' : ''} onclick="toggleComplete(${todo.ID})">
            <span class="${todo.completed ? 'completed' : ''}">${todo.title}</span>
            <button onclick="deleteTodo(${todo.ID})">Delete</button>
        `;
        taskList.appendChild(listItem);
    });
}

// Function to add a new todo
async function addTodo() {
    const title = taskInput.value;
    if (!title) return;

    const response = await fetch('/api/todos', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ title: title, completed: "0" }),
    });

    taskInput.value = '';
    await fetchTodos();
}

// Function to toggle todo completion status
async function toggleComplete(id) {
    const response = await fetch(`/api/todos/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ completed: "1" }),
    });

    await fetchTodos();
}

// Function to delete a todo
async function deleteTodo(id) {
    const response = await fetch(`/api/todos/${id}`, {
        method: 'DELETE',
    });

    await fetchTodos();
}

// Fetch todos on page load
fetchTodos();
