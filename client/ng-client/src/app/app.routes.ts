import { Routes } from '@angular/router';

import { PrivateComponent } from './private/private.component';
import { UserComponent } from './private/user/user.component';
import { PublicComponent } from './public/public.component';
import { LoginComponent } from './public/login/login.component';
import { isLoggedInGuard } from './guards/is-logged-in.guard';
import { isNotLoggedInGuard } from './guards/is-not-logged-in.guard';
import { EditComponent } from './private/user/edit/edit.component';
import { ViewComponent } from './private/user/view/view.component';
import { LogoutComponent } from './private/user/logout/logout.component';
import { RegisterComponent } from './public/register/register.component';
import { PendingComponent as RegisterPendingComponent } from './public/register/pending/pending.component';
import { ResetPasswordComponent } from './public/reset-password/reset-password.component';
import { PendingComponent as ResetPasswordPendingComponent } from './public/reset-password/pending/pending.component';
import { InvalidKeyComponent } from './public/invalid-key/invalid-key.component';
import { pendingResetPasswordGuard } from './guards/pending-reset-password.guard';
import { pendingRegistrationGuard } from './guards/pending-registration.guard';
import { pendingUpdateEmailGuard } from './guards/pending-update-email.guard';
import { DummyComponent } from './public/dummy/dummy.component';

export const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: '/user' },
  {
    path: '',
    component: PrivateComponent,
    canActivateChild: [isLoggedInGuard],
    children: [
      {
        path: 'user',
        component: UserComponent,
        children: [
          { path: '', pathMatch: 'full', redirectTo: '/user/view' },
          { path: 'view', component: ViewComponent },
          {
            path: 'edit',
            component: EditComponent,
          },
          {
            path: 'edit/email/:key',
            component: DummyComponent,
            canActivate: [isLoggedInGuard, pendingUpdateEmailGuard],
          },
        ],
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
      {
        path: 'login/reset',
        component: ResetPasswordComponent,
        canActivate: [isNotLoggedInGuard],
      },
      {
        path: 'login/reset/:key',
        component: ResetPasswordPendingComponent,
        canActivate: [isNotLoggedInGuard, pendingResetPasswordGuard],
      },
      {
        path: 'register',
        component: RegisterComponent,
        canActivate: [isNotLoggedInGuard],
      },
      {
        path: 'register/:key',
        component: RegisterPendingComponent,
        canActivate: [isNotLoggedInGuard, pendingRegistrationGuard],
      },
      { path: 'logout', component: LogoutComponent },
      { path: 'invalid-key', component: InvalidKeyComponent },
    ],
  },
];
