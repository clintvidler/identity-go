import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import {
  BehaviorSubject,
  Observable,
  catchError,
  map,
  of,
  tap,
  throwError,
} from 'rxjs';
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

  public currentUserSubject = new BehaviorSubject<User | null>(null);

  get currentUser(): User | null {
    return this.currentUserSubject.value;
  }

  set currentUser(user: User | null) {
    this.currentUserSubject.next(user);
  }

  // Check the access token with the server
  IsAuth(): Observable<boolean> {
    return this.http
      .get<Response>(`${environment.server}/auth`, this.httpOptions)
      .pipe(
        map((res: Response) => {
          return true;
        }),
        catchError((err) => {
          // console.warn(err);
          return null as any;
        })
      ) as Observable<boolean>;
  }

  // Read user profile
  CurrentUser(): Observable<User> {
    return this.http
      .get<Response>(`${environment.server}/user`, this.httpOptions)
      .pipe(
        map((res: Response) => {
          const user = res.body as unknown as User;
          this.currentUser = user;
          return user;
        }),
        catchError((err) => {
          // console.error('Error fetching user:', err);
          return throwError('Failed to fetch user data');
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

          this.cookieService.set('access', accessToken, { path: '/' });
          this.cookieService.set('refresh', refreshToken, { path: '/' });

          this.CurrentUser().subscribe();

          return res;
        }),
        catchError(this.handleError<any[]>('login', []))
      );
  }

  // Refresh token: The refresh interceptor uses this method to exchange a saved refresh token for a new refresh token and access token
  refreshToken(): Observable<any> {
    return this.http
      .get<Response>(`${environment.server}/refresh`, this.httpOptions)
      .pipe(
        map((res) => {
          // this.localStorage.setItem('refreshToken', res.body);
          var accessToken = res.headers.get('grpc-metadata-access-token') || '';
          var refreshToken =
            res.headers.get('grpc-metadata-refresh-token') || '';

          this.cookieService.set('access', accessToken, { path: '/' });
          this.cookieService.set('refresh', refreshToken, { path: '/' });

          return res;
        })
      );
  }

  // Logout: Delete browser JWT, and submit the refresh token to the backend handler
  logout(): Observable<void> {
    const refreshToken = this.cookieService.get('refresh');

    this.cookieService.delete('access', '/');
    this.cookieService.delete('refresh', '/');

    return this.http
      .post<void>(
        `${environment.server}/logout`,
        { RefreshToken: refreshToken },
        this.httpOptions
      )
      .pipe(
        map((res) => {
          this.router.navigate(['/login']);

          this.currentUser = null;

          return res;
        }),
        catchError(() => {
          this.router.navigate(['/login']);

          this.currentUser = null;

          return of();
        }),
        tap((res) => this.handleError<any[]>('logoutRequest', [res]))
      );
  }

  // Register

  startRegister(data: any): Observable<any> {
    return this.http
      .post<Response>(`${environment.server}/register`, data, this.httpOptions)
      .pipe(
        map((res) => {
          return res;
        }),
        catchError(this.handleError<any[]>('register', []))
      );
  }

  pendingRegister(key: string): Observable<Response> {
    return this.http.get<Response>(
      `${environment.server}/register/${key}`,
      this.httpOptions
    );
  }

  finishRegistration(data: any, key: string): Observable<any> {
    return this.http
      .post<Response>(
        `${environment.server}/register/${key}`,
        data,
        this.httpOptions
      )
      .pipe(
        map((res) => {
          return res;
        }),
        // tap((result) => console.log(result)),
        catchError(this.handleError<any[]>('register', []))
      );
  }

  // Reset password

  startResetPassword(data: any): Observable<any> {
    console.warn('startResetPassword', data);

    return this.http
      .post(`${environment.server}/reset-password`, data, this.httpOptions)
      .pipe(tap((result) => console.log(result)));
  }

  pendingResetPassword(key: string): Observable<any> {
    console.warn(key);

    return this.http.get(
      `${environment.server}/reset-password/${key}`,
      this.httpOptions
    );
  }

  finishResetPassword(data: any, key: string): Observable<any> {
    console.warn('finishResetPassword', data, key);

    return this.http
      .post(
        `${environment.server}/reset-password/${key}`,
        data,
        this.httpOptions
      )
      .pipe(tap((result) => console.log(result)));
  }

  // Update password

  updatePassword(data: any): Observable<any> {
    console.warn('updatePassword', data);

    return this.http.post<Response>(
      `${environment.server}/update-password`,
      data,
      this.httpOptions
    );
  }

  // Update username

  updateDisplayName(data: any): Observable<any> {
    console.warn('updateDisplayName', data);

    return this.http.post<Response>(
      `${environment.server}/update-display-name`,
      data,
      this.httpOptions
    );
  }

  // Update email

  startUpdateEmail(data: any): Observable<any> {
    console.warn('startUpdateEmail', data);

    return this.http.post(
      `${environment.server}/update-email`,
      data,
      this.httpOptions
    );
  }

  finishUpdateEmail(key: string): Observable<any> {
    console.warn('finishUpdateEmail', key);

    return this.http.post(
      `${environment.server}/update-email/${key}`,
      {},
      this.httpOptions
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
