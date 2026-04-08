import {Component} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {FormsModule} from '@angular/forms';
import {MatCheckbox} from '@angular/material/checkbox';
import {MatDivider} from '@angular/material/list';
import {MatInput} from '@angular/material/input';
import {MatIcon} from '@angular/material/icon';

@Component({
  selector: 'app-tasks',
  imports: [MatCardModule, MatButtonModule, FormsModule, MatCheckbox, MatDivider, MatInput, MatIcon],
  templateUrl: './tasks.html',
  styleUrl: './tasks.css',
})
export class Tasks {
  searchTerm = '';
  task:TaskModel = {id: '', title: '', description: '', completed: false};
  filteredTasks: TaskModel[] = [];
  tasks: TaskModel[] = [
    {id: '1', title: 'Task 1', description: 'Description 1', completed: false},
    {id: '2', title: 'Task 2', description: 'Description 2', completed: true},
    {id: '3', title: 'Task 3', description: 'Description 3', completed: false},
    {id: '4', title: 'Task 4', description: 'Description 4', completed: false},
  ];
constructor() {
  this.filteredTasks = this.tasks;
}


  addTask() {
   console.log("add task")
    this.task = {id: '', title: '', description: '', completed: false};
  }


  viewTask(task: TaskModel) {
    console.log(task);
  }


  search() {
    console.log("search")
    //Check if the search term is empty
    if (this.searchTerm.trim() === '') {
      this.filteredTasks = this.tasks;
      return;
    }
    // Filter tasks based on the search term
    this.filteredTasks = this.tasks.filter(task => {
      return task.id.toLowerCase().includes(this.searchTerm.toLowerCase()) || task.title.toLowerCase().includes(this.searchTerm.toLowerCase())
    });
  }
  //Toggle task
  toggleTask(id: string) {
    const task = this.tasks.find(t => t.id === id);
    if (task) {
      task.completed = !task.completed;
    }
  }

  //Update task
  updateTask(task: TaskModel) {
    const index = this.tasks.findIndex(t => t.id === task.id);
    if (index !== -1) {
      this.tasks[index] = task;
    }
  }

  //Delete task
  deleteTask(id: string) {
    const index = this.tasks.findIndex(t => t.id === id);
    if (index !== -1) {
      this.tasks.splice(index, 1);
    }
  }

}
