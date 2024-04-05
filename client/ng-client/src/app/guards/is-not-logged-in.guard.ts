import { Injectable, inject } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanActivateFn,
  Router,
  RouterStateSnapshot,
} from '@angular/router';
import { Observable, catchError, map, of } from 'rxjs';

import { IdentityService } from '../services/identity.service';

@Injectable()
export class IsNotLoggedInService {
  constructor(public router: Router, public identityService: IdentityService) {}

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<any> {
    return this.identityService.IsAuth().pipe(
      map(() => {
        this.router.navigate(['/user']);

        return of(false);
      }),
      catchError(() => {
        return of(true);
      })
    );
  }
}

export const isNotLoggedInGuard: CanActivateFn = (route, state) => {
  return inject(IsNotLoggedInService).canActivate(route, state);
};
