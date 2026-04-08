import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, map } from 'rxjs';
import { TaskModel } from '../../models/task.model';

@Injectable({
  providedIn: 'root',
})
export class Api {
  private apiUrl = 'http://localhost:3000/api/v1/tasks';

  constructor(private http: HttpClient) {}

  getTasks(): Observable<TaskModel[]> {
    return this.http.get<{ status: number; message: string; data: TaskModel[] }>(this.apiUrl).pipe(
      map(response => response.data || [])
    );
  }

  createTask(task: Partial<TaskModel>): Observable<TaskModel> {
    return this.http.post<{ status: number; message: string; data: TaskModel }>(this.apiUrl, task).pipe(
      map(response => response.data)
    );
  }

  updateTask(id: string, task: Partial<TaskModel>): Observable<TaskModel> {
    return this.http.put<{ status: number; message: string; data: TaskModel }>(`${this.apiUrl}/${id}`, task).pipe(
      map(response => response.data)
    );
  }

  deleteTask(id: string): Observable<any> {
    return this.http.delete(`${this.apiUrl}/${id}`);
  }
}
