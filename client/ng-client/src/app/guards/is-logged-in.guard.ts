import { Injectable, inject } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanActivateFn,
  Router,
  RouterStateSnapshot,
} from '@angular/router';
import { Observable, catchError, map, of } from 'rxjs';

import { IdentityService } from '../services/identity.service';
import { User } from '../interfaces/user';

@Injectable()
export class IsLoggedInService {
  constructor(public router: Router, public identityService: IdentityService) {}

  canActivateChild(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<any> {
    return this.identityService.currentUser().pipe(
      map((res: User) => {
        this.identityService.user = res;

        return of(true);
      }),
      catchError(() => {
        this.router.navigate(['/login']);

        return of(false);
      })
    );
  }
}

export const isLoggedInGuard: CanActivateFn = (route, state) => {
  // console.warn('isLoggedInGuard()');

  return inject(IsLoggedInService).canActivateChild(route, state);
};
