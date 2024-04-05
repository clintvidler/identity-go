import { Component } from '@angular/core';
import { ServerError } from '../../../interfaces/server-error';
import {
  FormControl,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { IdentityService } from '../../../services/identity.service';
import { SubmitInputComponent } from '../../../components/forms/submit-input/submit-input.component';
import { PasswordInputComponent } from '../../../components/forms/password-input/password-input.component';
import { TextInputComponent } from '../../../components/forms/text-input/text-input.component';
import { EmailInputComponent } from '../../../components/forms/email-input/email-input.component';
import { ServerErrorsComponent } from '../../../components/forms/server-errors/server-errors.component';

@Component({
  selector: 'app-edit',
  standalone: true,
  imports: [
    FormsModule,
    ReactiveFormsModule,
    SubmitInputComponent,
    PasswordInputComponent,
    TextInputComponent,
    EmailInputComponent,
    ServerErrorsComponent,
  ],
  templateUrl: './edit.component.html',
  styleUrl: './edit.component.scss',
})
export class EditComponent {
  constructor(private identityService: IdentityService) {}

  passwordErrors: ServerError[] = [];
  displaynameErrors: ServerError[] = [];
  emailErrors: ServerError[] = [];

  passwordForm = new FormGroup({
    old: new FormControl('', {
      validators: [Validators.required],
      updateOn: 'blur',
    }),
    new: new FormControl('', {
      validators: [Validators.required],
      updateOn: 'blur',
    }),
  });

  displayNameForm = new FormGroup({
    new: new FormControl('', {
      validators: [Validators.required],
      updateOn: 'blur',
    }),
  });

  emailForm = new FormGroup({
    new: new FormControl('', {
      validators: [Validators.required, Validators.email],
      updateOn: 'blur',
    }),
  });

  onPasswordSubmit(): void {
    this.identityService
      .updatePassword(this.passwordForm.value)
      .subscribe((res) => {
        this.identityService.logout().subscribe();
      });
  }

  onDisplaynameSubmit(): void {
    this.identityService
      .updateDisplayName(this.displayNameForm.value)
      .subscribe((res) => {});
  }

  onEmailSubmit(): void {
    this.identityService
      .startUpdateEmail(this.emailForm.value)
      .subscribe((res) => {});
  }
}
