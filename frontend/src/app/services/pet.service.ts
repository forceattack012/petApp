import { Dashboard } from './../models/dashboard';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Pet } from '../models/pet';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class PetService {

  constructor(private http: HttpClient) { }

  getPets(page: number, limit: number): Observable<Dashboard> {
    return this.http.get<Dashboard>(`${environment.petPath}?page=${page}&limit=${limit}`);
  }

  createPet(pet: Pet): Observable<any>{
    return this.http.post<any>(`${environment.petPath}`, pet)
  }

  remove(id: string): Observable<any>{
    return this.http.delete(`${environment.petPath}/${id}`)
  }

  getPetById(id : string): Observable<any>{
    return this.http.get<any>(`${environment.petPath}/${id}`)
  }

  update(id : string, pet: Pet): Observable<any>{
    return this.http.patch<any>(`${environment.petPath}/${id}`, pet)
  }
}
