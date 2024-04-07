import { TestBed } from '@angular/core/testing';
import { CanActivateFn } from '@angular/router';

import { pendingRegistrationGuard } from './pending-registration.guard';

describe('pendingRegistrationGuard', () => {
  const executeGuard: CanActivateFn = (...guardParameters) =>
    TestBed.runInInjectionContext(() =>
      pendingRegistrationGuard(...guardParameters)
    );

  beforeEach(() => {
    TestBed.configureTestingModule({});
  });

  it('should be created', () => {
    expect(executeGuard).toBeTruthy();
  });
});
