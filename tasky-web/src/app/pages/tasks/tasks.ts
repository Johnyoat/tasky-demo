import {Component} from '@angular/core';
import {MatButton} from '@angular/material/button';

@Component({
  selector: 'app-tasks',
  imports: [
    MatButton
  ],
  templateUrl: './tasks.html',
  styleUrl: './tasks.css',
})
export class Tasks {
  tasks: TaskModel[] = [
    {id: '1', title: 'Task 1', description: 'Description 1', completed: false},
    {id: '2', title: 'Task 2', description: 'Description 2', completed: true},
    {id: '3', title: 'Task 3', description: 'Description 3', completed: false},
    {id: '4', title: 'Task 4', description: 'Description 4', completed: false},
  ];
  task:TaskModel = {id: '', title: '', description: '', completed: false};

  constructor() {

  }


  addTask() {
   console.log("add task")
    this.task = {id: '', title: '', description: '', completed: false};
  }


  viewTask(task: TaskModel) {
    console.log(task);
  }
}
