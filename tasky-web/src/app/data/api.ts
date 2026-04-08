import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class Api {

  getTasks(){
    console.log("get")
  }

  updateTask(){
    console.log("update")
  }
  deleteTask(){
    console.log("delete")
  }

}
