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

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    FormsModule,
    ReactiveFormsModule,
    EmailInputComponent,
    PasswordInputComponent,
    SubmitInputComponent,
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

  serverError: string = '';

  errorMessages = {};

  loginForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required]),
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
          this.serverError = res?.error
            ? res.error
            : 'unexpected error has occured';
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
