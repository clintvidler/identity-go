import { Component } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import {
  FormControl,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';

import { IdentityService } from '../../services/identity.service';
import { EmailInputComponent } from '../../components/forms/email-input/email-input.component';
import { SubmitInputComponent } from '../../components/forms/submit-input/submit-input.component';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    EmailInputComponent,
    SubmitInputComponent,
    RouterLink,
  ],
  templateUrl: './register.component.html',
  styleUrl: './register.component.scss',
})
export class RegisterComponent {
  constructor(
    private identityService: IdentityService,
    private router: Router
  ) {}

  success = false;

  serverError: string = '';

  errorMessages = {};

  registerForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
  });

  onSubmit(): void {
    let email = this.registerForm.value.email || '';

    this.identityService.register({ email: email }).subscribe((res) => {
      switch (res?.status) {
        case 200:
          {
            this.success = true;
          }
          break;
        default:
          this.serverError = res?.error
            ? res.error
            : 'unexpected error has occured';
      }
    });
  }

  get registerFormEmail() {
    return this.registerForm.get('email');
  }
}
