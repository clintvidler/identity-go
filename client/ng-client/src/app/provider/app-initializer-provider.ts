import { Injectable } from '@angular/core';
import { IdentityService } from '../services/identity.service';

@Injectable({ providedIn: 'root' })
export class AppInitialiserProvider {
  constructor(private identityService: IdentityService) {}

  load(): Promise<void> {
    return new Promise((resolve, reject) => {
      this.identityService.CurrentUser().subscribe({
        next: () => {
          resolve();
        },
        error: (error) => {
          // reject(error);
          resolve();
        },
      });
    });
  }
}
