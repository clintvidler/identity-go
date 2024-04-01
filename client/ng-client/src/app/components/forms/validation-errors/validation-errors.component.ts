import { Component, Input, OnChanges, SimpleChanges } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-validation-errors',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './validation-errors.component.html',
  styleUrl: './validation-errors.component.scss',
})
export class ValidationErrorsComponent implements OnChanges {
  @Input() errors: Record<string, string> | null = {};
  @Input() customErrorMessages: Record<string, string> = {};

  // errorMessages: Record<string, string> = {
  //   required: 'This field is required',
  // };

  errorMessages: Record<string, string> = {};

  ngOnChanges(changes: SimpleChanges): void {
    const { customErrorMessages } = changes;

    if (customErrorMessages) {
      this.errors = {
        ...this.errorMessages,
        ...customErrorMessages.currentValue,
      };
    }
  }
}
