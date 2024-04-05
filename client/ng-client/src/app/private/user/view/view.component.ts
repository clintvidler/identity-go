import { Component } from '@angular/core';
import { IdentityService } from '../../../services/identity.service';
import { User } from '../../../interfaces/user';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-view',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './view.component.html',
  styleUrl: './view.component.scss',
})
export class ViewComponent {
  constructor(private identityService: IdentityService) {}

  user: User | null = null;

  ngOnInit(): void {
    this.identityService.currentUserSubject.subscribe((user: User | null) => {
      this.user = user;
    });
  }
}
