import { Component, OnInit } from '@angular/core';
import {
  FormControl,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { RouterLink } from '@angular/router';

import { IdentityService } from '../../services/identity.service';
import { ServerError } from '../../interfaces/server-error';
import { CommonModule } from '@angular/common';
import { SubmitInputComponent } from '../../components/forms/submit-input/submit-input.component';
import { EmailInputComponent } from '../../components/forms/email-input/email-input.component';
import { ServerErrorsComponent } from '../../components/forms/server-errors/server-errors.component';

@Component({
  selector: 'app-reset-password',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    RouterLink,
    SubmitInputComponent,
    EmailInputComponent,
    ServerErrorsComponent,
  ],
  templateUrl: './reset-password.component.html',
  styleUrl: './reset-password.component.scss',
})
export class ResetPasswordComponent implements OnInit {
  constructor(private identityService: IdentityService) {}

  ngOnInit() {
    this.resetForm.valueChanges.subscribe((_) => {
      this.errors = [];
    });
  }

  success = false;
  errors: ServerError[] = [];

  resetForm = new FormGroup({
    email: new FormControl('', {
      validators: [Validators.required, Validators.email],
      updateOn: 'blur',
    }),
  });

  onSubmit(): void {
    this.identityService
      .startResetPassword(this.resetForm.value)
      .subscribe((result) => {
        switch (result?.status) {
          case 200:
            this.success = true;
            break;
          default:
            this.errors = result.error
              ? result.error
              : 'unexpected error has occured';
        }
      });
  }
}
