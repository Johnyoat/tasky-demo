import {ChangeDetectionStrategy, ChangeDetectorRef, Component, OnInit, TemplateRef} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {MatCheckbox} from '@angular/material/checkbox';
import {MatDivider} from '@angular/material/list';
import {MatError, MatFormField, MatInput, MatLabel} from '@angular/material/input';
import {MatIcon} from '@angular/material/icon';
import {MatDialog, MatDialogClose, MatDialogContent} from '@angular/material/dialog';
import {Api} from '../../data/api';
import {TaskModel} from '../../../models/task.model';

@Component({
  selector: 'app-tasks',
  imports: [MatCardModule, MatButtonModule, FormsModule, MatCheckbox, MatDivider, MatIcon, MatDialogContent, ReactiveFormsModule, MatFormField, MatLabel, MatInput, MatDialogClose, MatError],
  templateUrl: './tasks.html',
  styleUrl: './tasks.css',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class Tasks implements OnInit {

  //Validation
  isTaskTitleValid = true;
  isTaskDescriptionValid = true;

  // Check if the task is new
  isTaskNew = false;

  searchTerm = '';
  task: TaskModel = {id: '', title: '', description: '', completed: false};


  filteredTasks: TaskModel[] = [];
  tasks: TaskModel[] = [];

  constructor(private dialog: MatDialog, private api: Api, private cdr: ChangeDetectorRef) {
  }

  ngOnInit() {
    this.loadTasks();
  }

  loadTasks() {
    this.api.getTasks().subscribe(tasks => {
      this.tasks = tasks;
      this.filteredTasks = tasks;
      this.cdr.markForCheck();
    });
  }


// adding task
  addTask(templateRef: TemplateRef<any>) {
    this.task = {id: '', title: '', description: '', completed: false};
    this.openDialog(templateRef, this.task, true);
  }

  search() {
    //Check if the search term is empty
    if (this.searchTerm.trim() === '') {
      this.filteredTasks = this.tasks;
      this.cdr.markForCheck();
      return;
    }
    // Filter tasks based on the search term
    this.filteredTasks = this.tasks.filter(task => {
      return task.id.toLowerCase().includes(this.searchTerm.toLowerCase()) || task.title.toLowerCase().includes(this.searchTerm.toLowerCase())
    });
    this.cdr.markForCheck();
  }

  //Toggle task
  toggleTask(id: string) {
    const task = this.tasks.find(t => t.id === id);
    if (task) {
      this.api.updateTask(id, {completed: !task.completed}).subscribe(updatedTask => {
        task.completed = updatedTask.completed;
        this.cdr.markForCheck();
      });
    }
  }

  //Update task
  saveOrUpdateTask() {
    const task = this.task;

    // Reset validation states
    this.isTaskTitleValid = true;
    this.isTaskDescriptionValid = true;

    // Validate task title
    if (task.title.trim() === '') {
      this.isTaskTitleValid = false;
    }

    // Validate task description
    if (task.description.trim() === '') {
      this.isTaskDescriptionValid = false;
    }

    if (!this.isTaskTitleValid || !this.isTaskDescriptionValid) {
      return;
    }

    if (this.isTaskNew) {
      this.api.createTask(task).subscribe(() => {
        this.loadTasks();
        this.dialog.closeAll();
      });
    } else {
      this.api.updateTask(task.id, task).subscribe(() => {
        this.loadTasks();
        this.dialog.closeAll();
      });
    }
  }

  //Delete task
  deleteTask(id: string) {
    this.api.deleteTask(id).subscribe(() => {
      this.loadTasks();
    });
  }

  openDialog(dialogRef: TemplateRef<any>, task: TaskModel, isNew = false) {
    this.isTaskNew = isNew;
    this.task = {...task};
    this.isTaskTitleValid = true;
    this.isTaskDescriptionValid = true;
    this.dialog.open(dialogRef, {
      width: '400px',
    })
  }

}
