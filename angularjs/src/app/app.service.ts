import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class AppService {

    url = process.env['API_URL'] || '';

    constructor(private http: HttpClient) { }

    save(board: any): Observable<any> {
        const headers = this.getHeader()
        return this.http.post(this.url.concat('/save'), board, { headers: headers });
    }

    get(): Observable<any> {
        const headers = this.getHeader()    
        return this.http.get(this.url.concat('/get'), { headers: headers });
    }

    getHeader(): HttpHeaders {
        const token: string | null = localStorage.getItem('token');
        const headers = new HttpHeaders({
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token
        });

        return headers;
    }
}