import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-submit-input',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './submit-input.component.html',
  styleUrl: './submit-input.component.scss',
})
export class SubmitInputComponent {
  @Input('label') label: string = 'Submit';
}
