import {
  HttpEvent,
  HttpHandler,
  HttpHandlerFn,
  HttpInterceptorFn,
  HttpRequest,
} from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { Router } from '@angular/router';
import {
  BehaviorSubject,
  Observable,
  catchError,
  concatMap,
  filter,
  finalize,
  take,
  throwError,
} from 'rxjs';
import { IdentityService } from '../services/identity.service';

@Injectable()
export class RefreshService {
  constructor(public router: Router, public identityService: IdentityService) {}

  isRefreshing = false;

  refreshed = new BehaviorSubject<boolean>(false);

  intercept(
    req: HttpRequest<any>,
    next: HttpHandlerFn
  ): Observable<HttpEvent<any>> {
    // console.warn('refreshInterceptor()intercept()');

    return next(req).pipe(
      catchError((err) => {
        // console.warn(err?.error?.message);
        if (
          // [401].includes(err.status)
          [500].includes(err.status) &&
          err?.error?.message == 'Access token: Token is expired'
          //  && this.localStorageService.getItem('refreshToken')
        ) {
          // console.warn(err?.error?.message);
          return this.handleRefresh(req, next);
        }

        return throwError(err);
      })
    );
  }

  handleRefresh(req: HttpRequest<any>, next: HttpHandlerFn): Observable<any> {
    // console.warn('RefreshInterceptor().handleRefresh()');

    if (this.isRefreshing) {
      return this.refreshed.pipe(
        filter(Boolean),
        take(1),
        concatMap(() => next(req))
      );
    }

    this.isRefreshing = true;

    // Subsequent requests to wait until token comes back from refresh call
    this.refreshed.next(false);

    return this.identityService.refreshToken().pipe(
      concatMap((res) => {
        // console.warn('refreshToken()Success');
        this.refreshed.next(true);

        return next(req);
      }),
      catchError((err) => {
        // console.warn('refreshToken()Error');
        // this.localStorageService.removeItem('refreshToken');

        return throwError(() => new Error(err));
      }),
      finalize(() => {
        this.isRefreshing = false;
      })
    );
  }
}

export const refreshInterceptor: HttpInterceptorFn = (req, next) => {
  // console.warn('refreshInterceptor()');

  // return next(req);

  return inject(RefreshService).intercept(req, next);
};
