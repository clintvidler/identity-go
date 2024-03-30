import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink, RouterOutlet } from '@angular/router';

import { User } from '../../interfaces/user';
import { IdentityService } from '../../services/identity.service';

@Component({
  selector: 'app-user',
  standalone: true,
  imports: [CommonModule, RouterOutlet, RouterLink],
  templateUrl: './user.component.html',
  styleUrl: './user.component.scss',
})
export class UserComponent implements OnInit {
  constructor(private identityService: IdentityService) {}

  user: User | null = null;

  ngOnInit(): void {
    this.user = this.identityService.user;
    this.identityService.profile().subscribe((res) => {
      console.warn(res);
      this.user = res;
    });
  }
}
