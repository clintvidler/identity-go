import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import {
  FormControl,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';

import { SubmitInputComponent } from '../../../components/forms/submit-input/submit-input.component';
import { ServerErrorsComponent } from '../../../components/forms/server-errors/server-errors.component';
import { IdentityService } from '../../../services/identity.service';
import { PasswordInputComponent } from '../../../components/forms/password-input/password-input.component';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { ServerError } from '../../../interfaces/server-error';

@Component({
  selector: 'app-pending',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    RouterLink,
    SubmitInputComponent,
    ServerErrorsComponent,
    PasswordInputComponent,
  ],
  templateUrl: './pending.component.html',
  styleUrl: './pending.component.scss',
})
export class PendingComponent implements OnInit {
  constructor(
    private identityService: IdentityService,
    private route: ActivatedRoute
  ) {
    this.key = this.route.snapshot.paramMap.get('key') || '';
  }

  ngOnInit(): void {
    this.identityService.pendingResetPassword(this.key).subscribe((res) => {
      const body = res.body as any;
      this.email = body?.email || '';
    });
  }

  key: string = '';
  success = false;
  errors: ServerError[] = [];
  email: string | undefined;

  resetForm = new FormGroup({
    password: new FormControl('', {
      validators: [Validators.required],
      updateOn: 'blur',
    }),
  });

  onSubmit(): void {
    this.identityService
      .finishResetPassword(this.resetForm.value, this.key)
      .subscribe((res) => {
        switch (res?.status) {
          case 200:
            this.success = true;
            break;
          default:
            this.errors.push(res?.error);
        }
      });
  }
}
