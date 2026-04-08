import {Component} from '@angular/core';

@Component({
  selector: 'app-tasks',
  imports: [],
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

  constructor() {

  }
}
