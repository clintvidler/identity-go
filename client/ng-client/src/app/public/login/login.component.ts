import { Component } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import {
  FormControl,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';

import { EmailInputComponent } from '../../components/forms/email-input/email-input.component';
import { PasswordInputComponent } from '../../components/forms/password-input/password-input.component';
import { SubmitInputComponent } from '../../components/forms/submit-input/submit-input.component';
import { IdentityService } from '../../services/identity.service';
import { LoginCredential } from '../../interfaces/user';
import { ServerErrorsComponent } from '../../components/forms/server-errors/server-errors.component';
import { ServerError } from '../../interfaces/server-error';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    FormsModule,
    ReactiveFormsModule,
    EmailInputComponent,
    PasswordInputComponent,
    SubmitInputComponent,
    ServerErrorsComponent,
    RouterLink,
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss',
})
export class LoginComponent {
  constructor(
    private identityService: IdentityService,
    private router: Router
  ) {}

  serverErrors: ServerError[] = [];

  loginForm = new FormGroup({
    email: new FormControl('', {
      validators: [Validators.required, Validators.email],
      updateOn: 'blur',
    }),
    password: new FormControl('', {
      validators: [Validators.required],
      updateOn: 'blur',
    }),
  });

  onSubmit(): void {
    let email = this.loginForm.value.email || '';
    let password = this.loginForm.value.password || '';

    const lc: LoginCredential = { email: email, password: password };

    this.identityService.login(lc).subscribe((res) => {
      switch (res?.status) {
        case 200:
          {
            this.router.navigate(['/user']);
          }
          break;
        default:
          this.serverErrors.push(res?.error);
      }
    });
  }

  get loginFormEmail() {
    return this.loginForm.get('email');
  }

  get loginFormPassword() {
    return this.loginForm.get('password');
  }
}
