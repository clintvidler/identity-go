import { ApplicationConfig } from '@angular/core';
import { provideRouter } from '@angular/router';
import { provideHttpClient, withInterceptors } from '@angular/common/http';

import { routes } from './app.routes';
import { credentialInterceptor } from './interceptors/credential-interceptor.interceptor';
import { IsLoggedInService } from './guards/is-logged-in.guard';
import { IsNotLoggedInService } from './guards/is-not-logged-in.guard';

export const appConfig: ApplicationConfig = {
  providers: [
    provideRouter(routes),
    provideHttpClient(withInterceptors([credentialInterceptor])),
    IsLoggedInService,
    IsNotLoggedInService
  ]
};
