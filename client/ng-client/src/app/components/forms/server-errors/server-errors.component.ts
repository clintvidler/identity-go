import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import { ServerError } from '../../../interfaces/server-error';

@Component({
  selector: 'app-server-errors',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './server-errors.component.html',
  styleUrl: './server-errors.component.scss',
})
export class ServerErrorsComponent {
  @Input() serverErrors: ServerError[] = [];
}
