import { Component } from '@angular/core';
import { IdentityService } from '../../../services/identity.service';

@Component({
  selector: 'app-logout',
  standalone: true,
  imports: [],
  templateUrl: './logout.component.html',
  styleUrl: './logout.component.scss',
})
export class LogoutComponent {
  constructor(private identityService: IdentityService) {}

  ngOnInit(): void {
    this.identityService.logout().subscribe(() => {});
  }
}
