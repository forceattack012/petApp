import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { User } from '../models/user';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient) { }

  login(user: User): Observable<any> {
    return this.http.post(`${environment.loginPath}`, user)
  }

  saveToken(token: string) {
    localStorage.setItem('token', token);
  }

  saveName(name: string) {
    localStorage.setItem('name', name);
  }

  saveUserId(userId: string) {
    localStorage.setItem('userId', userId);
  }

  getUserId(){
    return localStorage.getItem('userId');
  }

  getName(){
    return localStorage.getItem('name');
  }

  getToken() {
    return localStorage.getItem('token');
  }
}
