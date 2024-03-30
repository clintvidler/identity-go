import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { BehaviorSubject, Observable, catchError, map, of, tap } from 'rxjs';

import { environment } from '../../environments/environment';
import { LoginCredential, User } from '../interfaces/user';

@Injectable({
  providedIn: 'root'
})
export class IdentityService {
  constructor(private http: HttpClient, private router: Router) {}

  httpOptions: Object = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
    observe: 'response'
  };

  // Current user

  private userSubject = new BehaviorSubject<User | null>(null);

  get user(): User | null {
    return this.userSubject.value;
  }

  set user(user: User | null) {
    this.userSubject.next(user);
  }

  profile(): Observable<User> {
    return this.http
      .get<User>(`${environment.server}/user`, this.httpOptions)
      .pipe(
        map((res: any) => {
          console.warn(res);
          return res.body as User;
        }),
        catchError(err => {
          return null as any;
        })
      ) as Observable<User>;
  }

  currentUser(): Observable<User> {
    return this.http
      .get<User>(`${environment.server}/current-user`, this.httpOptions)
      .pipe(
        map((res: any) => {
          console.warn(res);
          return res.body as User;
        }),
        catchError(err => {
          return null as any;
        })
      ) as Observable<User>;
  }

  // Login

  login(data: LoginCredential): Observable<any> {
    return this.http
      .post<Response>(`${environment.server}/login`, data, this.httpOptions)
      .pipe(
        map(res => {
          var accessToken = res.headers.get('grpc-metadata-access-token');
          var refreshToken = res.headers.get('grpc-metadata-refresh-token');
          console.warn('Access token:', accessToken);
          console.warn('Refresh token:', refreshToken);

          return res;
        }),
        catchError(this.handleError<any[]>('login', []))
      );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      console.log(`${operation} failed: ${error.message}`);

      // // TODO: send the error to remote logging infrastructure
      // console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      // return of(result as T);

      return of(error as T);
    };
  }

  /** Log a HeroService message with the MessageService */
  private log(message: string) {
    // this.messageService.add(`HeroService: ${message}`);
  }
}
