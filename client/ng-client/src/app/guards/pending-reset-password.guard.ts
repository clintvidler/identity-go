import { Injectable, inject } from '@angular/core';
import {
  CanActivateFn,
  Router,
  ActivatedRouteSnapshot,
  RouterStateSnapshot,
} from '@angular/router';
import { Observable, of } from 'rxjs';
import { catchError, map } from 'rxjs/operators';

import { IdentityService } from '../services/identity.service';

@Injectable()
export class PendingResetPasswordService {
  constructor(
    private identityService: IdentityService,
    public router: Router
  ) {}

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<any> {
    return this.identityService.pendingResetPassword(route.params['key']).pipe(
      map((res: any) => {
        return of(res['status'] === 200);
      }),
      catchError(() => {
        this.router.navigate(['/invalid-key']);

        return of(false);
      })
    );
  }
}

export const pendingResetPasswordGuard: CanActivateFn = (route, state) => {
  return inject(PendingResetPasswordService).canActivate(route, state);
};
