import { HttpErrorResponse, HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { Observable } from "rxjs";
import { tap } from "rxjs/operators";
import { UserService } from "../services/user.service";

@Injectable()
export class Interceptor implements HttpInterceptor {
  constructor(private userService: UserService, private router: Router) {}

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    let request = this.authToken(req);
    return next.handle((request)).pipe(tap(()=>{},
        (err: any) => {
            if (err instanceof HttpErrorResponse) {
                if (err.status !== 401){
                    return;
                }
                alert("unauthorized please login !!")
            }
        }));
  }

  authToken(request: HttpRequest<any>){
    const token = this.userService.getToken() ?? "";

    return request.clone({
        setHeaders: {
            'TransId' : `x${this.getRandomInt(1000)}`,
            Authorization:`Bearer ${token}`
        }
    })
  }

  getRandomInt(max) {
    return Math.floor(Math.random() * max);
  }

}
