import {ChangeDetectionStrategy, Component, TemplateRef} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {MatCheckbox} from '@angular/material/checkbox';
import {MatDivider} from '@angular/material/list';
import {MatError, MatFormField, MatInput, MatLabel} from '@angular/material/input';
import {MatIcon} from '@angular/material/icon';
import {MatDialog, MatDialogClose, MatDialogContent} from '@angular/material/dialog';

@Component({
  selector: 'app-tasks',
  imports: [MatCardModule, MatButtonModule, FormsModule, MatCheckbox, MatDivider, MatIcon, MatDialogContent, ReactiveFormsModule, MatFormField, MatLabel, MatInput, MatDialogClose, MatError],
  templateUrl: './tasks.html',
  styleUrl: './tasks.css',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class Tasks {

  //Validation
  isTaskTitleValid = false;
  isTaskDescriptionValid = false;

  // Check if the task is new
  isTaskNew = false;

  searchTerm = '';
  task: TaskModel = {id: '', title: '', description: '', completed: false};


  filteredTasks: TaskModel[] = [];
  tasks: TaskModel[] = [
    {id: '1', title: 'Task 1', description: 'Description 1', completed: false},
    {id: '2', title: 'Task 2', description: 'Description 2', completed: true},
    {id: '3', title: 'Task 3', description: 'Description 3', completed: false},
    {id: '4', title: 'Task 4', description: 'Description 4', completed: false},
  ];

  constructor(private dialog: MatDialog) {
    this.filteredTasks = this.tasks;
  }

  ngOnInit() {
    this.filteredTasks = this.tasks;
  }


// adding task
  addTask(templateRef: TemplateRef<any>) {
    this.task = {id: '', title: '', description: '', completed: false};
    this.openDialog(templateRef, this.task, true);
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
  saveOrUpdateTask() {
    const task = this.task;
    const index = this.tasks.findIndex(t => t.id === task.id);


    // Validate task title
    if (task.title.trim() === '') {
      this.isTaskTitleValid = false;
      return;
    }

    // Validate task description
    if (task.description.trim() === '') {
      this.isTaskDescriptionValid = false;
      return;
    }

    if (index !== -1) {
      this.tasks[index] = task;
    } else {
      task.id = (this.tasks.length + 2).toString();
      this.tasks.push(task);
    }

    this.filteredTasks = this.tasks;
    this.dialog.closeAll();
  }

  //Delete task
  deleteTask(id: string) {
    const index = this.tasks.findIndex(t => t.id === id);
    if (index !== -1) {
      this.filteredTasks.splice(index, 1);
    }
  }

  openDialog(dialogRef: TemplateRef<any>, task: TaskModel, isNew = false) {
    this.isTaskNew = isNew;
    this.task = task;
    this.dialog.open(dialogRef, {
      width: '400px',
    })
  }

}
