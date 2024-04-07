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
export class PendingUpdateEmailService {
  constructor(
    private identityService: IdentityService,
    public router: Router
  ) {}

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<any> {
    return this.identityService.finishUpdateEmail(route.params['key']).pipe(
      map((res: any) => {
        this.identityService.logout().subscribe();
      }),
      catchError(() => {
        this.router.navigate(['/invalid-key']);

        return of(false);
      })
    );
  }
}

export const pendingUpdateEmailGuard: CanActivateFn = (route, state) => {
  return inject(PendingUpdateEmailService).canActivate(route, state);
};
