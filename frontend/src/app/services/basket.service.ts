import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Basket } from '../models/basket';

@Injectable({
  providedIn: 'root'
})
export class BasketService {

  constructor(private http: HttpClient) { }

  addBasket(basket: Basket) {
    return this.http.post('/api/basket', basket)
  }

  getBasket(name: string): Observable<Basket> {
    return this.http.get<Basket>(`/api/basket/${name}`)
  }
}
