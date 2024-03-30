import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';

import { IdentityService } from './services/identity.service';
import { LoginCredential } from './interfaces/user';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'ng-client';

  constructor(private identityService: IdentityService) {}

  login() {
    const lc: LoginCredential = { Email: 'x@x', Password: 'x' };

    this.identityService.login(lc).subscribe(res => {
      var accessToken = res.headers.get('grpc-metadata-access-token');
      var refreshToken = res.headers.get('grpc-metadata-refresh-token');
      console.warn('Access token:', accessToken);
      console.warn('Refresh token:', refreshToken);
    });
  }
}
