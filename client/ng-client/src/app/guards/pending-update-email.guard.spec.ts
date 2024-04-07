import { TestBed } from '@angular/core/testing';
import { CanActivateFn } from '@angular/router';

import { pendingUpdateEmailGuard } from './pending-update-email.guard';

describe('pendingUpdateEmailGuard', () => {
  const executeGuard: CanActivateFn = (...guardParameters) =>
    TestBed.runInInjectionContext(() =>
      pendingUpdateEmailGuard(...guardParameters)
    );

  beforeEach(() => {
    TestBed.configureTestingModule({});
  });

  it('should be created', () => {
    expect(executeGuard).toBeTruthy();
  });
});
