import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { BehaviorSubject, Observable, catchError, map, of, tap } from 'rxjs';

import { CookieService } from 'ngx-cookie-service';

import { environment } from '../../environments/environment';
import { LoginCredential, User } from '../interfaces/user';

@Injectable({
  providedIn: 'root',
})
export class IdentityService {
  constructor(
    private http: HttpClient,
    private router: Router,
    private cookieService: CookieService
  ) {}

  httpOptions: Object = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json',
    }),
    observe: 'response',
  };

  // Current user
  private userSubject = new BehaviorSubject<User | null>(null);

  get user(): User | null {
    return this.userSubject.value;
  }

  set user(user: User | null) {
    this.userSubject.next(user);
  }

  // TODO: this is currently doing the same as the knownUser method, these probably need different endpoints on the backend
  profile(): Observable<User> {
    return this.http
      .get<User>(`${environment.server}/user`, this.httpOptions)
      .pipe(
        map((res: any) => {
          console.warn(res);
          return res.body as User;
        }),
        catchError((err) => {
          return null as any;
        })
      ) as Observable<User>;
  }

  // To check if the user is known, used by the is/not logged in guards
  knownUser(): Observable<User> {
    return this.http
      .get<User>(`${environment.server}/user`, this.httpOptions)
      .pipe(
        map((res: any) => {
          return res.body as User;
        }),
        catchError((err) => {
          console.warn(err);
          return null as any;
        })
      ) as Observable<User>;
  }

  // Login: save JWTs to cookies
  login(data: LoginCredential): Observable<any> {
    return this.http
      .post<Response>(`${environment.server}/login`, data, this.httpOptions)
      .pipe(
        map((res) => {
          var accessToken = res.headers.get('grpc-metadata-access-token') || '';
          var refreshToken =
            res.headers.get('grpc-metadata-refresh-token') || '';

          this.cookieService.set('access', accessToken);
          this.cookieService.set('refresh', refreshToken);

          return res;
        }),
        catchError(this.handleError<any[]>('login', []))
      );
  }

  // Refresh token: The refresh interceptor uses this method to exchange a saved refresh token for a new refresh token and access token
  refreshToken(): Observable<any> {
    return this.http
      .get<Response>(
        `${environment.server}/refresh`,
        // { token: this.localStorage.getItem('refreshToken') },
        // { token: this.cookieService.get('rt') },
        this.httpOptions
      )
      .pipe(
        map((res) => {
          // this.localStorage.setItem('refreshToken', res.body);
          var accessToken = res.headers.get('grpc-metadata-access-token') || '';
          var refreshToken =
            res.headers.get('grpc-metadata-refresh-token') || '';

          this.cookieService.set('access', accessToken);
          this.cookieService.set('refresh', refreshToken);

          return res;
        })
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
