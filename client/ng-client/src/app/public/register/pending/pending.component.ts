import { Component } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import {
  FormControl,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';

import { IdentityService } from '../../../services/identity.service';
import { PasswordInputComponent } from '../../../components/forms/password-input/password-input.component';
import { SubmitInputComponent } from '../../../components/forms/submit-input/submit-input.component';
import { CommonModule } from '@angular/common';
import { TextInputComponent } from '../../../components/forms/text-input/text-input.component';
import { ServerError } from '../../../interfaces/server-error';
import { ServerErrorsComponent } from '../../../components/forms/server-errors/server-errors.component';

@Component({
  selector: 'app-pending',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    TextInputComponent,
    PasswordInputComponent,
    SubmitInputComponent,
    ServerErrorsComponent,
    RouterLink,
  ],
  templateUrl: './pending.component.html',
  styleUrl: './pending.component.scss',
})
export class PendingComponent {
  constructor(
    private identityService: IdentityService,
    private route: ActivatedRoute
  ) {
    this.key = this.route.snapshot.paramMap.get('key') || '';

    this.identityService.pendingRegister(this.key).subscribe((res) => {
      const body = res.body as any;
      this.email = body?.email || '';
    });
  }

  key: string = '';
  email: string = '';
  hideRegisterPassword = true;
  success = false;
  serverErrors: ServerError[] = [];

  registerForm = new FormGroup({
    display_name: new FormControl('', {
      validators: [Validators.required],
      updateOn: 'blur',
    }),
    password: new FormControl('', {
      validators: [Validators.required],
      updateOn: 'blur',
    }),
  });

  onSubmit(): void {
    let display_name = this.registerForm.value.display_name || '';
    let password = this.registerForm.value.password || '';

    const data: Object = { display_name: display_name, password: password };

    this.identityService.finishRegistration(data, this.key).subscribe((res) => {
      switch (res?.status) {
        case 200:
          this.success = true;
          break;
        default:
          this.serverErrors.push(res?.error);
      }
    });
  }

  get registerFormDisplayName() {
    return this.registerForm.get('display_name');
  }

  get registerFormPassword() {
    return this.registerForm.get('password');
  }
}
