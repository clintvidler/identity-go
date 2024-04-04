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
import { ServerError } from '../../interfaces/server-error';
import { ServerErrorsComponent } from '../../components/forms/server-errors/server-errors.component';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    EmailInputComponent,
    SubmitInputComponent,
    ServerErrorsComponent,
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
  serverErrors: ServerError[] = [];

  registerForm = new FormGroup({
    email: new FormControl('', {
      validators: [Validators.required, Validators.email],
      updateOn: 'blur',
    }),
  });

  onSubmit(): void {
    let email = this.registerForm.value.email || '';

    this.serverErrors = [];

    this.identityService.register({ email: email }).subscribe((res) => {
      switch (res?.status) {
        case 200:
          this.success = true;
          break;
        default:
          this.serverErrors.push(res?.error);
      }
    });
  }

  get registerFormEmail() {
    return this.registerForm.get('email');
  }
}
