import { TestBed } from '@angular/core/testing';
import { CanActivateFn } from '@angular/router';

import { pendingResetPasswordGuard } from './pending-reset-password.guard';

describe('pendingResetPasswordGuard', () => {
  const executeGuard: CanActivateFn = (...guardParameters) =>
    TestBed.runInInjectionContext(() =>
      pendingResetPasswordGuard(...guardParameters)
    );

  beforeEach(() => {
    TestBed.configureTestingModule({});
  });

  it('should be created', () => {
    expect(executeGuard).toBeTruthy();
  });
});
