import { Component, Input, forwardRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import {
  FormsModule,
  NG_VALUE_ACCESSOR,
  ReactiveFormsModule,
} from '@angular/forms';

import { ControlValueAccessorDirective } from '../../../directives/control-value-accessor.directive';
import { ValidationErrorsComponent } from '../validation-errors/validation-errors.component';

@Component({
  selector: 'app-password-input',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    FormsModule,
    ValidationErrorsComponent,
  ],
  templateUrl: './password-input.component.html',
  styleUrl: './password-input.component.scss',
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => PasswordInputComponent),
      multi: true,
    },
  ],
})
export class PasswordInputComponent<
  T
> extends ControlValueAccessorDirective<T> {
  @Input() inputId: string = 'password-input';
  @Input() label: string = 'Password';
  @Input() placeholder: string = 'Enter your password';
  @Input() sufix!: string;
  @Input() customErrorMessages: Record<string, string> = {};
}
