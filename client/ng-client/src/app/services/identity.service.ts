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
      // 'Grpc-Metadata-Access':
      //   'eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxMjcuMC4wLjE6NTQxMTgiLCJleHAiOjE3MTE4Mzc4ODQsInN1YiI6IjEifQ.CaIrFNG-v2bcIaJk29HfSNbDJy6YqGH-VqtJ3C0w6o4MXH34W6hKkzUcnHOG3aYybZ0kQUKaPfxi65YuzYh8zdXaraxaVEh_N7V6IloMsBS9knVXZGFR6KhYuRxEUull9HlKJxI4TUO60laQdySRQkGZuKBen2GhmmPAheLw7vnz_bILQV4rLfy9xMgTbCuk2xJ2hG5272v6TcMMTwewVsbdTDwsDN4rUrVhGy5jBMTKqtZchfw3jobgOYQasGQisfcEMUSDZaluTlAFHZThw-YqgRqFlVHKVRP4rbNehUKSyPZmlyICUsC2-2Z1P5dkwdLvwAsCmDx0LJdibrDAGFDn76C5IeqF0J8S-wAUkrBg6Cecu3F8eF7RbR10eIKLiF6jo5cd51Bk15CDYa3EXX260AkyZjxE_YXOf8hvPs9xUjbKDodlW1Rxpie146xj31aUCEKZ_Tcss-dIgMKP-3krZ9fSrqDES7hVsGACg_PE2c-lNbQxdtPC-3QEuxh3wwKGuPnEKr4TbSkjVgfhD5wczs4HPdtcFpG8Zs1UwNc5E5_zE8SX97zlLr8f3t9r_ENqUIDcLUx9ScTyslDTOAphXOjIMwt_FEJocfe0GS_ej7RBctwc6m6tNp0QKkKN0YFXvarR5NZIL7Uvjz1BOmBtAkU21Ujw7RDglaOaw74'
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

  currentUser(): Observable<User> {
    // return this.http
    //   .get<User>(`${environment.server}/current-user`, this.httpOptions)
    //   .pipe(
    //     map((res: any) => {
    //       console.warn(res);
    //       return res.body as User;
    //     }),
    //     catchError(err => {
    //       return null as any;
    //     })
    //   ) as Observable<User>;

    return this.http
      .get<User>(`${environment.server}/user`, this.httpOptions)
      .pipe(
        map((res: any) => {
          console.warn(res);
          return res.body as User;
        }),
        catchError((err) => {
          console.warn(err);
          return null as any;
        })
      ) as Observable<User>;
  }

  // Login

  login(data: LoginCredential): Observable<any> {
    return this.http
      .post<Response>(`${environment.server}/login`, data, this.httpOptions)
      .pipe(
        map((res) => {
          var accessToken = res.headers.get('grpc-metadata-access-token') || '';
          var refreshToken =
            res.headers.get('grpc-metadata-refresh-token') || '';

          this.cookieService.set('at', accessToken);
          this.cookieService.set('rt', refreshToken);

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
