import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UploadFileService {

  constructor(private http: HttpClient) { }

  upload(id: string, files: File[]): Observable<any> {
    console.log(files)
    const formData = new FormData();

    for(let i=0;i<files.length;i++) {
      formData.append("files[]", files[i])
    }

    return this.http.post(`/api/upload/${id}`, formData)
  }
}
