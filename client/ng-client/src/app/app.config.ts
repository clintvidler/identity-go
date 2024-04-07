import { APP_INITIALIZER, ApplicationConfig } from '@angular/core';
import { provideRouter } from '@angular/router';
import { provideHttpClient, withInterceptors } from '@angular/common/http';

import { routes } from './app.routes';
import { credentialInterceptor } from './interceptors/credential-interceptor.interceptor';
import { IsLoggedInService } from './guards/is-logged-in.guard';
import { IsNotLoggedInService } from './guards/is-not-logged-in.guard';
import {
  RefreshService,
  refreshInterceptor,
} from './interceptors/refresh.interceptor';
import { AppInitialiserProvider } from './providers/app-initializer-provider';
import { PendingResetPasswordService } from './guards/pending-reset-password.guard';
import { PendingRegistrationService } from './guards/pending-registration.guard';
import { PendingUpdateEmailService } from './guards/pending-update-email.guard';

export const appConfig: ApplicationConfig = {
  providers: [
    provideRouter(routes),
    provideHttpClient(
      withInterceptors([credentialInterceptor, refreshInterceptor])
    ),
    IsLoggedInService,
    IsNotLoggedInService,
    PendingRegistrationService,
    PendingResetPasswordService,
    PendingUpdateEmailService,
    RefreshService,
    {
      provide: APP_INITIALIZER,
      useFactory: (p: AppInitialiserProvider) => () => p.load(),
      deps: [AppInitialiserProvider],
      multi: true,
    },
  ],
};
