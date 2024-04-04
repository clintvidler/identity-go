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

  changed: boolean = false;

  ngOnChanges(changes: SimpleChanges): void {
    Object.keys(changes).forEach((key) => {
      if (changes[key].previousValue != undefined) {
        this.changed = !changes[key]?.firstChange;
      }
    });
  }

  getMessage(key: string) {
    switch (key) {
      case 'required':
        return 'Required';
      case 'email':
        return 'Enter a valid email';
      default:
        return `Unhandled validation error: ${key}`;
    }
  }
}
