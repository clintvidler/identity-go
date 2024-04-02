import { CommonModule } from '@angular/common';
import { Component, Input, forwardRef } from '@angular/core';
import {
  FormsModule,
  NG_VALUE_ACCESSOR,
  ReactiveFormsModule,
} from '@angular/forms';

import { ValidationErrorsComponent } from '../validation-errors/validation-errors.component';
import { ControlValueAccessorDirective } from '../../../directives/control-value-accessor.directive';

@Component({
  selector: 'app-text-input',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    FormsModule,
    ValidationErrorsComponent,
  ],
  templateUrl: './text-input.component.html',
  styleUrl: './text-input.component.scss',
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => TextInputComponent),
      multi: true,
    },
  ],
})
export class TextInputComponent<T> extends ControlValueAccessorDirective<T> {
  @Input() inputId: string = 'text-input';
  @Input() label: string = 'Text';
  @Input() placeholder: string = 'Enter text';
  @Input() customErrorMessages: Record<string, string> = {};
}
