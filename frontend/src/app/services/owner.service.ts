import { Mypet } from './../models/mypet';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Owner } from '../models/onwer';

@Injectable({
  providedIn: 'root'
})
export class OwnerService {

  constructor(private http: HttpClient) { }

  addOwners(owners: Owner[]){
    return this.http.post('/api/owner', owners)
  }

  getOwners(userId: string): Observable<Mypet[]>{
    return this.http.get<Mypet[]>(`/api/owner/${userId}`)
  }

  deleOwner(userId: string): Observable<any> {
    return this.http.delete(`/api/owner/${userId}`);
  }
}
