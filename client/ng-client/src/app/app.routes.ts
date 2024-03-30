import { Routes } from '@angular/router';

import { PrivateComponent } from './private/private.component';
import { UserComponent } from './private/user/user.component';
import { PublicComponent } from './public/public.component';
import { LoginComponent } from './public/login/login.component';
import { isLoggedInGuard } from './guards/is-logged-in.guard';
import { isNotLoggedInGuard } from './guards/is-not-logged-in.guard';

export const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: '/login' },
  {
    path: '',
    component: PrivateComponent,
    canActivateChild: [isLoggedInGuard],
    children: [
      {
        path: 'user',
        component: UserComponent,
        children: [],
      },
    ],
  },
  {
    path: '',
    component: PublicComponent,
    children: [
      {
        path: 'login',
        component: LoginComponent,
        canActivate: [isNotLoggedInGuard],
      },
    ],
  },
];
